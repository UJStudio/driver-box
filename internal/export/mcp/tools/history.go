package tools

import (
	"context"
	"fmt"
	"github.com/ibuilding-x/driver-box/driverbox/export/history"
	"github.com/ibuilding-x/driver-box/driverbox/helper"
	"github.com/mark3labs/mcp-go/mcp"
	"go.uber.org/zap"
)

var HistoryTableSchemaTool = mcp.NewTool("history_table_schema",
	mcp.WithDescription("查询当前网关数据库的表结构定义,有助于大模型开展数据分析"),
)

var HistoryTableSchemaHandler = func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	r, e := history.NewExport().QueryDataBySql("SELECT name,sql FROM sqlite_master WHERE type='table';")
	helper.Logger.Info("查询当前网关数据库的表结构定义", zap.Any("result", r), zap.Error(e))
	if e != nil {
		return nil, e
	}
	markdown := fmt.Sprintf("## 表结构定义（共 %d 张表）\n\n", len(r))
	for _, v := range r {
		markdown += fmt.Sprintf("### tableName: %s\n\n```sql\n%s\n```\n\n", v["name"], v["sql"])
	}
	return mcp.NewToolResultText(markdown), nil
}

var HistoryDataAnalysisTool = mcp.NewTool("history_data_analysis",
	mcp.WithDescription("执行大模型生成的SQL查询语句。要求查询SQL必须是网关中存在的表和字段，且符合 sqlite 语法。当查询设备数据时，需确保查询条件中的点位名同物模型定义保持一致。"),
	mcp.WithString("sql", mcp.Required(), mcp.Description("要执行的SQL查询语句")),
)

var HistoryDataAnalysisHandler = func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	sql := request.GetString("sql", "")
	if sql == "" {
		return nil, fmt.Errorf("请提供有效的SQL查询语句")
	}
	helper.Logger.Info("执行大模型生成的SQL查询语句", zap.String("sql", sql))
	r, e := history.NewExport().QueryDataBySql(sql)

	if e != nil {
		return nil, e
	}
	if len(r) == 0 {
		return mcp.NewToolResultText("无结果"), nil
	}
	markdown := fmt.Sprintf("## 查询结果（共 %d 条记录）\n\n", len(r))
	markdown += "|"
	for k, _ := range r[0] {
		markdown += fmt.Sprintf("%s |", k)
	}
	markdown += "\n|"
	for _, _ = range r[0] {
		markdown += fmt.Sprintf("---|")
	}
	markdown += "\n"
	for _, v := range r {
		markdown += "|"
		for k, _ := range v {
			markdown += fmt.Sprintf("%s |", v[k])
		}
		markdown += "\n"
	}
	helper.Logger.Info("执行大模型生成的SQL查询语句", zap.Any("result", r), zap.Error(e))
	return mcp.NewToolResultText(markdown), nil
}
