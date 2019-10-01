package main

import (
	"fmt"
	"strings"

	"github.com/Tnze/CoolQ-Golang-SDK/cqp"
)

var (
	group    []int64
	lyricsNO []int
)

//go:generate cqcfg -c .
// cqp: 名称: Creeper
// cqp: 版本: 1.0.0:0
// cqp: 作者: BaiMeow
// cqp: 简介: 简易Creeper接龙
func main() { /*此处应当留空*/ }

func init() {
	cqp.AppID = "cn.miaoscraft.creeper" // TODO: 修改为这个插件的ID
	cqp.GroupMsg = onGroupMsg
}

func onGroupMsg(subType, msgID int32, fromGroup, fromQQ int64, fromAnonymous, msg string, font int32) int32 {
	if latter := getnext(strings.ToLower(msg), fromGroup); latter != "" {
		cqp.SendGroupMsg(fromGroup, latter)
	}
	return 0
}

func getnext(former string, fromGroup int64) string {
	var i = 0
	//判断群是否有接龙记录
	if len(group) != 0 {
		for i = 0; i < len(group); i++ {
			if group[i] == fromGroup {
				break
			}
			i++
		}
		if i == len(group) {
			group = append(group, fromGroup)
			lyricsNO = append(lyricsNO, 0)
		}
	} else {
		group = append(group, fromGroup)
		lyricsNO = append(lyricsNO, 0)
	}
	cqp.AddLog(cqp.Info, "Creeper", fmt.Sprintf("i=%d", i))
	//接不上来的时候可以复读让机器人接下一句
	if former == strings.ToLower(lyrics[lyricsNO[i]]) {
		cqp.AddLog(cqp.Info, "Creeper", "群员复读")
		lyricsNO[i]++
		cqp.AddLog(cqp.Info, "Creeper", lyrics[lyricsNO[i]])
		if lyricsNO[i] == len(lyrics)-1 {
			defer func() {
				lyricsNO[i] = 0
				i = 0
			}()
		}
		return lyrics[lyricsNO[i]]
	}
	//正常接龙
	if former == strings.ToLower(lyrics[lyricsNO[i]+1]) {
		cqp.AddLog(cqp.Info, "Creeper", "群员接龙")
		lyricsNO[i] += 2
		cqp.AddLog(cqp.Info, "Creeper", lyrics[lyricsNO[i]])
		//群友接龙到末尾，重新初始化到第一句
		if lyricsNO[i] >= len(lyrics) {
			lyricsNO[i] = 0
			return ""
		}
		//机器人接龙到最后一个时重新初始化到第一句
		if lyricsNO[i] == len(lyrics)-1 {
			defer func() {
				lyricsNO[i] = 0
			}()
		}
		return lyrics[lyricsNO[i]]
	}
	//重置接龙
	if former == strings.ToLower(lyrics[0]) {
		cqp.AddLog(cqp.Info, "Creeper", "重置接龙")
		lyricsNO[i] = 1
		return lyrics[1]
	}
	return ""
}

var lyrics = []string{"Creeper?",
	"Awww man!",
	"So way back in the mine",
	"Got our pickaxe swinging from side to side",
	"Side,side to side...",
	"This task's a grueling one , hope to find some diamonds tonight,night,night",
	"Diamonds tonight...",
	"Heads up",
	"You hear a sound , turn around and look up",
	"Total shock fills your body",
	"Oh no it's you again , I could never forget those eyes,eyes,eyes",
	"Eyes,eyes,eyes...",
	"Cause baby tonight , the creeper's trying to steal all our stuff again",
	"Cause baby tonight , you grab your pick , shovel and bolt again , bolt again(gain)",
	"And run,run until it's done,done , until the sun comes up in the morn'",
	"Cause baby tonight , the creeper's trying to steal all our stuff again , stuff again(gain)",
	"Just when you think you're safe , overhear some hissing from , right behind",
	"Right,right behind...",
	"That's a nice life you have ,shame it's gotta end at this time,time,time",
	"Time,time,time , time...",
	"Blows up , then your health bar drops and you could use a 1-up",
	"Get inside , don't be tardy",
	"So now you are stuck in there , half a heart is left but don't die,die,die",
	"Die,die,die...",
	"Cause baby tonight , the creeper's trying to steal all our stuff again",
	"Cause baby tonight , you grab your pick , shovel and bolt again,bolt again(gain)",
	"And run,run until it's done,done , until the sun comes up in the morn'",
	"Cause baby tonight , the creeper's trying to steal all our stuff again",
	"Creepers , you're mine , haha",
	"Dig up diamonds , and craft those diamonds and make some armour , get it baby",
	"Go and forge that like you so , MLG pro , the sword's made of diamonds , so come at me bro",
	"Huh",
	"Training in your room under the torchlight , hone that form to get you ready for the big fight",
	"Every single day and the whole night , Creeper's out prowlin' , hmm - alright",
	"Look at me , look at you , take my revenge that's what I'm gonna do",
	"I'm a - warrior baby , what else is new ? And my blade's gonna tear (Cause baby tonight,) through you , bring it",
	"The creeper's trying to steal our stuff again",
	"Yeah , let's take back the world",
	"Hahhah",
	"Grab your sword , armour and go",
	"It's on",
	"Take your revenge (\"Woooo\")",
	"Ahhoahh",
	"So fight,fight like it's the last,last night of your life,life (ahhahah) show them your bite",
	"Wooo",
	"Cause baby tonight , (\"Ahhhhahaahaaah\") the creeper's trying to steal all our stuff again",
	"Cause baby tonight , you grab your pick , shovel and bolt again,bolt again(gain) (\"wooo\")",
	"And run,run until it's done,done , until the sun comes up in the morn'",
	"Cause baby tonight , the creeper's trying to steal all our stuff again ",
	"Woooo"}
