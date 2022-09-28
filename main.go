package main

import (
	"os"
	"strings"
	"time"

	"github.com/champ-r/data-maker/pkg/championgg"
	"github.com/champ-r/data-maker/pkg/common"
	la "github.com/champ-r/data-maker/pkg/lolalytics"
	mb "github.com/champ-r/data-maker/pkg/murderbridge"
	op "github.com/champ-r/data-maker/pkg/opgg"
	"github.com/champ-r/data-maker/pkg/ugg"
	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
)

var Version string

func main() {
	app := &cli.App{
		Name:    "data-maker",
		Version: Version,
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "sources",
				Aliases: []string{"s"},
				Usage:   "Fetch data from specified source",
				Value:   "",
			},
			&cli.BoolFlag{
				Name:    "summonersrift",
				Aliases: []string{"sr"},
				Usage:   "Fetch all sources for Summoner's Rift",
				Value:   false,
			},
			&cli.BoolFlag{
				Name:  "aram",
				Usage: "Fetch all sources for ARAM",
				Value: false,
			},
			&cli.BoolFlag{
				Name:  "urf",
				Usage: "Fetch all sources for URF",
				Value: false,
			},
			&cli.BoolFlag{
				Name:    "all",
				Aliases: []string{"a"},
				Usage:   "Fetch data from all available sources",
				Value:   false,
			},
			&cli.Int64Flag{
				Name:    "timestamp",
				Aliases: []string{"ts"},
				Usage:   "Set timestamp for builds",
			},
			&cli.BoolFlag{
				Name:    "debug",
				Aliases: []string{"d"},
				Usage:   "Enable debug mode",
				Value:   false,
			},
		},
		Action: func(ctx *cli.Context) error {
			sources := strings.Split(ctx.String("sources"), ",")
			for idx, s := range sources {
				sources[idx] = strings.TrimSpace(s)
			}

			enableOpgg := common.Some([]string{"op.gg", "opgg"}, sources) || ctx.Bool("summonersrift") || ctx.Bool("all")
			enableOpggAram := common.Some([]string{"op.gg-aram", "opgg-aram"}, sources) || ctx.Bool("aram") || ctx.Bool("all")
			enableOpggUrf := common.Some([]string{"op.gg-urf", "opgg-urf"}, sources) || ctx.Bool("urf") || ctx.Bool("all")
			enableMb := common.Some([]string{"mb", "murderbridge"}, sources) || ctx.Bool("aram") || ctx.Bool("all")
			enableLolalytics := common.Some([]string{"la", "lolalytics"}, sources) || ctx.Bool("summonersrift") || ctx.Bool("all")
			enableLolalyticsAram := common.Some([]string{"la-aram", "lolalytics-aram"}, sources) || ctx.Bool("aram") || ctx.Bool("all")
			enableLolalyticsUrf := common.Some([]string{"la-urf", "lolalytics-urf"}, sources) || ctx.Bool("urf") || ctx.Bool("all")
			enableUgg := common.Some([]string{"u.gg", "ugg"}, sources) || ctx.Bool("summonersrift") || ctx.Bool("all")
			enableUggAram := common.Some([]string{"u.gg-aram", "ugg-aram"}, sources) || ctx.Bool("aram") || ctx.Bool("all")
			//enableChampionGG := common.Some([]string{"champion.gg", "championgg"}, sources) || ctx.Bool("summonersrift") || ctx.Bool("all")
			enableChampionGG := false
			//enableChampionGGAram := common.Some([]string{"champion.gg-aram", "championgg-aram"}, sources) || ctx.Bool("aram") || ctx.Bool("all")
			enableChampionGGAram := false
			debug := ctx.Bool("debug")

			tasks := 0
			for _, i := range []bool{
				enableOpgg,
				enableOpggAram,
				enableOpggUrf,
				enableMb,
				enableLolalytics,
				enableLolalyticsAram,
				enableLolalyticsUrf,
				enableUgg,
				enableUggAram,
				enableChampionGG,
				enableChampionGGAram,
			} {
				if i {
					tasks++
				}
			}
			if tasks == 0 {
				log.Info("please at least specify 1 source")
				return nil
			}

			timestamp := time.Now().UnixMilli()
			allChampionData, officialVer, err := common.GetChampionList()
			if err != nil {
				log.Fatal(err)
			}
			runeLoopUp, allRunes, err := common.GetRunesReforged(officialVer)
			if err != nil {
				log.Fatal(err)
			}

			var championAliasList = make(map[string]string)
			for k, v := range allChampionData.Data {
				championAliasList[v.Name] = k
			}

			ch := make(chan string, tasks)
			//var opggRet, mbRet, opggAramRet, laRet, laAramRet, opggUrfRet string

			if enableOpgg {
				log.Info("[CMD] Fetch data from op.gg")
				go func() {
					ch <- op.Import(allChampionData.Data, championAliasList, officialVer, timestamp, debug)
				}()
			}
			cData := op.ChampionData{
				AllChampions: allChampionData.Data,
				AliasList:    championAliasList,
				OfficialVer:  officialVer,
				Timestamp:    timestamp,
			}
			if enableOpggAram {
				log.Info("[CMD] Fetch data from op.gg __ARAM__")
				go func() {
					config := op.PkgConfig{
						PkgName:        op.AramPkgName,
						ListUrl:        op.AramSourceUrl,
						IsFeaturedMode: true,
						Mode:           "aram",
					}
					ch <- op.ImportNoPosition(&cData, debug, &config)
				}()
			}
			if enableOpggUrf {
				log.Info("[CMD] Fetch data from op.gg __URF__")
				go func() {
					config := op.PkgConfig{
						PkgName:        op.UrfPkgName,
						ListUrl:        op.UrfSourceUrl,
						IsFeaturedMode: true,
						Mode:           "urf",
					}
					ch <- op.ImportNoPosition(&cData, debug, &config)
				}()
			}

			if enableMb {
				log.Info("[CMD] Fetch data from murderbridge.com")
				go func() {
					ch <- mb.Import(allChampionData.Data, timestamp, runeLoopUp, allRunes, debug)
				}()
			}

			if enableLolalytics {
				log.Info("[CMD] Fetch data from lolalytics.com")
				go func() {
					ch <- la.Import(allChampionData.Data, officialVer, timestamp, runeLoopUp, common.GameModeSR, debug)
				}()
			}
			if enableLolalyticsAram {
				log.Info("[CMD] Fetch data from lolalytics.com __ARAM__")
				go func() {
					ch <- la.Import(allChampionData.Data, officialVer, timestamp, runeLoopUp, common.GameModeARAM, debug)
				}()
			}
			if enableLolalyticsUrf {
				log.Info("[CMD] Fetch data from lolalytics.com __URF__")
				go func() {
					ch <- la.Import(allChampionData.Data, officialVer, timestamp, runeLoopUp, common.GameModeURF, debug)
				}()
			}
			if enableUgg {
				log.Info("[CMD] Fetch data from u.gg")
				go func() {
					ch <- ugg.Import(allChampionData.Data, officialVer, false, timestamp, debug)
				}()
			}
			if enableUggAram {
				log.Info("[CMD] Fetch data from u.gg")
				go func() {
					ch <- ugg.Import(allChampionData.Data, officialVer, true, timestamp, debug)
				}()
			}
			if enableChampionGG {
				log.Info("[CMD] Fetch data from champion.gg")
				go func() {
					ch <- championgg.Import(allChampionData.Data, &runeLoopUp, officialVer, false, timestamp, debug)
				}()
			}
			if enableChampionGGAram {
				log.Info("[CMD] Fetch data from champion.gg ARAM")
				go func() {
					ch <- championgg.Import(allChampionData.Data, &runeLoopUp, officialVer, true, timestamp, debug)
				}()
			}

			for i := 0; i < tasks; i++ {
				ret := <-ch
				log.Info(ret)
			}

			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
