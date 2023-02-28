// Code generated by mdatagen. DO NOT EDIT.

package metadata

import (
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/confmap/confmaptest"
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/pmetric"
	"go.opentelemetry.io/collector/receiver/receivertest"
	"go.uber.org/zap"
	"go.uber.org/zap/zaptest/observer"
)

type testConfigCollection int

const (
	testSetDefault testConfigCollection = iota
	testSetAll
	testSetNone
)

func TestMetricsBuilder(t *testing.T) {
	tests := []struct {
		name      string
		configSet testConfigCollection
	}{
		{
			name:      "default",
			configSet: testSetDefault,
		},
		{
			name:      "all_set",
			configSet: testSetAll,
		},
		{
			name:      "none_set",
			configSet: testSetNone,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			start := pcommon.Timestamp(1_000_000_000)
			ts := pcommon.Timestamp(1_000_001_000)
			observedZapCore, observedLogs := observer.New(zap.WarnLevel)
			settings := receivertest.NewNopCreateSettings()
			settings.Logger = zap.New(observedZapCore)
			mb := NewMetricsBuilder(loadConfig(t, test.name), settings, WithStartTime(start))

			expectedWarnings := 0
			assert.Equal(t, expectedWarnings, observedLogs.Len())

			defaultMetricsCount := 0
			allMetricsCount := 0

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordRedisClientsBlockedDataPoint(ts, 1)

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordRedisClientsConnectedDataPoint(ts, 1)

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordRedisClientsMaxInputBufferDataPoint(ts, 1)

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordRedisClientsMaxOutputBufferDataPoint(ts, 1)

			allMetricsCount++
			mb.RecordRedisCmdCallsDataPoint(ts, 1, "attr-val")

			allMetricsCount++
			mb.RecordRedisCmdUsecDataPoint(ts, 1, "attr-val")

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordRedisCommandsDataPoint(ts, 1)

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordRedisCommandsProcessedDataPoint(ts, 1)

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordRedisConnectionsReceivedDataPoint(ts, 1)

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordRedisConnectionsRejectedDataPoint(ts, 1)

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordRedisCPUTimeDataPoint(ts, 1, AttributeState(1))

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordRedisDbAvgTTLDataPoint(ts, 1, "attr-val")

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordRedisDbExpiresDataPoint(ts, 1, "attr-val")

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordRedisDbKeysDataPoint(ts, 1, "attr-val")

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordRedisKeysEvictedDataPoint(ts, 1)

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordRedisKeysExpiredDataPoint(ts, 1)

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordRedisKeyspaceHitsDataPoint(ts, 1)

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordRedisKeyspaceMissesDataPoint(ts, 1)

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordRedisLatestForkDataPoint(ts, 1)

			allMetricsCount++
			mb.RecordRedisMaxmemoryDataPoint(ts, 1)

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordRedisMemoryFragmentationRatioDataPoint(ts, 1)

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordRedisMemoryLuaDataPoint(ts, 1)

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordRedisMemoryPeakDataPoint(ts, 1)

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordRedisMemoryRssDataPoint(ts, 1)

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordRedisMemoryUsedDataPoint(ts, 1)

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordRedisNetInputDataPoint(ts, 1)

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordRedisNetOutputDataPoint(ts, 1)

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordRedisRdbChangesSinceLastSaveDataPoint(ts, 1)

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordRedisReplicationBacklogFirstByteOffsetDataPoint(ts, 1)

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordRedisReplicationOffsetDataPoint(ts, 1)

			allMetricsCount++
			mb.RecordRedisRoleDataPoint(ts, 1, AttributeRole(1))

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordRedisSlavesConnectedDataPoint(ts, 1)

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordRedisUptimeDataPoint(ts, 1)

			metrics := mb.Emit(WithRedisVersion("attr-val"))

			if test.configSet == testSetNone {
				assert.Equal(t, 0, metrics.ResourceMetrics().Len())
				return
			}

			assert.Equal(t, 1, metrics.ResourceMetrics().Len())
			rm := metrics.ResourceMetrics().At(0)
			attrCount := 0
			enabledAttrCount := 0
			attrVal, ok := rm.Resource().Attributes().Get("redis.version")
			attrCount++
			assert.Equal(t, mb.resourceAttributesSettings.RedisVersion.Enabled, ok)
			if mb.resourceAttributesSettings.RedisVersion.Enabled {
				enabledAttrCount++
				assert.EqualValues(t, "attr-val", attrVal.Str())
			}
			assert.Equal(t, enabledAttrCount, rm.Resource().Attributes().Len())
			assert.Equal(t, attrCount, 1)

			assert.Equal(t, 1, rm.ScopeMetrics().Len())
			ms := rm.ScopeMetrics().At(0).Metrics()
			if test.configSet == testSetDefault {
				assert.Equal(t, defaultMetricsCount, ms.Len())
			}
			if test.configSet == testSetAll {
				assert.Equal(t, allMetricsCount, ms.Len())
			}
			validatedMetrics := make(map[string]bool)
			for i := 0; i < ms.Len(); i++ {
				switch ms.At(i).Name() {
				case "redis.clients.blocked":
					assert.False(t, validatedMetrics["redis.clients.blocked"], "Found a duplicate in the metrics slice: redis.clients.blocked")
					validatedMetrics["redis.clients.blocked"] = true
					assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
					assert.Equal(t, "Number of clients pending on a blocking call", ms.At(i).Description())
					assert.Equal(t, "", ms.At(i).Unit())
					assert.Equal(t, false, ms.At(i).Sum().IsMonotonic())
					assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
					dp := ms.At(i).Sum().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
					assert.Equal(t, int64(1), dp.IntValue())
				case "redis.clients.connected":
					assert.False(t, validatedMetrics["redis.clients.connected"], "Found a duplicate in the metrics slice: redis.clients.connected")
					validatedMetrics["redis.clients.connected"] = true
					assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
					assert.Equal(t, "Number of client connections (excluding connections from replicas)", ms.At(i).Description())
					assert.Equal(t, "", ms.At(i).Unit())
					assert.Equal(t, false, ms.At(i).Sum().IsMonotonic())
					assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
					dp := ms.At(i).Sum().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
					assert.Equal(t, int64(1), dp.IntValue())
				case "redis.clients.max_input_buffer":
					assert.False(t, validatedMetrics["redis.clients.max_input_buffer"], "Found a duplicate in the metrics slice: redis.clients.max_input_buffer")
					validatedMetrics["redis.clients.max_input_buffer"] = true
					assert.Equal(t, pmetric.MetricTypeGauge, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Gauge().DataPoints().Len())
					assert.Equal(t, "Biggest input buffer among current client connections", ms.At(i).Description())
					assert.Equal(t, "", ms.At(i).Unit())
					dp := ms.At(i).Gauge().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
					assert.Equal(t, int64(1), dp.IntValue())
				case "redis.clients.max_output_buffer":
					assert.False(t, validatedMetrics["redis.clients.max_output_buffer"], "Found a duplicate in the metrics slice: redis.clients.max_output_buffer")
					validatedMetrics["redis.clients.max_output_buffer"] = true
					assert.Equal(t, pmetric.MetricTypeGauge, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Gauge().DataPoints().Len())
					assert.Equal(t, "Longest output list among current client connections", ms.At(i).Description())
					assert.Equal(t, "", ms.At(i).Unit())
					dp := ms.At(i).Gauge().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
					assert.Equal(t, int64(1), dp.IntValue())
				case "redis.cmd.calls":
					assert.False(t, validatedMetrics["redis.cmd.calls"], "Found a duplicate in the metrics slice: redis.cmd.calls")
					validatedMetrics["redis.cmd.calls"] = true
					assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
					assert.Equal(t, "Total number of calls for a command", ms.At(i).Description())
					assert.Equal(t, "", ms.At(i).Unit())
					assert.Equal(t, true, ms.At(i).Sum().IsMonotonic())
					assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
					dp := ms.At(i).Sum().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
					assert.Equal(t, int64(1), dp.IntValue())
					attrVal, ok := dp.Attributes().Get("cmd")
					assert.True(t, ok)
					assert.EqualValues(t, "attr-val", attrVal.Str())
				case "redis.cmd.usec":
					assert.False(t, validatedMetrics["redis.cmd.usec"], "Found a duplicate in the metrics slice: redis.cmd.usec")
					validatedMetrics["redis.cmd.usec"] = true
					assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
					assert.Equal(t, "Total time for all executions of this command", ms.At(i).Description())
					assert.Equal(t, "us", ms.At(i).Unit())
					assert.Equal(t, true, ms.At(i).Sum().IsMonotonic())
					assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
					dp := ms.At(i).Sum().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
					assert.Equal(t, int64(1), dp.IntValue())
					attrVal, ok := dp.Attributes().Get("cmd")
					assert.True(t, ok)
					assert.EqualValues(t, "attr-val", attrVal.Str())
				case "redis.commands":
					assert.False(t, validatedMetrics["redis.commands"], "Found a duplicate in the metrics slice: redis.commands")
					validatedMetrics["redis.commands"] = true
					assert.Equal(t, pmetric.MetricTypeGauge, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Gauge().DataPoints().Len())
					assert.Equal(t, "Number of commands processed per second", ms.At(i).Description())
					assert.Equal(t, "{ops}/s", ms.At(i).Unit())
					dp := ms.At(i).Gauge().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
					assert.Equal(t, int64(1), dp.IntValue())
				case "redis.commands.processed":
					assert.False(t, validatedMetrics["redis.commands.processed"], "Found a duplicate in the metrics slice: redis.commands.processed")
					validatedMetrics["redis.commands.processed"] = true
					assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
					assert.Equal(t, "Total number of commands processed by the server", ms.At(i).Description())
					assert.Equal(t, "", ms.At(i).Unit())
					assert.Equal(t, true, ms.At(i).Sum().IsMonotonic())
					assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
					dp := ms.At(i).Sum().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
					assert.Equal(t, int64(1), dp.IntValue())
				case "redis.connections.received":
					assert.False(t, validatedMetrics["redis.connections.received"], "Found a duplicate in the metrics slice: redis.connections.received")
					validatedMetrics["redis.connections.received"] = true
					assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
					assert.Equal(t, "Total number of connections accepted by the server", ms.At(i).Description())
					assert.Equal(t, "", ms.At(i).Unit())
					assert.Equal(t, true, ms.At(i).Sum().IsMonotonic())
					assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
					dp := ms.At(i).Sum().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
					assert.Equal(t, int64(1), dp.IntValue())
				case "redis.connections.rejected":
					assert.False(t, validatedMetrics["redis.connections.rejected"], "Found a duplicate in the metrics slice: redis.connections.rejected")
					validatedMetrics["redis.connections.rejected"] = true
					assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
					assert.Equal(t, "Number of connections rejected because of maxclients limit", ms.At(i).Description())
					assert.Equal(t, "", ms.At(i).Unit())
					assert.Equal(t, true, ms.At(i).Sum().IsMonotonic())
					assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
					dp := ms.At(i).Sum().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
					assert.Equal(t, int64(1), dp.IntValue())
				case "redis.cpu.time":
					assert.False(t, validatedMetrics["redis.cpu.time"], "Found a duplicate in the metrics slice: redis.cpu.time")
					validatedMetrics["redis.cpu.time"] = true
					assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
					assert.Equal(t, "System CPU consumed by the Redis server in seconds since server start", ms.At(i).Description())
					assert.Equal(t, "s", ms.At(i).Unit())
					assert.Equal(t, true, ms.At(i).Sum().IsMonotonic())
					assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
					dp := ms.At(i).Sum().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeDouble, dp.ValueType())
					assert.Equal(t, float64(1), dp.DoubleValue())
					attrVal, ok := dp.Attributes().Get("state")
					assert.True(t, ok)
					assert.Equal(t, "sys", attrVal.Str())
				case "redis.db.avg_ttl":
					assert.False(t, validatedMetrics["redis.db.avg_ttl"], "Found a duplicate in the metrics slice: redis.db.avg_ttl")
					validatedMetrics["redis.db.avg_ttl"] = true
					assert.Equal(t, pmetric.MetricTypeGauge, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Gauge().DataPoints().Len())
					assert.Equal(t, "Average keyspace keys TTL", ms.At(i).Description())
					assert.Equal(t, "ms", ms.At(i).Unit())
					dp := ms.At(i).Gauge().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
					assert.Equal(t, int64(1), dp.IntValue())
					attrVal, ok := dp.Attributes().Get("db")
					assert.True(t, ok)
					assert.EqualValues(t, "attr-val", attrVal.Str())
				case "redis.db.expires":
					assert.False(t, validatedMetrics["redis.db.expires"], "Found a duplicate in the metrics slice: redis.db.expires")
					validatedMetrics["redis.db.expires"] = true
					assert.Equal(t, pmetric.MetricTypeGauge, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Gauge().DataPoints().Len())
					assert.Equal(t, "Number of keyspace keys with an expiration", ms.At(i).Description())
					assert.Equal(t, "", ms.At(i).Unit())
					dp := ms.At(i).Gauge().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
					assert.Equal(t, int64(1), dp.IntValue())
					attrVal, ok := dp.Attributes().Get("db")
					assert.True(t, ok)
					assert.EqualValues(t, "attr-val", attrVal.Str())
				case "redis.db.keys":
					assert.False(t, validatedMetrics["redis.db.keys"], "Found a duplicate in the metrics slice: redis.db.keys")
					validatedMetrics["redis.db.keys"] = true
					assert.Equal(t, pmetric.MetricTypeGauge, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Gauge().DataPoints().Len())
					assert.Equal(t, "Number of keyspace keys", ms.At(i).Description())
					assert.Equal(t, "", ms.At(i).Unit())
					dp := ms.At(i).Gauge().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
					assert.Equal(t, int64(1), dp.IntValue())
					attrVal, ok := dp.Attributes().Get("db")
					assert.True(t, ok)
					assert.EqualValues(t, "attr-val", attrVal.Str())
				case "redis.keys.evicted":
					assert.False(t, validatedMetrics["redis.keys.evicted"], "Found a duplicate in the metrics slice: redis.keys.evicted")
					validatedMetrics["redis.keys.evicted"] = true
					assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
					assert.Equal(t, "Number of evicted keys due to maxmemory limit", ms.At(i).Description())
					assert.Equal(t, "", ms.At(i).Unit())
					assert.Equal(t, true, ms.At(i).Sum().IsMonotonic())
					assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
					dp := ms.At(i).Sum().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
					assert.Equal(t, int64(1), dp.IntValue())
				case "redis.keys.expired":
					assert.False(t, validatedMetrics["redis.keys.expired"], "Found a duplicate in the metrics slice: redis.keys.expired")
					validatedMetrics["redis.keys.expired"] = true
					assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
					assert.Equal(t, "Total number of key expiration events", ms.At(i).Description())
					assert.Equal(t, "", ms.At(i).Unit())
					assert.Equal(t, true, ms.At(i).Sum().IsMonotonic())
					assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
					dp := ms.At(i).Sum().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
					assert.Equal(t, int64(1), dp.IntValue())
				case "redis.keyspace.hits":
					assert.False(t, validatedMetrics["redis.keyspace.hits"], "Found a duplicate in the metrics slice: redis.keyspace.hits")
					validatedMetrics["redis.keyspace.hits"] = true
					assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
					assert.Equal(t, "Number of successful lookup of keys in the main dictionary", ms.At(i).Description())
					assert.Equal(t, "", ms.At(i).Unit())
					assert.Equal(t, true, ms.At(i).Sum().IsMonotonic())
					assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
					dp := ms.At(i).Sum().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
					assert.Equal(t, int64(1), dp.IntValue())
				case "redis.keyspace.misses":
					assert.False(t, validatedMetrics["redis.keyspace.misses"], "Found a duplicate in the metrics slice: redis.keyspace.misses")
					validatedMetrics["redis.keyspace.misses"] = true
					assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
					assert.Equal(t, "Number of failed lookup of keys in the main dictionary", ms.At(i).Description())
					assert.Equal(t, "", ms.At(i).Unit())
					assert.Equal(t, true, ms.At(i).Sum().IsMonotonic())
					assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
					dp := ms.At(i).Sum().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
					assert.Equal(t, int64(1), dp.IntValue())
				case "redis.latest_fork":
					assert.False(t, validatedMetrics["redis.latest_fork"], "Found a duplicate in the metrics slice: redis.latest_fork")
					validatedMetrics["redis.latest_fork"] = true
					assert.Equal(t, pmetric.MetricTypeGauge, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Gauge().DataPoints().Len())
					assert.Equal(t, "Duration of the latest fork operation in microseconds", ms.At(i).Description())
					assert.Equal(t, "us", ms.At(i).Unit())
					dp := ms.At(i).Gauge().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
					assert.Equal(t, int64(1), dp.IntValue())
				case "redis.maxmemory":
					assert.False(t, validatedMetrics["redis.maxmemory"], "Found a duplicate in the metrics slice: redis.maxmemory")
					validatedMetrics["redis.maxmemory"] = true
					assert.Equal(t, pmetric.MetricTypeGauge, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Gauge().DataPoints().Len())
					assert.Equal(t, "The value of the maxmemory configuration directive", ms.At(i).Description())
					assert.Equal(t, "By", ms.At(i).Unit())
					dp := ms.At(i).Gauge().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
					assert.Equal(t, int64(1), dp.IntValue())
				case "redis.memory.fragmentation_ratio":
					assert.False(t, validatedMetrics["redis.memory.fragmentation_ratio"], "Found a duplicate in the metrics slice: redis.memory.fragmentation_ratio")
					validatedMetrics["redis.memory.fragmentation_ratio"] = true
					assert.Equal(t, pmetric.MetricTypeGauge, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Gauge().DataPoints().Len())
					assert.Equal(t, "Ratio between used_memory_rss and used_memory", ms.At(i).Description())
					assert.Equal(t, "", ms.At(i).Unit())
					dp := ms.At(i).Gauge().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeDouble, dp.ValueType())
					assert.Equal(t, float64(1), dp.DoubleValue())
				case "redis.memory.lua":
					assert.False(t, validatedMetrics["redis.memory.lua"], "Found a duplicate in the metrics slice: redis.memory.lua")
					validatedMetrics["redis.memory.lua"] = true
					assert.Equal(t, pmetric.MetricTypeGauge, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Gauge().DataPoints().Len())
					assert.Equal(t, "Number of bytes used by the Lua engine", ms.At(i).Description())
					assert.Equal(t, "By", ms.At(i).Unit())
					dp := ms.At(i).Gauge().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
					assert.Equal(t, int64(1), dp.IntValue())
				case "redis.memory.peak":
					assert.False(t, validatedMetrics["redis.memory.peak"], "Found a duplicate in the metrics slice: redis.memory.peak")
					validatedMetrics["redis.memory.peak"] = true
					assert.Equal(t, pmetric.MetricTypeGauge, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Gauge().DataPoints().Len())
					assert.Equal(t, "Peak memory consumed by Redis (in bytes)", ms.At(i).Description())
					assert.Equal(t, "By", ms.At(i).Unit())
					dp := ms.At(i).Gauge().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
					assert.Equal(t, int64(1), dp.IntValue())
				case "redis.memory.rss":
					assert.False(t, validatedMetrics["redis.memory.rss"], "Found a duplicate in the metrics slice: redis.memory.rss")
					validatedMetrics["redis.memory.rss"] = true
					assert.Equal(t, pmetric.MetricTypeGauge, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Gauge().DataPoints().Len())
					assert.Equal(t, "Number of bytes that Redis allocated as seen by the operating system", ms.At(i).Description())
					assert.Equal(t, "By", ms.At(i).Unit())
					dp := ms.At(i).Gauge().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
					assert.Equal(t, int64(1), dp.IntValue())
				case "redis.memory.used":
					assert.False(t, validatedMetrics["redis.memory.used"], "Found a duplicate in the metrics slice: redis.memory.used")
					validatedMetrics["redis.memory.used"] = true
					assert.Equal(t, pmetric.MetricTypeGauge, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Gauge().DataPoints().Len())
					assert.Equal(t, "Total number of bytes allocated by Redis using its allocator", ms.At(i).Description())
					assert.Equal(t, "By", ms.At(i).Unit())
					dp := ms.At(i).Gauge().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
					assert.Equal(t, int64(1), dp.IntValue())
				case "redis.net.input":
					assert.False(t, validatedMetrics["redis.net.input"], "Found a duplicate in the metrics slice: redis.net.input")
					validatedMetrics["redis.net.input"] = true
					assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
					assert.Equal(t, "The total number of bytes read from the network", ms.At(i).Description())
					assert.Equal(t, "By", ms.At(i).Unit())
					assert.Equal(t, true, ms.At(i).Sum().IsMonotonic())
					assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
					dp := ms.At(i).Sum().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
					assert.Equal(t, int64(1), dp.IntValue())
				case "redis.net.output":
					assert.False(t, validatedMetrics["redis.net.output"], "Found a duplicate in the metrics slice: redis.net.output")
					validatedMetrics["redis.net.output"] = true
					assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
					assert.Equal(t, "The total number of bytes written to the network", ms.At(i).Description())
					assert.Equal(t, "By", ms.At(i).Unit())
					assert.Equal(t, true, ms.At(i).Sum().IsMonotonic())
					assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
					dp := ms.At(i).Sum().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
					assert.Equal(t, int64(1), dp.IntValue())
				case "redis.rdb.changes_since_last_save":
					assert.False(t, validatedMetrics["redis.rdb.changes_since_last_save"], "Found a duplicate in the metrics slice: redis.rdb.changes_since_last_save")
					validatedMetrics["redis.rdb.changes_since_last_save"] = true
					assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
					assert.Equal(t, "Number of changes since the last dump", ms.At(i).Description())
					assert.Equal(t, "", ms.At(i).Unit())
					assert.Equal(t, false, ms.At(i).Sum().IsMonotonic())
					assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
					dp := ms.At(i).Sum().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
					assert.Equal(t, int64(1), dp.IntValue())
				case "redis.replication.backlog_first_byte_offset":
					assert.False(t, validatedMetrics["redis.replication.backlog_first_byte_offset"], "Found a duplicate in the metrics slice: redis.replication.backlog_first_byte_offset")
					validatedMetrics["redis.replication.backlog_first_byte_offset"] = true
					assert.Equal(t, pmetric.MetricTypeGauge, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Gauge().DataPoints().Len())
					assert.Equal(t, "The master offset of the replication backlog buffer", ms.At(i).Description())
					assert.Equal(t, "", ms.At(i).Unit())
					dp := ms.At(i).Gauge().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
					assert.Equal(t, int64(1), dp.IntValue())
				case "redis.replication.offset":
					assert.False(t, validatedMetrics["redis.replication.offset"], "Found a duplicate in the metrics slice: redis.replication.offset")
					validatedMetrics["redis.replication.offset"] = true
					assert.Equal(t, pmetric.MetricTypeGauge, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Gauge().DataPoints().Len())
					assert.Equal(t, "The server's current replication offset", ms.At(i).Description())
					assert.Equal(t, "", ms.At(i).Unit())
					dp := ms.At(i).Gauge().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
					assert.Equal(t, int64(1), dp.IntValue())
				case "redis.role":
					assert.False(t, validatedMetrics["redis.role"], "Found a duplicate in the metrics slice: redis.role")
					validatedMetrics["redis.role"] = true
					assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
					assert.Equal(t, "Redis node's role", ms.At(i).Description())
					assert.Equal(t, "", ms.At(i).Unit())
					assert.Equal(t, false, ms.At(i).Sum().IsMonotonic())
					assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
					dp := ms.At(i).Sum().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
					assert.Equal(t, int64(1), dp.IntValue())
					attrVal, ok := dp.Attributes().Get("role")
					assert.True(t, ok)
					assert.Equal(t, "replica", attrVal.Str())
				case "redis.slaves.connected":
					assert.False(t, validatedMetrics["redis.slaves.connected"], "Found a duplicate in the metrics slice: redis.slaves.connected")
					validatedMetrics["redis.slaves.connected"] = true
					assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
					assert.Equal(t, "Number of connected replicas", ms.At(i).Description())
					assert.Equal(t, "", ms.At(i).Unit())
					assert.Equal(t, false, ms.At(i).Sum().IsMonotonic())
					assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
					dp := ms.At(i).Sum().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
					assert.Equal(t, int64(1), dp.IntValue())
				case "redis.uptime":
					assert.False(t, validatedMetrics["redis.uptime"], "Found a duplicate in the metrics slice: redis.uptime")
					validatedMetrics["redis.uptime"] = true
					assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
					assert.Equal(t, "Number of seconds since Redis server start", ms.At(i).Description())
					assert.Equal(t, "s", ms.At(i).Unit())
					assert.Equal(t, true, ms.At(i).Sum().IsMonotonic())
					assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
					dp := ms.At(i).Sum().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
					assert.Equal(t, int64(1), dp.IntValue())
				}
			}
		})
	}
}

func loadConfig(t *testing.T, name string) MetricsBuilderConfig {
	cm, err := confmaptest.LoadConf(filepath.Join("testdata", "config.yaml"))
	require.NoError(t, err)
	sub, err := cm.Sub(name)
	require.NoError(t, err)
	cfg := DefaultMetricsBuilderConfig()
	require.NoError(t, component.UnmarshalConfig(sub, &cfg))
	return cfg
}
