# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

from enum import Enum

__all__ = [
    'AccumulatedType',
    'ChartType',
    'ConnectorBillingModel',
    'CostAllocationPolicyType',
    'CostAllocationResourceType',
    'ExportType',
    'FormatType',
    'FunctionType',
    'GranularityType',
    'KpiTypeType',
    'MetricType',
    'OperatorType',
    'PivotTypeType',
    'RecurrenceType',
    'ReportColumnType',
    'ReportConfigColumnType',
    'ReportGranularityType',
    'ReportTimeframeType',
    'ReportType',
    'RuleStatus',
    'StatusType',
    'TimeframeType',
]


class AccumulatedType(str, Enum):
    """
    Show costs accumulated over time.
    """
    TRUE = "true"
    FALSE = "false"


class ChartType(str, Enum):
    """
    Chart type of the main view in Cost Analysis. Required.
    """
    AREA = "Area"
    LINE = "Line"
    STACKED_COLUMN = "StackedColumn"
    GROUPED_COLUMN = "GroupedColumn"
    TABLE = "Table"


class ConnectorBillingModel(str, Enum):
    """
    Connector billing model
    """
    TRIAL = "trial"
    AUTO_UPGRADE = "autoUpgrade"
    PREMIUM = "premium"
    EXPIRED = "expired"


class CostAllocationPolicyType(str, Enum):
    """
    Method of cost allocation for the rule
    """
    FIXED_PROPORTION = "FixedProportion"


class CostAllocationResourceType(str, Enum):
    """
    Type of resources contained in this cost allocation rule
    """
    DIMENSION = "Dimension"
    TAG = "Tag"


class ExportType(str, Enum):
    """
    The type of the export. Note that 'Usage' is equivalent to 'ActualCost' and is applicable to exports that do not yet provide data for charges or amortization for service reservations.
    """
    USAGE = "Usage"
    ACTUAL_COST = "ActualCost"
    AMORTIZED_COST = "AmortizedCost"


class FormatType(str, Enum):
    """
    The format of the report being delivered.
    """
    CSV = "Csv"


class FunctionType(str, Enum):
    """
    The name of the aggregation function to use.
    """
    SUM = "Sum"


class GranularityType(str, Enum):
    """
    The granularity of rows in the report.
    """
    DAILY = "Daily"
    HOURLY = "Hourly"


class KpiTypeType(str, Enum):
    """
    KPI type (Forecast, Budget).
    """
    FORECAST = "Forecast"
    BUDGET = "Budget"


class MetricType(str, Enum):
    """
    Metric to use when displaying costs.
    """
    ACTUAL_COST = "ActualCost"
    AMORTIZED_COST = "AmortizedCost"
    AHUB = "AHUB"


class OperatorType(str, Enum):
    """
    The operator to use for comparison.
    """
    IN_ = "In"
    CONTAINS = "Contains"


class PivotTypeType(str, Enum):
    """
    Data type to show in view.
    """
    DIMENSION = "Dimension"
    TAG_KEY = "TagKey"


class RecurrenceType(str, Enum):
    """
    The schedule recurrence.
    """
    DAILY = "Daily"
    WEEKLY = "Weekly"
    MONTHLY = "Monthly"
    ANNUALLY = "Annually"


class ReportColumnType(str, Enum):
    """
    Has type of the column to group.
    """
    TAG = "Tag"
    DIMENSION = "Dimension"


class ReportConfigColumnType(str, Enum):
    """
    Has type of the column to group.
    """
    TAG = "Tag"
    DIMENSION = "Dimension"


class ReportGranularityType(str, Enum):
    """
    The granularity of rows in the report.
    """
    DAILY = "Daily"
    MONTHLY = "Monthly"


class ReportTimeframeType(str, Enum):
    """
    The time frame for pulling data for the report. If custom, then a specific time period must be provided.
    """
    WEEK_TO_DATE = "WeekToDate"
    MONTH_TO_DATE = "MonthToDate"
    YEAR_TO_DATE = "YearToDate"
    CUSTOM = "Custom"


class ReportType(str, Enum):
    """
    The type of the report. Usage represents actual usage, forecast represents forecasted data and UsageAndForecast represents both usage and forecasted data. Actual usage and forecasted data can be differentiated based on dates.
    """
    USAGE = "Usage"


class RuleStatus(str, Enum):
    """
    Status of the rule
    """
    NOT_ACTIVE = "NotActive"
    ACTIVE = "Active"
    PROCESSING = "Processing"


class StatusType(str, Enum):
    """
    The status of the schedule. Whether active or not. If inactive, the report's scheduled execution is paused.
    """
    ACTIVE = "Active"
    INACTIVE = "Inactive"


class TimeframeType(str, Enum):
    """
    The time frame for pulling data for the report. If custom, then a specific time period must be provided.
    """
    WEEK_TO_DATE = "WeekToDate"
    MONTH_TO_DATE = "MonthToDate"
    CUSTOM = "Custom"
