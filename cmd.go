package main

import (
	"log"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	svp *viper.Viper // server viper
	rvp *viper.Viper // rule viper
)

var root = &cobra.Command{
	Use:   "",
	Short: "定时采集、搜索和推送热点新闻",
}

func RunCommand() {
	var ServeCmd = &cobra.Command{
		Use:   "serve",
		Short: "新闻数据的可视化、搜索和推送服务",
		Long:  "",
		Run: func(cmd *cobra.Command, args []string) {
			LoadSettings(cmd, args)
			RunServer()
		},
	}
	ServeCmd.Flags().StringP("settings_file", "s", "settings", "应用配置文件")

	var SchedulelCmd = &cobra.Command{
		Use:   "schedule",
		Short: "调度爬虫采集任务",
		Long:  "",
		Run: func(cmd *cobra.Command, args []string) {
			LoadSettings(cmd, args)
			LoadRules(cmd, args)
			RunScheduler()
		},
	}
	SchedulelCmd.Flags().StringP("settings_file", "s", "settings", "应用配置文件")
	SchedulelCmd.Flags().StringP("rules_file", "r", "rules", "解析规则文件")

	root.AddCommand(ServeCmd, SchedulelCmd)
	root.Execute()
}

func LoadSettings(cmd *cobra.Command, args []string) {
	var err error
	svp = viper.New()
	svp.AddConfigPath("config")
	fn, err := cmd.Flags().GetString("settings_file")
	if err != nil {
		log.Fatalf("cmd.Flags().GetString err: %v\n", err)
	}
	svp.SetConfigName(fn)
	err = svp.ReadInConfig()
	if err != nil {
		log.Fatalln("svp.ReadInConfig() err: ", err)
	}
	SetupSettings(svp)
}

func LoadRules(cmd *cobra.Command, args []string) {
	var err error
	rvp = viper.New()
	rvp.AddConfigPath("config")
	fn, err := cmd.Flags().GetString("rules_file")
	if err != nil {
		log.Fatalf("cmd.Flags().GetString err: %v\n", err)
	}
	rvp.SetConfigName(fn)
	err = rvp.ReadInConfig()
	if err != nil {
		log.Fatalln("rvp.ReadInConfig() err: ", err)
	}
}

func SetupSettings(svp *viper.Viper) {
	var err error
	err = svp.UnmarshalKey("MySQL", &dbSettings)
	if err != nil {
		log.Fatalln("svp.UnmarshalKey err: ", err)
	}
	err = svp.UnmarshalKey("Server", &serverSettings)
	if err != nil {
		log.Fatalln("svp.UnmarshalKey err: ", err)
	}
	err = svp.UnmarshalKey("Elastic", &elasticSettings)
	if err != nil {
		log.Fatalln("svp.UnmarshalKey err: ", err)
	}
}
