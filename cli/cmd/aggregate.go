package cmd

import (
	"fmt"
	"io/ioutil"
	// "bufio"
	"encoding/json"
	"sort"
	"strings"
	// "os"

	. "github.com/newlee/lemon/domain"
	"github.com/spf13/cobra"
)

var aggregateCmd *cobra.Command = &cobra.Command{
	Use:   "ag",
	Short: "aggregate methods",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		source := cmd.Flag("source").Value.String()
		configFile := cmd.Flag("config").Value.String()
		configData, err := ioutil.ReadFile(configFile)
		if err != nil {
			fmt.Print(err)
		}
		config := NewConfig(configData)

		fanData, err := ioutil.ReadFile(source)
		if err != nil {
			fmt.Print(err)
		}

		fans := make(map[string]map[string]string)
		json.Unmarshal(fanData, &fans)

		// transactions := [][]string{
		// 	{"beer", "nuts", "cheese"},
		// 	{"beer", "nuts", "jam"},
		// 	{"beer", "butter"},
		// 	{"nuts", "cheese"},
		// 	{"beer", "nuts", "cheese", "jam"},
		// 	{"butter"},
		// 	{"beer", "nuts", "jam", "butter"},
		// 	{"jam"},
		// }
		tmap := make(map[string]map[string]string)
		result := make(map[string]int)
		for member, mf := range fans {
			result[member] = 0
			for m, f := range mf {
				vv := m
				if config.ByFile {
					vv = f
				}
				if _, ok := tmap[vv]; !ok {
					tmap[vv] = make(map[string]string)
				}
				tmap[vv][member] = ""
			}
		}
		wordBase := make([][]string, 0)
		for _, v := range tmap {
			vv := make([]string, 0)
			for k, _ := range v {
				vv = append(vv, k)
			}
			wordBase = append(wordBase, vv)
		}
		fmt.Println(wordBase)
		// fmt.Print("[")
		// for ii, wb := range wordBase {
		// 	fmt.Print("[")
		// 	for index, w := range wb {
		// 		fmt.Printf("'%s'", w)
		// 		if index < (len(wb) - 1) {
		// 			fmt.Print(",")
		// 		}
		// 	}
		// 	fmt.Print("]")
		// 	if ii < (len(wordBase) - 1) {
		// 		fmt.Print(",")
		// 	}
		// }
		// fmt.Print("]")
		// fmt.Print("\n")
		// for _, wb := range wordBase {
		// 	for index, w := range wb {
		// 		fmt.Printf("%s", w)
		// 		if index < (len(wb) - 1) {
		// 			fmt.Print(",")
		// 		}
		// 	}
		// 	fmt.Print("\n")

		// }

		wordCount := WordCount(wordBase)
		headElems, headAddr := BuildHeadElems(wordCount, config.MaxLength)
		filteredWordBase := FilterWordBase(headAddr, wordBase)
		root := &FPRoot{}
		root.BuildFPTree(filteredWordBase, headAddr)

		root.ConditionalPattern(headElems, config.MaxLength, headAddr, 10)

		coWords := WordConcurrence(headAddr, 0)

		sort.Slice(coWords, func(i, j int) bool {
			return coWords[i].SupportCount > coWords[j].SupportCount
		})
		// data := FreqItemsToStrings(coWords)
		// for _, d := range data {
		// 	fmt.Println(d)
		// }
		rMap1, rMap2 := make(map[string]int), make(map[string]int)
		for _, co := range coWords {
			if co.Confidence >= config.MinConfidence {
				result[co.BaseWord] = 1
				result[co.Word] = 1
				// fmt.Println(co)
				continue
			}
			rMap1[co.BaseWord] = 0
			rMap2[co.Word] = 0
			// if co.Confidence >= 0.5 {
			// 	if _, ok := result[co.BaseWord]; !ok {
			// 		result[co.BaseWord] = 2
			// 	}
			// 	if _, ok := result[co.Word]; !ok {
			// 		result[co.Word] = 2
			// 	}
			// 	continue
			// }
			// if co.Confidence >= 0.1 {
			// 	if _, ok := result[co.BaseWord]; !ok {
			// 		result[co.BaseWord] = 1
			// 	}
			// 	if _, ok := result[co.Word]; !ok {
			// 		result[co.Word] = 1
			// 	}
			// 	continue
			// }
		}
		high, normal, low, other := make([]string, 0), make([]string, 0), make([]string, 0), make([]string, 0)

		for k, v := range result {
			tmp := strings.Split(k, " ")
			if v == 3 {
				high = append(high, k)
			}
			if v == 2 {
				normal = append(normal, k)
			}
			if v == 1 {
				low = append(low, tmp[len(tmp)-1])
			}
			if v == 0 {
				other = append(other, tmp[len(tmp)-1])
			}
		}

		r1, r2, r3 := make([]string, 0), make([]string, 0), make([]string, 0)
		for k, _ := range rMap1 {
			if _, ok := rMap2[k]; ok {
				r3 = append(r3, k)
				delete(rMap2, k)
			} else {
				r1 = append(r1, k)
			}
		}
		for k, _ := range rMap2 {
			r2 = append(r2, k)
		}
		sort.Strings(r1)
		sort.Strings(r2)
		sort.Strings(r3)
		fmt.Println("--------------")
		fmt.Println("Aggregate:")
		for _, k := range r1 {
			fmt.Println(k)
		}
		fmt.Println("--------------")
		fmt.Println("Aggregate:")
		for _, k := range r2 {
			fmt.Println(k)
		}
		fmt.Println("--------------")
		fmt.Println("Confuse:")
		for _, k := range r3 {
			fmt.Println(k)
		}
		fmt.Println("--------------")
		// fmt.Println(high)
		// fmt.Println(normal)

		fmt.Println("Confidence:")
		sort.Strings(low)
		for _, k := range low {
			fmt.Println(k)
		}
		fmt.Println("--------------")
		fmt.Println("Others:")
		sort.Strings(other)
		for _, k := range other {
			fmt.Println(k)
		}
	},
}

func init() {
	rootCmd.AddCommand(aggregateCmd)

	aggregateCmd.Flags().StringP("source", "s", "", "source code directory")
	aggregateCmd.Flags().StringP("config", "c", "", "config file")
}
