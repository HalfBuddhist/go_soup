package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type edge struct {
	Dest   string
	Weight uint
}

type node struct {
	ID      string
	Type    int
	Connect []edge
}

type nodeArray []node

func contains(s string, slice []string) bool {
	for _, item := range slice {
		if item == s {
			return true
		}
	}
	return false
}

var nodesLid []string

func getSwsNodesInfo(filePath string) ([][]string, []string, map[string]string) {

	var switchLinkInfoList [][]string
	var swsLid []string
	nodesLidName := make(map[string]string)
	var currentSwitchInfo []string
	//var nodesLid []string
	var nodesName []string

	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("打开文件时出错:", err)
		return nil, nil, nil
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		line = strings.TrimSpace(line)
		nodesLidRegex := regexp.MustCompile(`^([0-9a-z])+\s+(\d+)\s+(\d+)\[\s+\]\s+==\(\s+(\d+\S)+\s+((-?\d+)(\.\d+)?\s+[A-Za-z]+)\s+([A-Za-z]+\/+\s+[A-Za-z]+)\)==>+\s+(\d+)\s+(\d+)\[\s+]\s+("([^"]+)")\s+\((.*?)\)$`)
		nodesNameRegex := regexp.MustCompile(`^CA:+\s+(\S+)\s(\S+):`)
		swsLidRegex := regexp.MustCompile(`^(\d+)\s+(\d+)\[\s+\]\s+==\(\s+(\d+\S)+\s+((-?\d+)(\.\d+)?\s+[A-Za-z]+)\s+([A-Za-z]+\/+\s+[A-Za-z]+)\)==>+\s+(\d+)\s+(\d+)\[\s+]\s+("([^"]+)")\s+\((.*?)\)$`)

		nodesLidMatch := nodesLidRegex.FindString(line)
		nodesNameMatch := nodesNameRegex.FindString(line)
		swsLidMatch := swsLidRegex.FindString(line)

		if nodesLidMatch != "" {
			nodesLidMatches := nodesLidRegex.FindStringSubmatch(line)
			if nodesLidMatches != nil {
				nodesLid = append(nodesLid, nodesLidMatches[2])
			}

		}

		if swsLidMatch != "" {
			swsLidMatches := swsLidRegex.FindStringSubmatch(line)

			if !contains(swsLidMatches[1], swsLid) {
				swsLid = append(swsLid, swsLidMatches[1])
			}

		}
		if strings.HasPrefix(line, "CA:") {
			if nodesNameMatch != "" {
				nodesNameMatches := nodesNameRegex.FindStringSubmatch(line)
				nodesName = append(nodesName, nodesNameMatches[1])
			} else {
				nodesName = append(nodesName, "")
			}
		}

		if strings.HasPrefix(line, "Switch:") {
			if currentSwitchInfo != nil {
				switchLinkInfoList = append(switchLinkInfoList, currentSwitchInfo)
				currentSwitchInfo = make([]string, 0)

			}
		}
		if swsLidMatch != "" {
			currentSwitchInfo = append(currentSwitchInfo, line)
		}

	}
	if currentSwitchInfo != nil {
		switchLinkInfoList = append(switchLinkInfoList, currentSwitchInfo)

	}
	for i, value := range nodesName {

		nodesLidName[nodesLid[i]] = value
	}

	return switchLinkInfoList, swsLid, nodesLidName
}

func getSwsNodesLink(switchLinkInfoList [][]string, swsLid []string, nodesLidName map[string]string) (nodeArray, nodeArray, [][]string, int) {
	switchNum := 0
	var allSwConnect nodeArray
	var allNodesConnect nodeArray
	switchAllNodes := make([][]string, len(switchLinkInfoList))
	for index, switchInfo := range switchLinkInfoList {
		var Connects []edge
		var alreadyCount []string

		for _, line := range switchInfo {
			regex := regexp.MustCompile(`^(\d+)\s+(\d+)\[\s+\]\s+==\(\s+(\d+\S)+\s+((-?\d+)(\.\d+)?\s+[A-Za-z]+)\s+([A-Za-z]+\/+\s+[A-Za-z]+)\)==>+\s+(\d+)\s+(\d+)\[\s+]\s+("([^"]+)")\s+\((.*?)\)$`)

			match := regex.FindStringSubmatch(line)
			if match != nil {
				sourceNode := match[1]
				targetNode := match[8]
				//sourcePort := match[2]
				//linkWidth := match[3]
				//linkSpeed := match[4]
				//linkStatus := match[7]
				//targetPort := match[9]
				//nodeName := match[10]
				//fmt.Println("Source Node:", sourceNode)
				//fmt.Println("Source Port:", sourcePort)
				//fmt.Println("Link Width:", linkWidth)
				//fmt.Println("Link NodeSpeed:", linkSpeed)
				//fmt.Println("Link Status:", linkStatus)
				//fmt.Println("Target Node:", targetNode)
				//fmt.Println("Target Port:", targetPort)
				//fmt.Println("Node Name:", nodeName)

				// switch nodes

				Connect := edge{
					Dest:   targetNode,
					Weight: 1,
				}

				//for key, _ := range nodesLidName {
				//	nodesLid = append(nodesLid, key)
				//}
				if !contains(targetNode, alreadyCount) && contains(targetNode, nodesLid) {
					if nodesLidName[targetNode] != "" {
						Connects = append(Connects, Connect)
						alreadyCount = append(alreadyCount, targetNode)
						switchAllNodes[switchNum] = append(switchAllNodes[switchNum], targetNode)

					}

				} else if !contains(targetNode, alreadyCount) && contains(targetNode, swsLid) {
					Connects = append(Connects, Connect)
					alreadyCount = append(alreadyCount, targetNode)
					switchAllNodes[switchNum] = append(switchAllNodes[switchNum], targetNode)
				}

				if !contains(targetNode, swsLid) {
					if nodesLidName[targetNode] != "" {
						NodesConnect := edge{
							Dest:   sourceNode,
							Weight: 1,
						}

						NodesLinks := node{
							ID:      targetNode,
							Type:    0,
							Connect: []edge{NodesConnect},
						}

						allNodesConnect = append(allNodesConnect, NodesLinks)

					}
				}

			}
		}
		//fmt.Println("Connects", Connects)
		linkStatus := node{
			ID:      swsLid[index],
			Type:    1,
			Connect: Connects,
		}
		switchNum += 1
		allSwConnect = append(allSwConnect, linkStatus)

	}
	return allSwConnect, allNodesConnect, switchAllNodes, switchNum

}

func getNetLevel(switchNum int, switchAllNodes [][]string, swIndex []string) (map[string]string, map[string]int) {
	spine0Index := 0
	leaf0Index := 0
	isSwSpine := make(map[string]string)
	levelSw := make(map[string]int)
	spineNums := 0
	leafNum := 0

	for i := 0; i < switchNum; i++ {
		flag := 0

		for _, swNodes := range switchAllNodes[i] {
			if contains(swNodes, nodesLid) {
				flag += 1

			}
		}
		if flag == 0 {
			if spineNums == 0 {
				isSwSpine[swIndex[i]] = "spine"
				spine0Index = i
				levelSw[swIndex[i]] = 2

			} else {
				isSwSpine[swIndex[spine0Index]] = "spine0"
				isSwSpine[swIndex[i]] = "spine" + strconv.Itoa(spineNums)
				levelSw[swIndex[i]] = 2
			}
			spineNums += 1

		} else {
			if leafNum == 0 {
				isSwSpine[swIndex[i]] = "leaf"
				leaf0Index = i
				levelSw[swIndex[i]] = 1
			} else {
				isSwSpine[swIndex[leaf0Index]] = "leaf0"
				isSwSpine[swIndex[i]] = "leaf" + strconv.Itoa(leafNum)
				levelSw[swIndex[i]] = 1
			}
			leafNum += 1

		}
	}
	return isSwSpine, levelSw

}

func replaceSwsName(swType map[string]string, levelSw map[string]int, allSwConnect nodeArray, allNodesConnect nodeArray, nodesLidName map[string]string) (nodeArray, nodeArray) {
	var swTypeLid []string
	for key, _ := range swType {
		swTypeLid = append(swTypeLid, key)
	}

	for i, swConnect := range allSwConnect {

		for j, n := range swConnect.Connect {
			if contains(n.Dest, swTypeLid) {
				DestSw := n.Dest
				n.Dest = swType[DestSw]

			} else {
				DestNode := n.Dest
				n.Dest = nodesLidName[DestNode]
			}
			swConnect.Connect[j] = n

		}
		swId := allSwConnect[i].ID
		allSwConnect[i].ID = swType[swId]
		allSwConnect[i].Type = levelSw[swId]

	}
	for m, nodeConnect := range allNodesConnect {
		nodeId := allNodesConnect[m].ID
		allNodesConnect[m].ID = nodesLidName[nodeId]
		for n, sw := range nodeConnect.Connect {
			nodeDest := sw.Dest
			if contains(nodeDest, swTypeLid) {
				sw.Dest = swType[nodeDest]

			} else {
				sw.Dest = nodesLidName[nodeDest]
			}
			nodeConnect.Connect[n] = sw

		}
	}
	return allSwConnect, allNodesConnect
}

func main() {
	filePath := "iblinkinfo.log"

	switchLinkInfoList, swsLid, nodesLidName := getSwsNodesInfo(filePath)

	allSwConnect, allNodesConnect, switchAllNodes, switchNums := getSwsNodesLink(switchLinkInfoList, swsLid, nodesLidName)

	swType, levelSw := getNetLevel(switchNums, switchAllNodes, swsLid)

	allSwConnect, allNodesConnect = replaceSwsName(swType, levelSw, allSwConnect, allNodesConnect, nodesLidName)
	allDevicesConnect := append(allSwConnect, allNodesConnect...)

	jsonData, err := json.MarshalIndent(allDevicesConnect, "", "    ")
	if err != nil {
		fmt.Println("JSON编码错误:", err)
		return
	}

	// 打开文件进行写入（如果文件不存在会创建，如果存在会覆盖）
	file, err := os.Create("IB_Top.json")
	if err != nil {
		fmt.Println("创建文件错误:", err)
		return
	}
	defer file.Close()

	// 将JSON数据写入文件
	_, err = file.Write(jsonData)
	if err != nil {
		fmt.Println("写入文件错误:", err)
		return
	}

	fmt.Println("拓扑解析数据已成功写入到 IB_Top.json 文件.")
}
