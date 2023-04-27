package features

//lint:file-ignore U1000 we want to introduce this feature flag unused.

var (
	// csvExport enables CSV export of search results.
	csvExport = registerFeature("Enable CSV export of search results", "ROX_CSV_EXPORT", false)

	// NetworkDetectionBaselineSimulation enables new features related to the baseline simulation part of the network detection experience.
	NetworkDetectionBaselineSimulation = registerFeature("Enable network detection baseline simulation", "ROX_NETWORK_DETECTION_BASELINE_SIMULATION", true)

	// IntegrationsAsConfig enables loading integrations from config
	IntegrationsAsConfig = registerFeature("Enable loading integrations from config", "ROX_INTEGRATIONS_AS_CONFIG", false)

	// ComplianceOperatorCheckResults enables getting compliance results from the compliance operator
	ComplianceOperatorCheckResults = registerFeature("Enable fetching of compliance operator results", "ROX_COMPLIANCE_OPERATOR_INTEGRATION", true)

	// SystemHealthPatternFly enables the Pattern Fly version of System Health page. (used in the front-end app only)
	SystemHealthPatternFly = registerFeature("Enable Pattern Fly version of System Health page", "ROX_SYSTEM_HEALTH_PF", false)

	// NetworkPolicySystemPolicy enables two system policy fields (Missing (Ingress|Egress) Network Policy) to check deployments
	// against network policies applied in the secured cluster.
	NetworkPolicySystemPolicy = registerFeature("Enable NetworkPolicy-related system policy fields", "ROX_NETPOL_FIELDS", true)

	// QuayRobotAccounts enables Robot accounts as credentials in Quay Image Integration.
	QuayRobotAccounts = registerFeature("Enable Robot accounts in Quay Image Integration", "ROX_QUAY_ROBOT_ACCOUNTS", true)

	// RoxctlNetpolGenerate enables 'roxctl netpol generate' command which integrates with NP-Guard
	RoxctlNetpolGenerate = registerFeature("Enable 'roxctl generate netpol' command", "ROX_ROXCTL_NETPOL_GENERATE", true)

	// NetworkGraphPatternFly enables the PatternFly version of NetworkGraph. (used in the front-end app only)
	NetworkGraphPatternFly = registerFeature("Enable PatternFly version of NetworkGraph", "ROX_NETWORK_GRAPH_PATTERNFLY", true)

	// ClairV4Scanner enables Clair v4 as an Image Integration option
	ClairV4Scanner = registerFeature("Enable Clair v4 as an Image Integration option", "ROX_CLAIR_V4_SCANNING", true)

	// RoxSyslogExtraFields enables user to add additional key value pairs in syslog alert notification in cef format.
	RoxSyslogExtraFields = registerFeature("Enable extra fields for syslog integration", "ROX_SYSLOG_EXTRA_FIELDS", true)

	// SourcedAutogeneratedIntegrations enables adding a "source" to autogenerated integrations.
	SourcedAutogeneratedIntegrations = registerFeature("Enable autogenerated integrations with cluster/namespace/secret source", "ROX_SOURCED_AUTOGENERATED_INTEGRATIONS", false)

	// VulnMgmtWorkloadCVEs enables APIs and UI pages for the VM Workload CVE enhancements
	VulnMgmtWorkloadCVEs = registerFeature("Vuln Mgmt Workload CVEs", "ROX_VULN_MGMT_WORKLOAD_CVES", false)

	// PostgresBlobStore enables the creation of the Postgres Blob Store
	PostgresBlobStore = registerFeature("Postgres Blob Store", "ROX_POSTGRES_BLOB_STORE", false)

	// VulnMgmtReportingEnhancements enables APIs and UI pages for VM Reporting enhancements including downloadable reports
	VulnMgmtReportingEnhancements = registerFeature("Vuln Mgmt Reporting Enhancements", "ROX_VULN_MGMT_REPORTING_ENHANCEMENTS", false)

	// StoreEventHashes stores the hashes of successfully processed objects we receive from Sensor into the database
	StoreEventHashes = registerFeature("Store Event Hashes", "ROX_STORE_EVENT_HASHES", false)
)
