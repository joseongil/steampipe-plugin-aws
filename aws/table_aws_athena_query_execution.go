package aws

import (
	"context"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/service/athena"
	"github.com/turbot/steampipe-plugin-sdk/v3/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v3/plugin"
)

//// TABLE DEFINITION

func tableAwsAthenaQueryExecution(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "aws_athena_query_execution",
		Description: "AWS Athena Query Execution history",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("query_id"),
			Hydrate:    getAwsAthenaQueryExecution,
		},
		List: &plugin.ListConfig{
			// Hydrate: listAwsAthenaQueryExecutions,
			KeyColumns: []*plugin.KeyColumn{
				{Name: "query_id", Require: plugin.Required},
			},
			Hydrate: getAwsAthenaQueryExecution,
		},
		GetMatrixItem: BuildRegionList,
		// Columns: awsRegionalColumns([]*plugin.Column{
		Columns: awsDefaultColumns([]*plugin.Column{
			{
				Name:        "query_id",
				Description: "The id of the query execution.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "query",
				Description: "The statement of the query",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "catalog",
				Description: "The catalog name query executed on",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "database",
				Description: "The database name query executed on",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "workgroup",
				Description: "The workgroup name query executed",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "result_config_output_location",
				Description: "The configuration of output location",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "result_config_expected_bucket_owner",
				Description: "The configuration of expected bucket owner",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "statement_type",
				Description: "The statement type of query",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "statistics_datamenifest_location",
				Description: "The statistics of datamenifest location",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "statistics_data_scanned_byte",
				Description: "The statistics of DataScannedInBytes",
				Type:        proto.ColumnType_INT,
			},
			{
				Name:        "statistics_engine_executiontime_ms",
				Description: "The statistics of EngineExecutionTimeInMillis",
				Type:        proto.ColumnType_INT,
			},
			{
				Name:        "statistics_query_planningtime_ms",
				Description: "The statistics of QueryPlanningTimeInMillis",
				Type:        proto.ColumnType_INT,
			},
			{
				Name:        "statistics_query_queuetime_ms",
				Description: "The statistics of QueryQueueTimeInMillis",
				Type:        proto.ColumnType_INT,
			},
			{
				Name:        "statistics_service_processingtime_ms",
				Description: "The statistics of ServiceProcessingTimeInMillis",
				Type:        proto.ColumnType_INT,
			},
			{
				Name:        "statistics_total_executiontime_ms",
				Description: "The statistics of TotalExecutionTimeInMillis",
				Type:        proto.ColumnType_INT,
			},
			{
				Name:        "athena_error_message",
				Description: "The error message of query execution",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "athena_error_categroy",
				Description: "The error categroy of query execution",
				Type:        proto.ColumnType_INT,
			},
			{
				Name:        "athena_error_type",
				Description: "The error type of query execution",
				Type:        proto.ColumnType_INT,
			},
			{
				Name:        "state",
				Description: "The state of query execution",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "state_change_reason",
				Description: "The reason for state change",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "submission_time",
				Description: "The time at submission of query",
				Type:        proto.ColumnType_TIMESTAMP,
			},
			{
				Name:        "completion_time",
				Description: "The time at completion of query",
				Type:        proto.ColumnType_TIMESTAMP,
			},
		}),
	}
}

type AthenaQueryExecution struct {
	QueryId                           string
	Query                             string
	Catalog                           string
	Database                          string
	Workgroup                         string
	ResultConfigOutputLocation        string
	ResultConfigExpectedBucketOwner   string
	StatementType                     string
	StatisticsDatamenifestLocation    string
	StatisticsDataScannedByte         int64
	StatisticsEngineExecutiontimeMs   int64
	StatisticsQueryPlanningtimeMs     int64
	StatisticsQueryQueuetimeMs        int64
	StatisticsServiceProcessingtimeMs int64
	StatisticsTotalExecutiontimeMs    int64
	AthenaErrorMessage                string
	AthenaErrorCategoroy              int32
	AthenaErrorType                   int32
	State                             string
	StateChangeReason                 string
	SubmissionTime                    time.Time
	CompletionTime                    time.Time
}

func MakeQueryExecutionRow(qe *athena.GetQueryExecutionOutput) AthenaQueryExecution {
	var aqes AthenaQueryExecution

	aqes.QueryId = *qe.QueryExecution.QueryExecutionId
	if qe.QueryExecution.Query != nil {
		aqes.Query = *qe.QueryExecution.Query
	}
	if qe.QueryExecution.QueryExecutionContext.Catalog != nil {
		aqes.Catalog = *qe.QueryExecution.QueryExecutionContext.Catalog
	}
	if qe.QueryExecution.QueryExecutionContext.Database != nil {
		aqes.Database = *qe.QueryExecution.QueryExecutionContext.Database
	}
	if qe.QueryExecution.WorkGroup != nil {
		aqes.Workgroup = *qe.QueryExecution.WorkGroup
	}

	if qe.QueryExecution.ResultConfiguration.OutputLocation != nil {
		aqes.ResultConfigOutputLocation = *qe.QueryExecution.ResultConfiguration.OutputLocation
	}
	if qe.QueryExecution.ResultConfiguration.ExpectedBucketOwner != nil {
		aqes.ResultConfigExpectedBucketOwner = *qe.QueryExecution.ResultConfiguration.ExpectedBucketOwner
	}
	if qe.QueryExecution.StatementType != nil {
		aqes.StatementType = *qe.QueryExecution.StatementType
	}
	if qe.QueryExecution.Statistics.DataManifestLocation != nil {
		aqes.StatisticsDatamenifestLocation = *qe.QueryExecution.Statistics.DataManifestLocation
	}
	if qe.QueryExecution.Statistics.DataScannedInBytes != nil {
		aqes.StatisticsDataScannedByte = *qe.QueryExecution.Statistics.DataScannedInBytes
	}
	if qe.QueryExecution.Statistics.EngineExecutionTimeInMillis != nil {
		aqes.StatisticsEngineExecutiontimeMs = *qe.QueryExecution.Statistics.EngineExecutionTimeInMillis
	}
	if qe.QueryExecution.Statistics.QueryPlanningTimeInMillis != nil {
		aqes.StatisticsQueryPlanningtimeMs = *qe.QueryExecution.Statistics.QueryPlanningTimeInMillis
	}
	if qe.QueryExecution.Statistics.QueryQueueTimeInMillis != nil {
		aqes.StatisticsQueryQueuetimeMs = *qe.QueryExecution.Statistics.QueryQueueTimeInMillis
	}
	if qe.QueryExecution.Statistics.ServiceProcessingTimeInMillis != nil {
		aqes.StatisticsServiceProcessingtimeMs = *qe.QueryExecution.Statistics.ServiceProcessingTimeInMillis
	}
	if qe.QueryExecution.Statistics.TotalExecutionTimeInMillis != nil {
		aqes.StatisticsTotalExecutiontimeMs = *qe.QueryExecution.Statistics.TotalExecutionTimeInMillis
	}

	if qe.QueryExecution.Status.State != nil {
		aqes.State = *qe.QueryExecution.Status.State
	}
	if qe.QueryExecution.Status.StateChangeReason != nil {
		aqes.StateChangeReason = *qe.QueryExecution.Status.StateChangeReason
	}

	if qe.QueryExecution.Status.SubmissionDateTime != nil {
		aqes.SubmissionTime = *qe.QueryExecution.Status.SubmissionDateTime
	}
	if qe.QueryExecution.Status.CompletionDateTime != nil {
		aqes.CompletionTime = *qe.QueryExecution.Status.CompletionDateTime
	}

	return aqes
}

//// LIST FUNCTION

func listAwsAthenaQueryExecutions(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	// Create Session
	svc, err := AthenaService(ctx, d)
	if err != nil {
		return nil, err
	}
	limit := aws.Int64(50)
	if d.QueryContext.Limit != nil {
		if *d.QueryContext.Limit < 5 {
			limit = aws.Int64(5)
		} else if *d.QueryContext.Limit < 50 {
			limit = d.QueryContext.Limit
		}
	}
	dummy := "Initial"
	nextToken := &dummy
	input := &athena.ListQueryExecutionsInput{
		MaxResults: limit,
	}
	for nextToken != nil {
		if dummy != "Initial" {
			input = &athena.ListQueryExecutionsInput{
				MaxResults: limit,
				NextToken:  nextToken,
			}
		}
		queryResult, err := svc.ListQueryExecutions(input)

		if err != nil {
			plugin.Logger(ctx).Error("aws_athena_query_execution.listQueryExecution", "api_error", err)
			return nil, err
		}
		nextToken = queryResult.NextToken
		for _, query_id := range queryResult.QueryExecutionIds {
			params := &athena.GetQueryExecutionInput{
				QueryExecutionId: query_id,
			}
			qe, err := svc.GetQueryExecution(params)
			if err != nil {
				plugin.Logger(ctx).Error("GetQueryExecution Error:", err)
				continue
			}
			var aqes AthenaQueryExecution
			if qe != nil {
				aqes = MakeQueryExecutionRow(qe)
			}
			d.StreamListItem(ctx, aqes)
			// Context may get cancelled due to manual cancellation or if the limit has been reached
			if d.QueryStatus.RowsRemaining(ctx) == 0 {
				return nil, nil
			}
		}
	}

	return nil, nil
}

func getAwsAthenaQueryExecution(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	plugin.Logger(ctx).Trace("getAwsAthenaQueryExecution")

	query_id := d.KeyColumnQuals["query_id"].GetStringValue()

	// create service
	svc, err := AthenaService(ctx, d)
	if err != nil {
		return nil, err
	}

	// Build the params
	params := &athena.GetQueryExecutionInput{
		QueryExecutionId: &query_id,
	}

	// Get call
	qe, err := svc.GetQueryExecution(params)
	if err != nil {
		if a, ok := err.(awserr.Error); ok {
			if a.Code() == "InvalidRequestException" {
				return nil, nil
			}
		}
		plugin.Logger(ctx).Error("aws_athena_query_execution.getQueryExecution", "api_error", err)
		return nil, err
	}

	// plugin.Logger(ctx).Error("@DBG3-0:")
	// __jsonb, err := json.Marshal(qe)
	// plugin.Logger(ctx).Error("@DBG3-0-1:" + string(__jsonb) + "\n")
	// if err != nil {
	// 	return nil, nil
	// }

	var aqes AthenaQueryExecution
	if qe != nil {
		aqes = MakeQueryExecutionRow(qe)
		return aqes, nil
	}

	return nil, nil
}
