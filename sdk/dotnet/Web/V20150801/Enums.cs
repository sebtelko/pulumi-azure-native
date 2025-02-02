// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.ComponentModel;
using Pulumi;

namespace Pulumi.AzureNative.Web.V20150801
{
    [EnumType]
    public readonly struct AccessControlEntryAction : IEquatable<AccessControlEntryAction>
    {
        private readonly string _value;

        private AccessControlEntryAction(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static AccessControlEntryAction Permit { get; } = new AccessControlEntryAction("Permit");
        public static AccessControlEntryAction Deny { get; } = new AccessControlEntryAction("Deny");

        public static bool operator ==(AccessControlEntryAction left, AccessControlEntryAction right) => left.Equals(right);
        public static bool operator !=(AccessControlEntryAction left, AccessControlEntryAction right) => !left.Equals(right);

        public static explicit operator string(AccessControlEntryAction value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is AccessControlEntryAction other && Equals(other);
        public bool Equals(AccessControlEntryAction other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// ActionType - predefined action to be taken
    /// </summary>
    [EnumType]
    public readonly struct AutoHealActionType : IEquatable<AutoHealActionType>
    {
        private readonly string _value;

        private AutoHealActionType(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static AutoHealActionType Recycle { get; } = new AutoHealActionType("Recycle");
        public static AutoHealActionType LogEvent { get; } = new AutoHealActionType("LogEvent");
        public static AutoHealActionType CustomAction { get; } = new AutoHealActionType("CustomAction");

        public static bool operator ==(AutoHealActionType left, AutoHealActionType right) => left.Equals(right);
        public static bool operator !=(AutoHealActionType left, AutoHealActionType right) => !left.Equals(right);

        public static explicit operator string(AutoHealActionType value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is AutoHealActionType other && Equals(other);
        public bool Equals(AutoHealActionType other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// Azure resource type
    /// </summary>
    [EnumType]
    public readonly struct AzureResourceType : IEquatable<AzureResourceType>
    {
        private readonly string _value;

        private AzureResourceType(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static AzureResourceType Website { get; } = new AzureResourceType("Website");
        public static AzureResourceType TrafficManager { get; } = new AzureResourceType("TrafficManager");

        public static bool operator ==(AzureResourceType left, AzureResourceType right) => left.Equals(right);
        public static bool operator !=(AzureResourceType left, AzureResourceType right) => !left.Equals(right);

        public static explicit operator string(AzureResourceType value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is AzureResourceType other && Equals(other);
        public bool Equals(AzureResourceType other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// Type of the backup
    /// </summary>
    [EnumType]
    public readonly struct BackupRestoreOperationType : IEquatable<BackupRestoreOperationType>
    {
        private readonly string _value;

        private BackupRestoreOperationType(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static BackupRestoreOperationType Default { get; } = new BackupRestoreOperationType("Default");
        public static BackupRestoreOperationType Clone { get; } = new BackupRestoreOperationType("Clone");
        public static BackupRestoreOperationType Relocation { get; } = new BackupRestoreOperationType("Relocation");

        public static bool operator ==(BackupRestoreOperationType left, BackupRestoreOperationType right) => left.Equals(right);
        public static bool operator !=(BackupRestoreOperationType left, BackupRestoreOperationType right) => !left.Equals(right);

        public static explicit operator string(BackupRestoreOperationType value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is BackupRestoreOperationType other && Equals(other);
        public bool Equals(BackupRestoreOperationType other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// Gets or sets the default authentication provider to use when multiple providers are configured.
    ///             This setting is only needed if multiple providers are configured and the unauthenticated client
    ///             action is set to "RedirectToLoginPage".
    /// </summary>
    [EnumType]
    public readonly struct BuiltInAuthenticationProvider : IEquatable<BuiltInAuthenticationProvider>
    {
        private readonly string _value;

        private BuiltInAuthenticationProvider(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static BuiltInAuthenticationProvider AzureActiveDirectory { get; } = new BuiltInAuthenticationProvider("AzureActiveDirectory");
        public static BuiltInAuthenticationProvider Facebook { get; } = new BuiltInAuthenticationProvider("Facebook");
        public static BuiltInAuthenticationProvider Google { get; } = new BuiltInAuthenticationProvider("Google");
        public static BuiltInAuthenticationProvider MicrosoftAccount { get; } = new BuiltInAuthenticationProvider("MicrosoftAccount");
        public static BuiltInAuthenticationProvider Twitter { get; } = new BuiltInAuthenticationProvider("Twitter");

        public static bool operator ==(BuiltInAuthenticationProvider left, BuiltInAuthenticationProvider right) => left.Equals(right);
        public static bool operator !=(BuiltInAuthenticationProvider left, BuiltInAuthenticationProvider right) => !left.Equals(right);

        public static explicit operator string(BuiltInAuthenticationProvider value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is BuiltInAuthenticationProvider other && Equals(other);
        public bool Equals(BuiltInAuthenticationProvider other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// Shared or dedicated web app hosting
    /// </summary>
    [EnumType]
    public readonly struct ComputeModeOptions : IEquatable<ComputeModeOptions>
    {
        private readonly string _value;

        private ComputeModeOptions(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static ComputeModeOptions Shared { get; } = new ComputeModeOptions("Shared");
        public static ComputeModeOptions Dedicated { get; } = new ComputeModeOptions("Dedicated");
        public static ComputeModeOptions Dynamic { get; } = new ComputeModeOptions("Dynamic");

        public static bool operator ==(ComputeModeOptions left, ComputeModeOptions right) => left.Equals(right);
        public static bool operator !=(ComputeModeOptions left, ComputeModeOptions right) => !left.Equals(right);

        public static explicit operator string(ComputeModeOptions value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is ComputeModeOptions other && Equals(other);
        public bool Equals(ComputeModeOptions other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// Custom DNS record type
    /// </summary>
    [EnumType]
    public readonly struct CustomHostNameDnsRecordType : IEquatable<CustomHostNameDnsRecordType>
    {
        private readonly string _value;

        private CustomHostNameDnsRecordType(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static CustomHostNameDnsRecordType CName { get; } = new CustomHostNameDnsRecordType("CName");
        public static CustomHostNameDnsRecordType A { get; } = new CustomHostNameDnsRecordType("A");

        public static bool operator ==(CustomHostNameDnsRecordType left, CustomHostNameDnsRecordType right) => left.Equals(right);
        public static bool operator !=(CustomHostNameDnsRecordType left, CustomHostNameDnsRecordType right) => !left.Equals(right);

        public static explicit operator string(CustomHostNameDnsRecordType value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is CustomHostNameDnsRecordType other && Equals(other);
        public bool Equals(CustomHostNameDnsRecordType other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// Type of database
    /// </summary>
    [EnumType]
    public readonly struct DatabaseServerType : IEquatable<DatabaseServerType>
    {
        private readonly string _value;

        private DatabaseServerType(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static DatabaseServerType MySql { get; } = new DatabaseServerType("MySql");
        public static DatabaseServerType SQLServer { get; } = new DatabaseServerType("SQLServer");
        public static DatabaseServerType SQLAzure { get; } = new DatabaseServerType("SQLAzure");
        public static DatabaseServerType Custom { get; } = new DatabaseServerType("Custom");

        public static bool operator ==(DatabaseServerType left, DatabaseServerType right) => left.Equals(right);
        public static bool operator !=(DatabaseServerType left, DatabaseServerType right) => !left.Equals(right);

        public static explicit operator string(DatabaseServerType value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is DatabaseServerType other && Equals(other);
        public bool Equals(DatabaseServerType other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// How often should be the backup executed (e.g. for weekly backup, this should be set to Day and FrequencyInterval should be set to 7)
    /// </summary>
    [EnumType]
    public readonly struct FrequencyUnit : IEquatable<FrequencyUnit>
    {
        private readonly string _value;

        private FrequencyUnit(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static FrequencyUnit Day { get; } = new FrequencyUnit("Day");
        public static FrequencyUnit Hour { get; } = new FrequencyUnit("Hour");

        public static bool operator ==(FrequencyUnit left, FrequencyUnit right) => left.Equals(right);
        public static bool operator !=(FrequencyUnit left, FrequencyUnit right) => !left.Equals(right);

        public static explicit operator string(FrequencyUnit value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is FrequencyUnit other && Equals(other);
        public bool Equals(FrequencyUnit other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// Host name type
    /// </summary>
    [EnumType]
    public readonly struct HostNameType : IEquatable<HostNameType>
    {
        private readonly string _value;

        private HostNameType(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static HostNameType Verified { get; } = new HostNameType("Verified");
        public static HostNameType Managed { get; } = new HostNameType("Managed");

        public static bool operator ==(HostNameType left, HostNameType right) => left.Equals(right);
        public static bool operator !=(HostNameType left, HostNameType right) => !left.Equals(right);

        public static explicit operator string(HostNameType value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is HostNameType other && Equals(other);
        public bool Equals(HostNameType other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// Current status of the hostingEnvironment (App Service Environment)
    /// </summary>
    [EnumType]
    public readonly struct HostingEnvironmentStatus : IEquatable<HostingEnvironmentStatus>
    {
        private readonly string _value;

        private HostingEnvironmentStatus(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static HostingEnvironmentStatus Preparing { get; } = new HostingEnvironmentStatus("Preparing");
        public static HostingEnvironmentStatus Ready { get; } = new HostingEnvironmentStatus("Ready");
        public static HostingEnvironmentStatus Scaling { get; } = new HostingEnvironmentStatus("Scaling");
        public static HostingEnvironmentStatus Deleting { get; } = new HostingEnvironmentStatus("Deleting");

        public static bool operator ==(HostingEnvironmentStatus left, HostingEnvironmentStatus right) => left.Equals(right);
        public static bool operator !=(HostingEnvironmentStatus left, HostingEnvironmentStatus right) => !left.Equals(right);

        public static explicit operator string(HostingEnvironmentStatus value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is HostingEnvironmentStatus other && Equals(other);
        public bool Equals(HostingEnvironmentStatus other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// Specifies which endpoints to serve internally in the hostingEnvironment's (App Service Environment) VNET
    /// </summary>
    [EnumType]
    public readonly struct InternalLoadBalancingMode : IEquatable<InternalLoadBalancingMode>
    {
        private readonly string _value;

        private InternalLoadBalancingMode(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static InternalLoadBalancingMode None { get; } = new InternalLoadBalancingMode("None");
        public static InternalLoadBalancingMode Web { get; } = new InternalLoadBalancingMode("Web");
        public static InternalLoadBalancingMode Publishing { get; } = new InternalLoadBalancingMode("Publishing");

        public static bool operator ==(InternalLoadBalancingMode left, InternalLoadBalancingMode right) => left.Equals(right);
        public static bool operator !=(InternalLoadBalancingMode left, InternalLoadBalancingMode right) => !left.Equals(right);

        public static explicit operator string(InternalLoadBalancingMode value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is InternalLoadBalancingMode other && Equals(other);
        public bool Equals(InternalLoadBalancingMode other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// Log level
    /// </summary>
    [EnumType]
    public readonly struct LogLevel : IEquatable<LogLevel>
    {
        private readonly string _value;

        private LogLevel(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static LogLevel Off { get; } = new LogLevel("Off");
        public static LogLevel Verbose { get; } = new LogLevel("Verbose");
        public static LogLevel Information { get; } = new LogLevel("Information");
        public static LogLevel Warning { get; } = new LogLevel("Warning");
        public static LogLevel Error { get; } = new LogLevel("Error");

        public static bool operator ==(LogLevel left, LogLevel right) => left.Equals(right);
        public static bool operator !=(LogLevel left, LogLevel right) => !left.Equals(right);

        public static explicit operator string(LogLevel value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is LogLevel other && Equals(other);
        public bool Equals(LogLevel other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// Managed pipeline mode
    /// </summary>
    [EnumType]
    public readonly struct ManagedPipelineMode : IEquatable<ManagedPipelineMode>
    {
        private readonly string _value;

        private ManagedPipelineMode(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static ManagedPipelineMode Integrated { get; } = new ManagedPipelineMode("Integrated");
        public static ManagedPipelineMode Classic { get; } = new ManagedPipelineMode("Classic");

        public static bool operator ==(ManagedPipelineMode left, ManagedPipelineMode right) => left.Equals(right);
        public static bool operator !=(ManagedPipelineMode left, ManagedPipelineMode right) => !left.Equals(right);

        public static explicit operator string(ManagedPipelineMode value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is ManagedPipelineMode other && Equals(other);
        public bool Equals(ManagedPipelineMode other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// Provisioning state of the hostingEnvironment (App Service Environment)
    /// </summary>
    [EnumType]
    public readonly struct ProvisioningState : IEquatable<ProvisioningState>
    {
        private readonly string _value;

        private ProvisioningState(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static ProvisioningState Succeeded { get; } = new ProvisioningState("Succeeded");
        public static ProvisioningState Failed { get; } = new ProvisioningState("Failed");
        public static ProvisioningState Canceled { get; } = new ProvisioningState("Canceled");
        public static ProvisioningState InProgress { get; } = new ProvisioningState("InProgress");
        public static ProvisioningState Deleting { get; } = new ProvisioningState("Deleting");

        public static bool operator ==(ProvisioningState left, ProvisioningState right) => left.Equals(right);
        public static bool operator !=(ProvisioningState left, ProvisioningState right) => !left.Equals(right);

        public static explicit operator string(ProvisioningState value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is ProvisioningState other && Equals(other);
        public bool Equals(ProvisioningState other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// Site load balancing
    /// </summary>
    [EnumType]
    public readonly struct SiteLoadBalancing : IEquatable<SiteLoadBalancing>
    {
        private readonly string _value;

        private SiteLoadBalancing(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static SiteLoadBalancing WeightedRoundRobin { get; } = new SiteLoadBalancing("WeightedRoundRobin");
        public static SiteLoadBalancing LeastRequests { get; } = new SiteLoadBalancing("LeastRequests");
        public static SiteLoadBalancing LeastResponseTime { get; } = new SiteLoadBalancing("LeastResponseTime");
        public static SiteLoadBalancing WeightedTotalTraffic { get; } = new SiteLoadBalancing("WeightedTotalTraffic");
        public static SiteLoadBalancing RequestHash { get; } = new SiteLoadBalancing("RequestHash");

        public static bool operator ==(SiteLoadBalancing left, SiteLoadBalancing right) => left.Equals(right);
        public static bool operator !=(SiteLoadBalancing left, SiteLoadBalancing right) => !left.Equals(right);

        public static explicit operator string(SiteLoadBalancing value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is SiteLoadBalancing other && Equals(other);
        public bool Equals(SiteLoadBalancing other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// SSL type
    /// </summary>
    [EnumType]
    public readonly struct SslState : IEquatable<SslState>
    {
        private readonly string _value;

        private SslState(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static SslState Disabled { get; } = new SslState("Disabled");
        public static SslState SniEnabled { get; } = new SslState("SniEnabled");
        public static SslState IpBasedEnabled { get; } = new SslState("IpBasedEnabled");

        public static bool operator ==(SslState left, SslState right) => left.Equals(right);
        public static bool operator !=(SslState left, SslState right) => !left.Equals(right);

        public static explicit operator string(SslState value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is SslState other && Equals(other);
        public bool Equals(SslState other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// Gets or sets the action to take when an unauthenticated client attempts to access the app.
    /// </summary>
    [EnumType]
    public readonly struct UnauthenticatedClientAction : IEquatable<UnauthenticatedClientAction>
    {
        private readonly string _value;

        private UnauthenticatedClientAction(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static UnauthenticatedClientAction RedirectToLoginPage { get; } = new UnauthenticatedClientAction("RedirectToLoginPage");
        public static UnauthenticatedClientAction AllowAnonymous { get; } = new UnauthenticatedClientAction("AllowAnonymous");

        public static bool operator ==(UnauthenticatedClientAction left, UnauthenticatedClientAction right) => left.Equals(right);
        public static bool operator !=(UnauthenticatedClientAction left, UnauthenticatedClientAction right) => !left.Equals(right);

        public static explicit operator string(UnauthenticatedClientAction value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is UnauthenticatedClientAction other && Equals(other);
        public bool Equals(UnauthenticatedClientAction other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }

    /// <summary>
    /// Size of the machines
    /// </summary>
    [EnumType]
    public readonly struct WorkerSizeOptions : IEquatable<WorkerSizeOptions>
    {
        private readonly string _value;

        private WorkerSizeOptions(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static WorkerSizeOptions Default { get; } = new WorkerSizeOptions("Default");
        public static WorkerSizeOptions Small { get; } = new WorkerSizeOptions("Small");
        public static WorkerSizeOptions Medium { get; } = new WorkerSizeOptions("Medium");
        public static WorkerSizeOptions Large { get; } = new WorkerSizeOptions("Large");

        public static bool operator ==(WorkerSizeOptions left, WorkerSizeOptions right) => left.Equals(right);
        public static bool operator !=(WorkerSizeOptions left, WorkerSizeOptions right) => !left.Equals(right);

        public static explicit operator string(WorkerSizeOptions value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is WorkerSizeOptions other && Equals(other);
        public bool Equals(WorkerSizeOptions other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;
    }
}
