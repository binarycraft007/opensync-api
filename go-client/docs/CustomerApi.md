# \CustomerApi

All URIs are relative to *https://localhost/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CustomerConfirm**](CustomerApi.md#CustomerConfirm) | **Get** /Customers/confirm | Confirm a user registration with identity verification token.
[**CustomerCount**](CustomerApi.md#CustomerCount) | **Get** /Customers/count | Count instances of the model matched by where from the data source.
[**CustomerCreateOrUpdateUser**](CustomerApi.md#CustomerCreateOrUpdateUser) | **Put** /Customers/createOrUpdateUser | Create or update a Plume NOC user.
[**CustomerCustomCreate**](CustomerApi.md#CustomerCustomCreate) | **Post** /Customers | Create a Plume customer.
[**CustomerEmailExists**](CustomerApi.md#CustomerEmailExists) | **Get** /Customers/exists | 
[**CustomerEmailPasswordlessToken**](CustomerApi.md#CustomerEmailPasswordlessToken) | **Post** /Customers/passwordLessToken | Generate two accessTokens with special scopes for the account with the email address and send a verification email.
[**CustomerFind**](CustomerApi.md#CustomerFind) | **Get** /Customers | Find all instances of the model matched by filter from the data source.
[**CustomerFindById**](CustomerApi.md#CustomerFindById) | **Get** /Customers/{id} | Find a model instance by {{id}} from the data source.
[**CustomerImportData**](CustomerApi.md#CustomerImportData) | **Post** /Customers/import | Import customer data
[**CustomerLogin**](CustomerApi.md#CustomerLogin) | **Post** /Customers/login | Login a user with username/email and password.
[**CustomerLogoutPostCustomersLogout**](CustomerApi.md#CustomerLogoutPostCustomersLogout) | **Post** /Customers/logout | Logout a user with access token.
[**CustomerPrototypeAddDeviceVisibleToGuests**](CustomerApi.md#CustomerPrototypeAddDeviceVisibleToGuests) | **Post** /Customers/{id}/locations/{locationId}/wifiNetwork/accessZones/home/devicesVisibleToGuests/{mac} | DEPRECATED: Update home devices visible to guests.
[**CustomerPrototypeAppFacadeHome**](CustomerApi.md#CustomerPrototypeAppFacadeHome) | **Get** /Customers/{id}/locations/{locationId}/appFacade/home | Retrieve timezone, capabilities, summary, ... for this location.
[**CustomerPrototypeApproveDevices**](CustomerApi.md#CustomerPrototypeApproveDevices) | **Post** /Customers/{id}/locations/{locationId}/networkAccess/networks/{networkId}/approved | Approve devices in the network
[**CustomerPrototypeApproveWhitelistRequest**](CustomerApi.md#CustomerPrototypeApproveWhitelistRequest) | **Put** /Customers/{id}/locations/{locationId}/securityPolicy/websites/whitelist/approvalRequests/{requestId} | Approve a persons whitelist request and add it to the security policy.
[**CustomerPrototypeBlockDevices**](CustomerApi.md#CustomerPrototypeBlockDevices) | **Post** /Customers/{id}/locations/{locationId}/networkAccess/blocked | Block devices
[**CustomerPrototypeClaimNode**](CustomerApi.md#CustomerPrototypeClaimNode) | **Post** /Customers/{id}/locations/{locationId}/nodes | Claim a node and all nodes still associated to its Package ID for a Location ID.
[**CustomerPrototypeCreateAccessTokens**](CustomerApi.md#CustomerPrototypeCreateAccessTokens) | **Post** /Customers/{id}/accessTokens | Creates a new instance in accessTokens of this model.
[**CustomerPrototypeCreateGetMarketingExportDataAccessToken**](CustomerApi.md#CustomerPrototypeCreateGetMarketingExportDataAccessToken) | **Post** /Customers/{id}/createGetMarketingExportDataAccessToken | Create access token to get marketing data by CRM for campaigns
[**CustomerPrototypeCreateIpLimitedAccessToken**](CustomerApi.md#CustomerPrototypeCreateIpLimitedAccessToken) | **Post** /Customers/{id}/createIpLimitedAccessToken | Create access token with limited privileges as defined for IP authenticated customers
[**CustomerPrototypeCreateLocation**](CustomerApi.md#CustomerPrototypeCreateLocation) | **Post** /Customers/{id}/locations | Create a new location.
[**CustomerPrototypeCreateMigration**](CustomerApi.md#CustomerPrototypeCreateMigration) | **Post** /Customers/{id}/_migration | Creates a new instance in _migration of this model.
[**CustomerPrototypeCreateNewPasswordlessToken**](CustomerApi.md#CustomerPrototypeCreateNewPasswordlessToken) | **Post** /Customers/{id}/accessToken | Generates usable passwordless accessToken for the account with the email address.
[**CustomerPrototypeCreateOauthAccessToken**](CustomerApi.md#CustomerPrototypeCreateOauthAccessToken) | **Post** /Customers/{id}/createOauthAccessToken | Create access token with ouath scope.
[**CustomerPrototypeCreateOauthRefreshToken**](CustomerApi.md#CustomerPrototypeCreateOauthRefreshToken) | **Post** /Customers/{id}/createOauthRefreshToken | Create refresh token for a specific access token with ouath scope.
[**CustomerPrototypeCreatePatchServiceLevelAccessToken**](CustomerApi.md#CustomerPrototypeCreatePatchServiceLevelAccessToken) | **Post** /Customers/{id}/createPatchServiceLevelAccessToken | Create access token to patch customer serviceLevel used by ZUORA
[**CustomerPrototypeCreateReadDnsAccessToken**](CustomerApi.md#CustomerPrototypeCreateReadDnsAccessToken) | **Post** /Customers/{id}/createReadDnsAccessToken | Create access token to read data related to DNS security policies
[**CustomerPrototypeDeleteAllDeviceFreezes**](CustomerApi.md#CustomerPrototypeDeleteAllDeviceFreezes) | **Delete** /Customers/{id}/locations/{locationId}/devices/{mac}/freezes | Delete/clear all device freezes templateIds for a mac.
[**CustomerPrototypeDeleteAppPrioritizationLocationConfig**](CustomerApi.md#CustomerPrototypeDeleteAppPrioritizationLocationConfig) | **Delete** /Customers/{id}/locations/{locationId}/qos/appPrioritization/customSetting | Set custom setting to default for app prioritization.
[**CustomerPrototypeDeleteCaptivePortal**](CustomerApi.md#CustomerPrototypeDeleteCaptivePortal) | **Delete** /Customers/{id}/locations/{locationId}/secondaryNetworks/captivePortals/{networkId} | Delete a CaptivePortal for a Location ID.
[**CustomerPrototypeDeleteCaptivePortalAuthorizedClients**](CustomerApi.md#CustomerPrototypeDeleteCaptivePortalAuthorizedClients) | **Delete** /Customers/{id}/locations/{locationId}/secondaryNetworks/captivePortals/{networkId}/authorizedClients/{mac} | Delete Authorized Client
[**CustomerPrototypeDeleteConfigs**](CustomerApi.md#CustomerPrototypeDeleteConfigs) | **Delete** /Customers/{id}/locations/{locationId}/configs | Delete specified location settings, while keeping claimed nodes intact
[**CustomerPrototypeDeleteCustomSharedSchedule**](CustomerApi.md#CustomerPrototypeDeleteCustomSharedSchedule) | **Delete** /Customers/{id}/locations/{locationId}/schedules/{templateId} | Delete \&quot;custom shared\&quot; schedule shared by all persons and devices in a location.
[**CustomerPrototypeDeleteCustomer**](CustomerApi.md#CustomerPrototypeDeleteCustomer) | **Delete** /Customers/{id} | Delete a model instance by {{id}} from the data source.
[**CustomerPrototypeDeleteDeviceFreeze**](CustomerApi.md#CustomerPrototypeDeleteDeviceFreeze) | **Delete** /Customers/{id}/locations/{locationId}/devices/{mac}/freeze/{freezeTemplateId} | Delete a device to be frozen for a Location ID.
[**CustomerPrototypeDeleteDeviceFreezeAutoExpire**](CustomerApi.md#CustomerPrototypeDeleteDeviceFreezeAutoExpire) | **Delete** /Customers/{id}/locations/{locationId}/devices/{mac}/freeze/autoExpire | Delete a device to be frozen for a Location ID.
[**CustomerPrototypeDeleteDeviceFreezeForever**](CustomerApi.md#CustomerPrototypeDeleteDeviceFreezeForever) | **Delete** /Customers/{id}/locations/{locationId}/devices/{mac}/freeze/forever | Delete a device forever freeze for a Location ID.
[**CustomerPrototypeDeleteDeviceFreezeResidentialGwManaged**](CustomerApi.md#CustomerPrototypeDeleteDeviceFreezeResidentialGwManaged) | **Delete** /Customers/{id}/locations/{locationId}/devices/{mac}/freeze/residentialGwManaged | Delete a device residentialGwManaged freeze for a Location ID.
[**CustomerPrototypeDeleteDeviceFreezeSuspend**](CustomerApi.md#CustomerPrototypeDeleteDeviceFreezeSuspend) | **Delete** /Customers/{id}/locations/{locationId}/devices/{mac}/freeze/suspend | Delete a device suspend for a Location ID.
[**CustomerPrototypeDeleteDeviceFromAccessZone**](CustomerApi.md#CustomerPrototypeDeleteDeviceFromAccessZone) | **Delete** /Customers/{id}/locations/{locationId}/wifiNetwork/accessZones/{zoneId}/accessibleDevices/{mac} | Delete a device mac from a WiFi Access Zone
[**CustomerPrototypeDeleteDeviceFromPerson**](CustomerApi.md#CustomerPrototypeDeleteDeviceFromPerson) | **Delete** /Customers/{id}/locations/{locationId}/persons/{personId}/devices/{mac} | Unassign a device from Person for a location ID.
[**CustomerPrototypeDeleteDeviceGroup**](CustomerApi.md#CustomerPrototypeDeleteDeviceGroup) | **Delete** /Customers/{id}/locations/{locationId}/networkAccess/networks/{networkId}/deviceGroups/{groupId} | Delete a device group from a network.
[**CustomerPrototypeDeleteDeviceQosPrioritization**](CustomerApi.md#CustomerPrototypeDeleteDeviceQosPrioritization) | **Delete** /Customers/{id}/locations/{locationId}/devices/{mac}/qos/prioritization | Delete prioritization of a single device
[**CustomerPrototypeDeleteDeviceSecurityPolicyAnomalyExperience**](CustomerApi.md#CustomerPrototypeDeleteDeviceSecurityPolicyAnomalyExperience) | **Delete** /Customers/{id}/locations/{locationId}/devices/{mac}/securityPolicy/anomaly/experience | Delete an Anomaly Experience (demo) for a Device on a location.
[**CustomerPrototypeDeleteDeviceSecurityPolicyAnomalyWhitelist**](CustomerApi.md#CustomerPrototypeDeleteDeviceSecurityPolicyAnomalyWhitelist) | **Delete** /Customers/{id}/locations/{locationId}/devices/{mac}/securityPolicy/anomaly/websites/whitelist/{fqdn} | Update a Location&#39;s Anomaly Security Policy for a location ID to remove a whitelisted DNS entry.
[**CustomerPrototypeDeleteDhcpReservation**](CustomerApi.md#CustomerPrototypeDeleteDhcpReservation) | **Delete** /Customers/{id}/locations/{locationId}/networkConfiguration/dhcpReservations/{mac} | Delete a current DHCP IP reservation and the associated port forwarding entries for a particular MAC address at a Location ID.
[**CustomerPrototypeDeleteForcedSteer**](CustomerApi.md#CustomerPrototypeDeleteForcedSteer) | **Delete** /Customers/{id}/locations/{locationId}/devices/{mac}/forcedSteer | Disable 2.4Ghz band enforcement early.
[**CustomerPrototypeDeleteFromDeviceSecurityPolicyWebsitesBlacklist**](CustomerApi.md#CustomerPrototypeDeleteFromDeviceSecurityPolicyWebsitesBlacklist) | **Delete** /Customers/{id}/locations/{locationId}/devices/{mac}/securityPolicy/websites/blacklist/{dns} | Update a Device&#39;s Security Policy for a location ID to remove a blacklisted DNS entry.
[**CustomerPrototypeDeleteFromDeviceSecurityPolicyWebsitesWhitelist**](CustomerApi.md#CustomerPrototypeDeleteFromDeviceSecurityPolicyWebsitesWhitelist) | **Delete** /Customers/{id}/locations/{locationId}/devices/{mac}/securityPolicy/websites/whitelist/{dns} | Update a Device&#39;s Security Policy for a location ID to remove a whitelisted DNS entry.
[**CustomerPrototypeDeleteFromGroupOfUnassignedDevicesSecurityPolicyWebsitesBlacklist**](CustomerApi.md#CustomerPrototypeDeleteFromGroupOfUnassignedDevicesSecurityPolicyWebsitesBlacklist) | **Delete** /Customers/{id}/locations/{locationId}/groupOfUnassignedDevices/securityPolicy/websites/blacklist/{dns} | Update a Location&#39;s Default Device Group Security Policy for a location ID to remove a blacklisted DNS entry.
[**CustomerPrototypeDeleteFromGroupOfUnassignedDevicesSecurityPolicyWebsitesWhitelist**](CustomerApi.md#CustomerPrototypeDeleteFromGroupOfUnassignedDevicesSecurityPolicyWebsitesWhitelist) | **Delete** /Customers/{id}/locations/{locationId}/groupOfUnassignedDevices/securityPolicy/websites/whitelist/{dns} | Update a Location&#39;s Default Device Group Security Policy for a location ID to remove a whitelisted DNS entry.
[**CustomerPrototypeDeleteFromLocationSecurityPolicyWebsitesBlacklist**](CustomerApi.md#CustomerPrototypeDeleteFromLocationSecurityPolicyWebsitesBlacklist) | **Delete** /Customers/{id}/locations/{locationId}/securityPolicy/websites/blacklist/{dns} | Update a Location&#39;s Security Policy for a location ID to remove a blacklisted DNS entry.
[**CustomerPrototypeDeleteFromLocationSecurityPolicyWebsitesWhitelist**](CustomerApi.md#CustomerPrototypeDeleteFromLocationSecurityPolicyWebsitesWhitelist) | **Delete** /Customers/{id}/locations/{locationId}/securityPolicy/websites/whitelist/{dns} | Update a Locations&#39;s Security Policy for a location ID to remove a whitelisted DNS entry.
[**CustomerPrototypeDeleteFromPersonSecurityPolicyWebsitesBlacklist**](CustomerApi.md#CustomerPrototypeDeleteFromPersonSecurityPolicyWebsitesBlacklist) | **Delete** /Customers/{id}/locations/{locationId}/persons/{personId}/securityPolicy/websites/blacklist/{dns} | Update a Person&#39;s Security Policy for a location ID to remove a blacklisted DNS entry.
[**CustomerPrototypeDeleteFromPersonSecurityPolicyWebsitesWhitelist**](CustomerApi.md#CustomerPrototypeDeleteFromPersonSecurityPolicyWebsitesWhitelist) | **Delete** /Customers/{id}/locations/{locationId}/persons/{personId}/securityPolicy/websites/whitelist/{dns} | Update a Person&#39;s Security Policy for a location ID to remove a whitelisted DNS entry.
[**CustomerPrototypeDeleteFrontHaul**](CustomerApi.md#CustomerPrototypeDeleteFrontHaul) | **Delete** /Customers/{id}/locations/{locationId}/secondaryNetworks/fronthauls/{networkId} | Delete a Front Haul for a Location ID.
[**CustomerPrototypeDeleteGdprCaptivePortalsData**](CustomerApi.md#CustomerPrototypeDeleteGdprCaptivePortalsData) | **Post** /Customers/{id}/locations/{locationId}/secondaryNetworks/captivePortal/{networkId}/gdprForget/guests | Delete the Gdpr Captive Portals data for a guest.
[**CustomerPrototypeDeleteGroupOfUnassignedDevicesFreezeAutoExpire**](CustomerApi.md#CustomerPrototypeDeleteGroupOfUnassignedDevicesFreezeAutoExpire) | **Delete** /Customers/{id}/locations/{locationId}/groupOfUnassignedDevices/freeze/autoExpire | Delete GroupOfUnassignedDevices autoExpire freeze for a Location ID.
[**CustomerPrototypeDeleteGroupOfUnassignedDevicesFreezeForever**](CustomerApi.md#CustomerPrototypeDeleteGroupOfUnassignedDevicesFreezeForever) | **Delete** /Customers/{id}/locations/{locationId}/groupOfUnassignedDevices/freeze/forever | Delete GroupOfUnassignedDevices forever freeze for a Location ID.
[**CustomerPrototypeDeleteGroupOfUnassignedDevicesFreezeSuspend**](CustomerApi.md#CustomerPrototypeDeleteGroupOfUnassignedDevicesFreezeSuspend) | **Delete** /Customers/{id}/locations/{locationId}/groupOfUnassignedDevices/freeze/suspend | Delete GroupOfUnassignedDevices suspend for a Location ID.
[**CustomerPrototypeDeleteGroupOfUnassignedDevicesFreezeTemplateId**](CustomerApi.md#CustomerPrototypeDeleteGroupOfUnassignedDevicesFreezeTemplateId) | **Delete** /Customers/{id}/locations/{locationId}/groupOfUnassignedDevices/freeze/{freezeTemplateId} | Delete GroupOfUnassignedDevices uuid freeze for a Location ID.
[**CustomerPrototypeDeleteGroupOfUnassignedDevicesFreezes**](CustomerApi.md#CustomerPrototypeDeleteGroupOfUnassignedDevicesFreezes) | **Delete** /Customers/{id}/locations/{locationId}/groupOfUnassignedDevices/freezes | Delete All GroupOfUnassignedDevices freeze except autoExpire for a Location ID.
[**CustomerPrototypeDeleteKvConfigs**](CustomerApi.md#CustomerPrototypeDeleteKvConfigs) | **Delete** /Customers/{id}/locations/{locationId}/nodes/{nodeId}/kvConfigs/{module}/{key} | Delete kvConfigs with selected module and key on a particular Node for a Location ID.
[**CustomerPrototypeDeleteLinkedAccount**](CustomerApi.md#CustomerPrototypeDeleteLinkedAccount) | **Delete** /Customers/{id}/linkedAccounts/{provider}/{userId} | link the outside account, such as Samsung user.
[**CustomerPrototypeDeleteLocation**](CustomerApi.md#CustomerPrototypeDeleteLocation) | **Delete** /Customers/{id}/locations/{locationId} | Archive a location.
[**CustomerPrototypeDeleteLocationEventsHistory**](CustomerApi.md#CustomerPrototypeDeleteLocationEventsHistory) | **Delete** /Customers/{id}/locations/{locationId}/securityPolicy/events | Delete a Location&#39;s Security Events history for a location ID.
[**CustomerPrototypeDeleteLocationFreezeAutoExpire**](CustomerApi.md#CustomerPrototypeDeleteLocationFreezeAutoExpire) | **Delete** /Customers/{id}/locations/{locationId}/freeze/autoExpire | Delete the location freeze/autoExpire for a Location ID.
[**CustomerPrototypeDeleteManagerAccess**](CustomerApi.md#CustomerPrototypeDeleteManagerAccess) | **Delete** /Customers/{id}/locations/{locationId}/managers/{managerId} | Delete manager access for location and destroy access tokens for that manager\&quot;.
[**CustomerPrototypeDeleteNodeLocked**](CustomerApi.md#CustomerPrototypeDeleteNodeLocked) | **Delete** /Customers/{id}/nodes/{nodeId} | Delete a node model based on its id.
[**CustomerPrototypeDeleteNodePersistentConfigs**](CustomerApi.md#CustomerPrototypeDeleteNodePersistentConfigs) | **Delete** /Customers/{id}/locations/{locationId}/nodes/{nodeId}/persistentConfigs | Delete persistent data/configs from node in runtime.
[**CustomerPrototypeDeletePerson**](CustomerApi.md#CustomerPrototypeDeletePerson) | **Delete** /Customers/{id}/locations/{locationId}/persons/{personId} | Delete a Person for a location ID.
[**CustomerPrototypeDeletePersonAllFreeze**](CustomerApi.md#CustomerPrototypeDeletePersonAllFreeze) | **Delete** /Customers/{id}/locations/{locationId}/persons/{personId}/freezes | Delete a person to be frozen for a Location ID.
[**CustomerPrototypeDeletePersonFreeze**](CustomerApi.md#CustomerPrototypeDeletePersonFreeze) | **Delete** /Customers/{id}/locations/{locationId}/persons/{personId}/freeze/{freezeTemplateId} | Delete a person to be frozen for a Location ID.
[**CustomerPrototypeDeletePersonFreezeAutoExpire**](CustomerApi.md#CustomerPrototypeDeletePersonFreezeAutoExpire) | **Delete** /Customers/{id}/locations/{locationId}/persons/{personId}/freeze/autoExpire | Delete all devices from a person to be frozen for a Location ID.
[**CustomerPrototypeDeletePersonFreezeForever**](CustomerApi.md#CustomerPrototypeDeletePersonFreezeForever) | **Delete** /Customers/{id}/locations/{locationId}/persons/{personId}/freeze/forever | Delete a person forever freeze for a Location ID.
[**CustomerPrototypeDeletePersonFreezeSuspend**](CustomerApi.md#CustomerPrototypeDeletePersonFreezeSuspend) | **Delete** /Customers/{id}/locations/{locationId}/persons/{personId}/freeze/suspend | Delete person suspend for a Location ID.
[**CustomerPrototypeDeletePersonProfile**](CustomerApi.md#CustomerPrototypeDeletePersonProfile) | **Delete** /Customers/{id}/locations/{locationId}/persons/{personId}/profile | Delete a Person&#39;s Profile for a location ID.
[**CustomerPrototypeDeletePortForward**](CustomerApi.md#CustomerPrototypeDeletePortForward) | **Delete** /Customers/{id}/locations/{locationId}/networkConfiguration/dhcpReservations/{mac}/portForward/{externalPort} | Delete an existing Port Forwarding entry for an existing DHCP IP reservation tied to a MAC address at a Location ID.
[**CustomerPrototypeDeleteRemoteConnectionsAllow**](CustomerApi.md#CustomerPrototypeDeleteRemoteConnectionsAllow) | **Delete** /Customers/{id}/locations/{locationId}/devices/{mac}/securityPolicy/remoteConnections/allow/{ipaddr} | Delete a Remote Connection Allow IpAddress/ttl for the given device and Location ID.
[**CustomerPrototypeDeleteRemoteConnectionsAllowAll**](CustomerApi.md#CustomerPrototypeDeleteRemoteConnectionsAllowAll) | **Delete** /Customers/{id}/locations/{locationId}/devices/{mac}/securityPolicy/remoteConnections/allowAll | Delete a Remote Connection Allow All for the given device and Location ID.
[**CustomerPrototypeDeleteRoom**](CustomerApi.md#CustomerPrototypeDeleteRoom) | **Delete** /Customers/{id}/locations/{locationId}/rooms/{roomId} | Delete a Room for a location ID.
[**CustomerPrototypeDeleteWifiAccessZone**](CustomerApi.md#CustomerPrototypeDeleteWifiAccessZone) | **Delete** /Customers/{id}/locations/{locationId}/wifiNetwork/accessZones/{accessZoneId} | Delete a WiFi Access Zone
[**CustomerPrototypeDeleteWifiKey**](CustomerApi.md#CustomerPrototypeDeleteWifiKey) | **Delete** /Customers/{id}/locations/{locationId}/wifiNetwork/accessZones/{accessZone}/keys/{keyId} | Delete a WiFi Password
[**CustomerPrototypeDestroyMigration**](CustomerApi.md#CustomerPrototypeDestroyMigration) | **Delete** /Customers/{id}/_migration | Deletes _migration of this model.
[**CustomerPrototypeDisableLogin**](CustomerApi.md#CustomerPrototypeDisableLogin) | **Put** /Customers/{id}/disable | Disable customer from logging in until their account is reactivated.
[**CustomerPrototypeEnableDeviceTypeSniffing**](CustomerApi.md#CustomerPrototypeEnableDeviceTypeSniffing) | **Post** /Customers/{id}/locations/{locationId}/devices/{mac}/resniff | Re-enables deviceType sniffing for a particular device.
[**CustomerPrototypeEnableLogin**](CustomerApi.md#CustomerPrototypeEnableLogin) | **Put** /Customers/{id}/enable | Enable customer log in, after it has been disabled.
[**CustomerPrototypeFactoryReset**](CustomerApi.md#CustomerPrototypeFactoryReset) | **Delete** /Customers/{id}/locations/{locationId}/factoryReset | Reset specified location settings to default, while keeping claimed nodes intact. Some of the flags can cause a node to be reeboted.
[**CustomerPrototypeFindLocationById**](CustomerApi.md#CustomerPrototypeFindLocationById) | **Get** /Customers/{id}/locations/{locationId} | Get a Location&#39;s combined State and Config by LocationId.
[**CustomerPrototypeGetAccessTokenForManagedLocation**](CustomerApi.md#CustomerPrototypeGetAccessTokenForManagedLocation) | **Post** /Customers/{id}/entitledAccess/{locationId}/accessTokens | Get an access token for a location where you are assigned as a manager
[**CustomerPrototypeGetAlerts**](CustomerApi.md#CustomerPrototypeGetAlerts) | **Get** /Customers/{id}/locations/{locationId}/alerts | Retrieve active alerts for this location.
[**CustomerPrototypeGetAppEngagementTimer**](CustomerApi.md#CustomerPrototypeGetAppEngagementTimer) | **Get** /Customers/{id}/locations/{locationId}/appEngagementTimer | Get information about app engagement timer details for a location
[**CustomerPrototypeGetAppIdInfoCaptivePortalNetwork**](CustomerApi.md#CustomerPrototypeGetAppIdInfoCaptivePortalNetwork) | **Get** /Customers/{id}/locations/{locationId}/secondaryNetworks/captivePortals/ownerAppIds | Get the AppId info for the given location.
[**CustomerPrototypeGetAppPrioritizationLocationConfig**](CustomerApi.md#CustomerPrototypeGetAppPrioritizationLocationConfig) | **Get** /Customers/{id}/locations/{locationId}/qos/appPrioritization | Get status for app prioritization.
[**CustomerPrototypeGetAppPrioritizationMonitoring**](CustomerApi.md#CustomerPrototypeGetAppPrioritizationMonitoring) | **Post** /Customers/{id}/locations/{locationId}/qos/appPrioritization/monitoring | Get monitoring metrics for app prioritization.
[**CustomerPrototypeGetAppPrioritizationTemplateConfig**](CustomerApi.md#CustomerPrototypeGetAppPrioritizationTemplateConfig) | **Get** /Customers/{id}/locations/{locationId}/qos/appPrioritization/templateConfig | Get AppPrioritization template configs
[**CustomerPrototypeGetAppQoeStatsByTrafficClass**](CustomerApi.md#CustomerPrototypeGetAppQoeStatsByTrafficClass) | **Post** /Customers/{id}/locations/{locationId}/appqoe/AppQoeStatsByTrafficClass | Get App QoE metrics by traffic classes / devices / apps.
[**CustomerPrototypeGetAppQoeTrafficClassMetricsGetCustomersidLocationslocationIdAppqoeTrafficClassStats**](CustomerApi.md#CustomerPrototypeGetAppQoeTrafficClassMetricsGetCustomersidLocationslocationIdAppqoeTrafficClassStats) | **Get** /Customers/{id}/locations/{locationId}/appqoe/traffic_class_stats | Get App QoE metrics for traffic classes.
[**CustomerPrototypeGetAppQoeTrafficClassMetricsPostCustomersidLocationslocationIdAppqoeTrafficClassStats**](CustomerApi.md#CustomerPrototypeGetAppQoeTrafficClassMetricsPostCustomersidLocationslocationIdAppqoeTrafficClassStats) | **Post** /Customers/{id}/locations/{locationId}/appqoe/traffic_class_stats | Get App QoE metrics for traffic classes.
[**CustomerPrototypeGetAppTimeIpFlows**](CustomerApi.md#CustomerPrototypeGetAppTimeIpFlows) | **Get** /Customers/{id}/locations/{locationId}/appTime/ipFlows | Get IP flows config
[**CustomerPrototypeGetAuditTrailForCustomer**](CustomerApi.md#CustomerPrototypeGetAuditTrailForCustomer) | **Get** /Customers/{id}/auditTrail | Get audit trail for a customer.
[**CustomerPrototypeGetAuditTrailForLocation**](CustomerApi.md#CustomerPrototypeGetAuditTrailForLocation) | **Get** /Customers/{id}/locations/{locationId}/auditTrail | Get audit trail for location.
[**CustomerPrototypeGetAuthorizations**](CustomerApi.md#CustomerPrototypeGetAuthorizations) | **Get** /Customers/{id}/locations/{locationId}/authorizations | Get the number of authorized leaf pods for a Location ID.
[**CustomerPrototypeGetBackhaul**](CustomerApi.md#CustomerPrototypeGetBackhaul) | **Get** /Customers/{id}/locations/{locationId}/backhaul | 
[**CustomerPrototypeGetCampaignCaptivePortalNetwork**](CustomerApi.md#CustomerPrototypeGetCampaignCaptivePortalNetwork) | **Get** /Customers/{id}/locations/{locationId}/secondaryNetworks/captivePortals/{networkId}/campaign | Get the Captive Portal campaign for a given Location ID/NetworkId.
[**CustomerPrototypeGetCaptivePortalAuthorizedClients**](CustomerApi.md#CustomerPrototypeGetCaptivePortalAuthorizedClients) | **Get** /Customers/{id}/locations/{locationId}/secondaryNetworks/captivePortals/{networkId}/authorizedClients | Get CaptivePortal authorized clients
[**CustomerPrototypeGetCaptivePortalGuestEmailCollectionInfo**](CustomerApi.md#CustomerPrototypeGetCaptivePortalGuestEmailCollectionInfo) | **Get** /Customers/{id}/locations/{locationId}/secondaryNetworks/captivePortals/{networkId}/guestEmailCollectionInfo | Fetch the Captive Portal Network guest info download availability for the given network.
[**CustomerPrototypeGetCaptivePortalGuests**](CustomerApi.md#CustomerPrototypeGetCaptivePortalGuests) | **Get** /Customers/{id}/locations/{locationId}/secondaryNetworks/captivePortals/{networkId}/guests | Fetch the list of Guests which were logged into the given captivePortal network during the current day.
[**CustomerPrototypeGetCaptivePortalNetworks**](CustomerApi.md#CustomerPrototypeGetCaptivePortalNetworks) | **Get** /Customers/{id}/locations/{locationId}/secondaryNetworks/captivePortals | Get the Captive Portal configs for a given Location ID.
[**CustomerPrototypeGetCaptivePortalSendDetails**](CustomerApi.md#CustomerPrototypeGetCaptivePortalSendDetails) | **Get** /Customers/{id}/locations/{locationId}/secondaryNetworks/captivePortals/{networkId}/downloadGuestDetails | Download Captive Portal Guest details for a given Location ID/NetworkId.
[**CustomerPrototypeGetCaptivePortalSendDetailsDirect**](CustomerApi.md#CustomerPrototypeGetCaptivePortalSendDetailsDirect) | **Get** /Customers/{id}/locations/{locationId}/secondaryNetworks/captivePortals/{networkId}/downloadGuestDetailsDirect | Download Captive Portal Guest details for a given Location ID/NetworkId without accessing Amazon S3.
[**CustomerPrototypeGetCommands**](CustomerApi.md#CustomerPrototypeGetCommands) | **Get** /Customers/{id}/locations/{locationId}/command | Returns list of linked command accounts for the location
[**CustomerPrototypeGetCompanyInfoCaptivePortalNetwork**](CustomerApi.md#CustomerPrototypeGetCompanyInfoCaptivePortalNetwork) | **Post** /Customers/{id}/locations/{locationId}/secondaryNetworks/captivePortals/companyInfo/search | Get the companyInfo for the given url (domain).
[**CustomerPrototypeGetControlMode**](CustomerApi.md#CustomerPrototypeGetControlMode) | **Get** /Customers/{id}/locations/{locationId}/controlMode | Get control mode for a Location ID.
[**CustomerPrototypeGetCustomerNodeById**](CustomerApi.md#CustomerPrototypeGetCustomerNodeById) | **Get** /Customers/{id}/nodes/{nodeId} | Returns a single Node for a Customer ID.
[**CustomerPrototypeGetCustomerSupportConfigurations**](CustomerApi.md#CustomerPrototypeGetCustomerSupportConfigurations) | **Get** /Customers/{id}/locations/{locationId}/customerSupportConfigurations | Returns partner customer support configuration.
[**CustomerPrototypeGetDashboard**](CustomerApi.md#CustomerPrototypeGetDashboard) | **Get** /Customers/{id}/locations/{locationId}/flex/dashboard | Daily/Weekly/Monthly device usage summary report based on location
[**CustomerPrototypeGetDefaultNetworkAppTimeAppsDataUsage**](CustomerApi.md#CustomerPrototypeGetDefaultNetworkAppTimeAppsDataUsage) | **Get** /Customers/{id}/locations/{locationId}/wifiNetwork/appTime/apps/dataUsage | Fetch the AppTime Apps Data Usage Stats for captivePortal network.
[**CustomerPrototypeGetDefaultNetworkAppTimeAppsOnlineTime**](CustomerApi.md#CustomerPrototypeGetDefaultNetworkAppTimeAppsOnlineTime) | **Get** /Customers/{id}/locations/{locationId}/wifiNetwork/appTime/apps/onlineTime | Fetch the AppTime Apps Online Time Stats for default network.
[**CustomerPrototypeGetDefaultNetworkAppTimeCategoriesDataUsage**](CustomerApi.md#CustomerPrototypeGetDefaultNetworkAppTimeCategoriesDataUsage) | **Get** /Customers/{id}/locations/{locationId}/wifiNetwork/appTime/categories/dataUsage | Fetch the AppTime Categories Data Usage Stats for captivePortal network.
[**CustomerPrototypeGetDefaultNetworkAppTimeCategoriesOnlineTime**](CustomerApi.md#CustomerPrototypeGetDefaultNetworkAppTimeCategoriesOnlineTime) | **Get** /Customers/{id}/locations/{locationId}/wifiNetwork/appTime/categories/onlineTime | Fetch the AppTime Categories Online Time Stats for captivePortal network.
[**CustomerPrototypeGetDeviceAlarms**](CustomerApi.md#CustomerPrototypeGetDeviceAlarms) | **Get** /Customers/{id}/locations/{locationId}/flex/devices/{mac}/alarms | Device alarm history graph array for a particular MAC address.
[**CustomerPrototypeGetDeviceAppTimeAppsDataUsage**](CustomerApi.md#CustomerPrototypeGetDeviceAppTimeAppsDataUsage) | **Get** /Customers/{id}/locations/{locationId}/devices/{mac}/appTime/apps/dataUsage | Fetch the AppTime Apps Data Usage Stats for a Device.
[**CustomerPrototypeGetDeviceAppTimeAppsOnlineTime**](CustomerApi.md#CustomerPrototypeGetDeviceAppTimeAppsOnlineTime) | **Get** /Customers/{id}/locations/{locationId}/devices/{mac}/appTime/apps/onlineTime | Fetch the AppTime Apps Online Time Stats for a Device.
[**CustomerPrototypeGetDeviceAppTimeCategoriesDataUsage**](CustomerApi.md#CustomerPrototypeGetDeviceAppTimeCategoriesDataUsage) | **Get** /Customers/{id}/locations/{locationId}/devices/{mac}/appTime/categories/dataUsage | Fetch the AppTime Categories Data Usage Stats for a Device.
[**CustomerPrototypeGetDeviceAppTimeCategoriesOnlineTime**](CustomerApi.md#CustomerPrototypeGetDeviceAppTimeCategoriesOnlineTime) | **Get** /Customers/{id}/locations/{locationId}/devices/{mac}/appTime/categories/onlineTime | Fetch the AppTime Categories Online Time Stats for a Device.
[**CustomerPrototypeGetDeviceBandSteeringStats**](CustomerApi.md#CustomerPrototypeGetDeviceBandSteeringStats) | **Get** /Customers/{id}/locations/{locationId}/flex/devices/{mac}/bandSteeringStats | Device band steering stats with all nodes for a particular MAC address.
[**CustomerPrototypeGetDeviceByMac**](CustomerApi.md#CustomerPrototypeGetDeviceByMac) | **Get** /Customers/{id}/locations/{locationId}/devices/{mac} | Returns a single Device for a Location ID.
[**CustomerPrototypeGetDeviceClientSteeringStats**](CustomerApi.md#CustomerPrototypeGetDeviceClientSteeringStats) | **Get** /Customers/{id}/locations/{locationId}/flex/devices/{mac}/clientSteeringStats | Device client steering stats with all nodes for a particular MAC address.
[**CustomerPrototypeGetDeviceGroups**](CustomerApi.md#CustomerPrototypeGetDeviceGroups) | **Get** /Customers/{id}/locations/{locationId}/networkAccess/networks/{networkId}/deviceGroups | Get a list of device groups in a network, along with a list of member devices and group shares.
[**CustomerPrototypeGetDeviceQoeMetrics**](CustomerApi.md#CustomerPrototypeGetDeviceQoeMetrics) | **Get** /Customers/{id}/locations/{locationId}/flex/devices/{mac}/qoeMetrics | Device or pod QoE 15 minutes data.
[**CustomerPrototypeGetDeviceSecurity**](CustomerApi.md#CustomerPrototypeGetDeviceSecurity) | **Get** /Customers/{id}/locations/{locationId}/devices/{mac}/securityPolicy | Returns the security policy Device for a Location ID.
[**CustomerPrototypeGetDeviceSecurityPolicyEvents**](CustomerApi.md#CustomerPrototypeGetDeviceSecurityPolicyEvents) | **Get** /Customers/{id}/locations/{locationId}/devices/{mac}/securityPolicy/events | Get a Security Policy Events for Device for a Location ID.
[**CustomerPrototypeGetDeviceSecurityPolicyHourlyCounts**](CustomerApi.md#CustomerPrototypeGetDeviceSecurityPolicyHourlyCounts) | **Get** /Customers/{id}/locations/{locationId}/devices/{mac}/securityPolicy/hourlyBlockedCounts | Get a Security Policy Hourly Blocked Counts for a Device for a Location ID.
[**CustomerPrototypeGetDeviceSoundingState**](CustomerApi.md#CustomerPrototypeGetDeviceSoundingState) | **Get** /Customers/{id}/locations/{locationId}/homeSecurity/devices/sounding | Fetch the sounding states for eligible devices in this location
[**CustomerPrototypeGetDeviceSteeringWithAthena**](CustomerApi.md#CustomerPrototypeGetDeviceSteeringWithAthena) | **Get** /Customers/{id}/locations/{locationId}/flex/devices/{mac}/clientSteeringTriggers | Find all instances of the model.
[**CustomerPrototypeGetDevices**](CustomerApi.md#CustomerPrototypeGetDevices) | **Get** /Customers/{id}/locations/{locationId}/devices | Get all the devices for a Location ID, including the device name, icon to use, MAC and IP  address, connecting nodes and more.
[**CustomerPrototypeGetDhcp**](CustomerApi.md#CustomerPrototypeGetDhcp) | **Get** /Customers/{id}/locations/{locationId}/networkConfiguration/dhcp | Get current DHCP Configuration details for a Location ID.
[**CustomerPrototypeGetDhcpReservation**](CustomerApi.md#CustomerPrototypeGetDhcpReservation) | **Get** /Customers/{id}/locations/{locationId}/networkConfiguration/dhcpReservations/{mac} | Get current DHCP IP reservation details for a Location ID.
[**CustomerPrototypeGetDhcpReservations**](CustomerApi.md#CustomerPrototypeGetDhcpReservations) | **Get** /Customers/{id}/locations/{locationId}/networkConfiguration/dhcpReservations | Get current DHCP IP reservation details for a Location ID.
[**CustomerPrototypeGetDnsServers**](CustomerApi.md#CustomerPrototypeGetDnsServers) | **Get** /Customers/{id}/locations/{locationId}/networkConfiguration/dnsServers | Get the current DNS IP addresses and settings for a Location ID.
[**CustomerPrototypeGetDpp**](CustomerApi.md#CustomerPrototypeGetDpp) | **Get** /Customers/{id}/locations/{locationId}/dpp | Get the current DPP configuration for a Location ID.
[**CustomerPrototypeGetDppAnnouncementsFromController**](CustomerApi.md#CustomerPrototypeGetDppAnnouncementsFromController) | **Get** /Customers/{id}/locations/{locationId}/dpp/announcements | Returns DPP announcements from controller
[**CustomerPrototypeGetEmployeeNetworkAppTimeAppsDataUsage**](CustomerApi.md#CustomerPrototypeGetEmployeeNetworkAppTimeAppsDataUsage) | **Get** /Customers/{id}/locations/{locationId}/secondaryNetworks/fronthauls/{networkId}/appTime/apps/dataUsage | Fetch the AppTime Apps Data Usage Stats for fronthaul network.
[**CustomerPrototypeGetEmployeeNetworkAppTimeAppsOnlineTime**](CustomerApi.md#CustomerPrototypeGetEmployeeNetworkAppTimeAppsOnlineTime) | **Get** /Customers/{id}/locations/{locationId}/secondaryNetworks/fronthauls/{networkId}/appTime/apps/onlineTime | Fetch the AppTime Apps Online Time Stats for fronthaul network.
[**CustomerPrototypeGetEmployeeNetworkAppTimeCategoriesDataUsage**](CustomerApi.md#CustomerPrototypeGetEmployeeNetworkAppTimeCategoriesDataUsage) | **Get** /Customers/{id}/locations/{locationId}/secondaryNetworks/fronthauls/{networkId}/appTime/categories/dataUsage | Fetch the AppTime Categories Data Usage Stats for fronthaul network.
[**CustomerPrototypeGetEmployeeNetworkAppTimeCategoriesOnlineTime**](CustomerApi.md#CustomerPrototypeGetEmployeeNetworkAppTimeCategoriesOnlineTime) | **Get** /Customers/{id}/locations/{locationId}/secondaryNetworks/fronthauls/{networkId}/appTime/categories/onlineTime | Fetch the AppTime Categories Online Time Stats for fronthaul network.
[**CustomerPrototypeGetEntitledAccessList**](CustomerApi.md#CustomerPrototypeGetEntitledAccessList) | **Get** /Customers/{id}/entitledAccess | Get a list of all locations on which you are assigned as a manager.
[**CustomerPrototypeGetEventHistory**](CustomerApi.md#CustomerPrototypeGetEventHistory) | **Get** /Customers/{id}/locations/{locationId}/homeSecurity/events/history | Fetch the event history for this location
[**CustomerPrototypeGetFastInterference**](CustomerApi.md#CustomerPrototypeGetFastInterference) | **Get** /Customers/{id}/locations/{locationId}/fastInterference | Get from Controller Fast interference status.
[**CustomerPrototypeGetFirmwareUpgradeStatus**](CustomerApi.md#CustomerPrototypeGetFirmwareUpgradeStatus) | **Get** /Customers/{id}/locations/{locationId}/firmware | Firmware Upgrade Status
[**CustomerPrototypeGetFlowStats**](CustomerApi.md#CustomerPrototypeGetFlowStats) | **Get** /Customers/{id}/locations/{locationId}/flowStats | GET the flow stats configuration
[**CustomerPrototypeGetForceGraph**](CustomerApi.md#CustomerPrototypeGetForceGraph) | **Get** /Customers/{id}/locations/{locationId}/forceGraph | HTML or JSON (vertices[] + edges[]) used to display a Network Topology.
[**CustomerPrototypeGetFrontHaulNetworks**](CustomerApi.md#CustomerPrototypeGetFrontHaulNetworks) | **Get** /Customers/{id}/locations/{locationId}/secondaryNetworks/fronthauls | Get the Front Haul Portal configs for a given Location ID.
[**CustomerPrototypeGetFrontHaulsDpp**](CustomerApi.md#CustomerPrototypeGetFrontHaulsDpp) | **Get** /Customers/{id}/locations/{locationId}/secondaryNetworks/fronthauls/{networkId}/dpp | Get the current DPP configurator for a Location ID.
[**CustomerPrototypeGetFrontlineStorage**](CustomerApi.md#CustomerPrototypeGetFrontlineStorage) | **Get** /Customers/{id}/locations/{locationId}/frontline/storage | Fetch the frontline storage for this location
[**CustomerPrototypeGetGdprCaptivePortalsData**](CustomerApi.md#CustomerPrototypeGetGdprCaptivePortalsData) | **Get** /Customers/{id}/locations/{locationId}/networks/{networkId}/gdprData | Fetch the Gdpr Captive Portals data for a guest.
[**CustomerPrototypeGetGroupOfUnassignedDevicesFreezePolicy**](CustomerApi.md#CustomerPrototypeGetGroupOfUnassignedDevicesFreezePolicy) | **Get** /Customers/{id}/locations/{locationId}/groupOfUnassignedDevices/freezePolicy | Get GroupOfUnassignedDevices freeze policy for a Location ID.
[**CustomerPrototypeGetGroupOfUnassignedDevicesSecurityPolicy**](CustomerApi.md#CustomerPrototypeGetGroupOfUnassignedDevicesSecurityPolicy) | **Get** /Customers/{id}/locations/{locationId}/groupOfUnassignedDevices/securityPolicy | Get a Security Policy for a Location ID.
[**CustomerPrototypeGetGroupOfUnassignedDevicesSecurityPolicyEvents**](CustomerApi.md#CustomerPrototypeGetGroupOfUnassignedDevicesSecurityPolicyEvents) | **Get** /Customers/{id}/locations/{locationId}/groupOfUnassignedDevices/securityPolicy/events | Get a Security Policy Events for groupOfUnassignedDevices for a Location ID.
[**CustomerPrototypeGetGroupOfUnassignedDevicesSecurityPolicyHourlyCounts**](CustomerApi.md#CustomerPrototypeGetGroupOfUnassignedDevicesSecurityPolicyHourlyCounts) | **Get** /Customers/{id}/locations/{locationId}/groupOfUnassignedDevices/securityPolicy/hourlyBlockedCounts | Get a Security Policy Hourly Blocked Counts for group Of Unassigned Devices for a Location ID.
[**CustomerPrototypeGetGroups**](CustomerApi.md#CustomerPrototypeGetGroups) | **Get** /Customers/{id}/groups | 
[**CustomerPrototypeGetGuestNetworkAppTimeAppsDataUsage**](CustomerApi.md#CustomerPrototypeGetGuestNetworkAppTimeAppsDataUsage) | **Get** /Customers/{id}/locations/{locationId}/secondaryNetworks/captivePortals/{networkId}/appTime/apps/dataUsage | Fetch the AppTime Apps Data Usage Stats for captivePortal network.
[**CustomerPrototypeGetGuestNetworkAppTimeAppsOnlineTime**](CustomerApi.md#CustomerPrototypeGetGuestNetworkAppTimeAppsOnlineTime) | **Get** /Customers/{id}/locations/{locationId}/secondaryNetworks/captivePortals/{networkId}/appTime/apps/onlineTime | Fetch the AppTime Apps Online Time Stats for captivePortal network.
[**CustomerPrototypeGetGuestNetworkAppTimeCategoriesDataUsage**](CustomerApi.md#CustomerPrototypeGetGuestNetworkAppTimeCategoriesDataUsage) | **Get** /Customers/{id}/locations/{locationId}/secondaryNetworks/captivePortals/{networkId}/appTime/categories/dataUsage | Fetch the AppTime Categories Data Usage Stats for captivePortal network.
[**CustomerPrototypeGetGuestNetworkAppTimeCategoriesOnlineTime**](CustomerApi.md#CustomerPrototypeGetGuestNetworkAppTimeCategoriesOnlineTime) | **Get** /Customers/{id}/locations/{locationId}/secondaryNetworks/captivePortals/{networkId}/appTime/categories/onlineTime | Fetch the AppTime Categories Online Time Stats for captivePortal network.
[**CustomerPrototypeGetHomeAwayLocationEvents**](CustomerApi.md#CustomerPrototypeGetHomeAwayLocationEvents) | **Get** /Customers/{id}/locations/{locationId}/homeAway/events | Fetch the all the Homeaway events history for this location
[**CustomerPrototypeGetHomeSecurity**](CustomerApi.md#CustomerPrototypeGetHomeSecurity) | **Get** /Customers/{id}/locations/{locationId}/homeSecurity | Fetch the home security configuration for this location
[**CustomerPrototypeGetIPv6**](CustomerApi.md#CustomerPrototypeGetIPv6) | **Get** /Customers/{id}/locations/{locationId}/ipv6 | 
[**CustomerPrototypeGetKvConfigs**](CustomerApi.md#CustomerPrototypeGetKvConfigs) | **Get** /Customers/{id}/locations/{locationId}/nodes/{nodeId}/kvConfigs | Retrieve all kvConfigs on a particular Node for a Location ID.
[**CustomerPrototypeGetLocationAppTime**](CustomerApi.md#CustomerPrototypeGetLocationAppTime) | **Get** /Customers/{id}/locations/{locationId}/appTime | Get a Location&#39;s AppTime config by location ID.
[**CustomerPrototypeGetLocationCapabilities**](CustomerApi.md#CustomerPrototypeGetLocationCapabilities) | **Get** /Customers/{id}/locations/{locationId}/capabilities | Get the non-base feature capabilities supported by a particular Location ID.
[**CustomerPrototypeGetLocationConfigAuditEvents**](CustomerApi.md#CustomerPrototypeGetLocationConfigAuditEvents) | **Get** /Customers/{id}/locations/{locationId}/configAudit/events | Get a Config Audit Trail Events for a Location ID.
[**CustomerPrototypeGetLocationFreezeAutoExpire**](CustomerApi.md#CustomerPrototypeGetLocationFreezeAutoExpire) | **Get** /Customers/{id}/locations/{locationId}/freeze/autoExpire | Get all devices/persons except some to be frozen for a Location ID.
[**CustomerPrototypeGetLocationGuardEventStats**](CustomerApi.md#CustomerPrototypeGetLocationGuardEventStats) | **Post** /Customers/{id}/locations/{locationId}/securityPolicy/guard/eventStats | Get the Guard Event Stats for a Location ID.
[**CustomerPrototypeGetLocationGuardEventsTldOrIp**](CustomerApi.md#CustomerPrototypeGetLocationGuardEventsTldOrIp) | **Post** /Customers/{id}/locations/{locationId}/securityPolicy/guard/events | Get the Guard Event Domain Groups for a Location ID.
[**CustomerPrototypeGetLocationGuardPersonEventsSummary**](CustomerApi.md#CustomerPrototypeGetLocationGuardPersonEventsSummary) | **Get** /Customers/{id}/locations/{locationId}/securityPolicy/guard/personEventsSummary | Get the Guard Event Stats for all persons in a Location ID.
[**CustomerPrototypeGetLocationKvStates**](CustomerApi.md#CustomerPrototypeGetLocationKvStates) | **Get** /Customers/{id}/locations/{locationId}/kvStates | Retrieve all kvStates on a particular Node for a Location ID.
[**CustomerPrototypeGetLocationOverlordConfigs**](CustomerApi.md#CustomerPrototypeGetLocationOverlordConfigs) | **Get** /Customers/{id}/locations/{locationId}/v2/configAndState | Gets all the configs from Overlord for a specified location.
[**CustomerPrototypeGetLocationQoe**](CustomerApi.md#CustomerPrototypeGetLocationQoe) | **Get** /Customers/{id}/locations/{locationId}/flex/qoe | Get QoE recent 1 minute data for a whole location.
[**CustomerPrototypeGetLocationRooms**](CustomerApi.md#CustomerPrototypeGetLocationRooms) | **Get** /Customers/{id}/locations/{locationId}/rooms | Get a Location&#39;s Rooms config by location ID.
[**CustomerPrototypeGetLocationSecurityPolicy**](CustomerApi.md#CustomerPrototypeGetLocationSecurityPolicy) | **Get** /Customers/{id}/locations/{locationId}/securityPolicy | Get a Security Policy for a Location ID.
[**CustomerPrototypeGetLocationSecurityPolicyEvents**](CustomerApi.md#CustomerPrototypeGetLocationSecurityPolicyEvents) | **Get** /Customers/{id}/locations/{locationId}/securityPolicy/events | Get a Security Policy Events for a Location ID.
[**CustomerPrototypeGetLocationSecurityPolicyHourlyCounts**](CustomerApi.md#CustomerPrototypeGetLocationSecurityPolicyHourlyCounts) | **Get** /Customers/{id}/locations/{locationId}/securityPolicy/hourlyBlockedCounts | Get a Security Policy Hourly Blocked Counts for a Location ID.
[**CustomerPrototypeGetLocationWanConfiguration**](CustomerApi.md#CustomerPrototypeGetLocationWanConfiguration) | **Get** /Customers/{id}/locations/{locationId}/wanConfiguration | Get WAN Configuration for a Location ID.
[**CustomerPrototypeGetLocationWanSettings**](CustomerApi.md#CustomerPrototypeGetLocationWanSettings) | **Get** /Customers/{id}/locations/{locationId}/wanSettings | DEPRECATED: Get the WAN Settings for a Location ID.
[**CustomerPrototypeGetLocations**](CustomerApi.md#CustomerPrototypeGetLocations) | **Get** /Customers/{id}/locations | Queries locations of Customer.
[**CustomerPrototypeGetManagersListForLocation**](CustomerApi.md#CustomerPrototypeGetManagersListForLocation) | **Get** /Customers/{id}/locations/{locationId}/managers | Get a list of all managers the are assigned to manage your location.
[**CustomerPrototypeGetMigration**](CustomerApi.md#CustomerPrototypeGetMigration) | **Get** /Customers/{id}/_migration | Fetches hasOne relation _migration.
[**CustomerPrototypeGetModulesFromController**](CustomerApi.md#CustomerPrototypeGetModulesFromController) | **Get** /Customers/{id}/locations/{locationId}/firmware/modules | Retrieve all firmaware modules for a Location ID.
[**CustomerPrototypeGetMotionHistory**](CustomerApi.md#CustomerPrototypeGetMotionHistory) | **Get** /Customers/{id}/locations/{locationId}/homeSecurity/motionHistory | Fetch the motion density history for this location
[**CustomerPrototypeGetMotionStateHistory**](CustomerApi.md#CustomerPrototypeGetMotionStateHistory) | **Get** /Customers/{id}/locations/{locationId}/homeSecurity/motionHistory/state | Fetch the motion state history for this location
[**CustomerPrototypeGetNetworkAccessNetworks**](CustomerApi.md#CustomerPrototypeGetNetworkAccessNetworks) | **Get** /Customers/{id}/locations/{locationId}/networkAccess/networks | Get information about networkAccess networks
[**CustomerPrototypeGetNetworkConfigurationHome**](CustomerApi.md#CustomerPrototypeGetNetworkConfigurationHome) | **Get** /Customers/{id}/locations/{locationId}/networkConfiguration/home | Get the current overall settings and status of the Advanced Networking settings for a Location ID.
[**CustomerPrototypeGetNetworkMode**](CustomerApi.md#CustomerPrototypeGetNetworkMode) | **Get** /Customers/{id}/locations/{locationId}/networkMode | Get the current Network Mode for a Location ID.
[**CustomerPrototypeGetNodeBlePairingPin**](CustomerApi.md#CustomerPrototypeGetNodeBlePairingPin) | **Get** /Customers/{id}/locations/{locationId}/nodes/{nodeId}/blePairingPin | Get BLE pairing pin for a node that is claimed by the selected location
[**CustomerPrototypeGetNodeBySerialNumber**](CustomerApi.md#CustomerPrototypeGetNodeBySerialNumber) | **Get** /Customers/{id}/locations/{locationId}/nodes/{nodeId} | Returns a single Node for a Location ID with its list of connected devices.
[**CustomerPrototypeGetNodeKvStates**](CustomerApi.md#CustomerPrototypeGetNodeKvStates) | **Get** /Customers/{id}/locations/{locationId}/nodes/{nodeId}/kvStates | Retrieve all kvStates on a particular Node for a Location ID.
[**CustomerPrototypeGetNodes**](CustomerApi.md#CustomerPrototypeGetNodes) | **Get** /Customers/{id}/locations/{locationId}/nodes | Retrieve the Node settings and status for a Location ID.
[**CustomerPrototypeGetOhpLocationIdentifier**](CustomerApi.md#CustomerPrototypeGetOhpLocationIdentifier) | **Get** /Customers/{id}/locations/{locationId}/securityPolicy/ohp/locationIdentifier | Get the current OHP identifier for a Location ID.
[**CustomerPrototypeGetOnboardingLocationIdentifier**](CustomerApi.md#CustomerPrototypeGetOnboardingLocationIdentifier) | **Get** /Customers/{id}/locations/{locationId}/onboardingLocationIdentifier | Get the onboarding identifier for a Location ID.
[**CustomerPrototypeGetPersonAppTimeAppsDataUsage**](CustomerApi.md#CustomerPrototypeGetPersonAppTimeAppsDataUsage) | **Get** /Customers/{id}/locations/{locationId}/persons/{personId}/appTime/apps/dataUsage | Fetch the AppTime Apps Data Usage Stats for a Person.
[**CustomerPrototypeGetPersonAppTimeAppsOnlineTime**](CustomerApi.md#CustomerPrototypeGetPersonAppTimeAppsOnlineTime) | **Get** /Customers/{id}/locations/{locationId}/persons/{personId}/appTime/apps/onlineTime | Fetch the AppTime Apps Online Time Stats for a Person.
[**CustomerPrototypeGetPersonAppTimeCategoriesDataUsage**](CustomerApi.md#CustomerPrototypeGetPersonAppTimeCategoriesDataUsage) | **Get** /Customers/{id}/locations/{locationId}/persons/{personId}/appTime/categories/dataUsage | Fetch the AppTime Categories Data Usage Stats for a Person.
[**CustomerPrototypeGetPersonAppTimeCategoriesOnlineTime**](CustomerApi.md#CustomerPrototypeGetPersonAppTimeCategoriesOnlineTime) | **Get** /Customers/{id}/locations/{locationId}/persons/{personId}/appTime/categories/onlineTime | Fetch the AppTime Categories Online Time Stats for a Person.
[**CustomerPrototypeGetPersonById**](CustomerApi.md#CustomerPrototypeGetPersonById) | **Get** /Customers/{id}/locations/{locationId}/persons/{personId} | Get a Person by ID for a Location ID.
[**CustomerPrototypeGetPersonSecurityPolicyEvents**](CustomerApi.md#CustomerPrototypeGetPersonSecurityPolicyEvents) | **Get** /Customers/{id}/locations/{locationId}/persons/{personId}/securityPolicy/events | Get a Security Policy Events for Person for a Location ID.
[**CustomerPrototypeGetPersonSecurityPolicyHourlyCounts**](CustomerApi.md#CustomerPrototypeGetPersonSecurityPolicyHourlyCounts) | **Get** /Customers/{id}/locations/{locationId}/persons/{personId}/securityPolicy/hourlyBlockedCounts | Get a Security Policy Hourly Blocked Counts for a Person for a Location ID.
[**CustomerPrototypeGetPersons**](CustomerApi.md#CustomerPrototypeGetPersons) | **Get** /Customers/{id}/locations/{locationId}/persons | Get all Persons for a Location ID.
[**CustomerPrototypeGetPortForwards**](CustomerApi.md#CustomerPrototypeGetPortForwards) | **Get** /Customers/{id}/locations/{locationId}/networkConfiguration/dhcpReservations/{mac}/portForwards | Get all existing Port Forwarding entries for an existing DHCP IP reservation tied to a MAC address at a Location ID.
[**CustomerPrototypeGetPrimarySecondaryNetworks**](CustomerApi.md#CustomerPrototypeGetPrimarySecondaryNetworks) | **Get** /Customers/{id}/locations/{locationId}/primarySecondaryNetworks | Get networks for wpa3 transition flow
[**CustomerPrototypeGetQoe1Minute**](CustomerApi.md#CustomerPrototypeGetQoe1Minute) | **Get** /Customers/{id}/locations/{locationId}/flex/devices/{mac}/qoe/liveModeStream | Device or pod QoE live mode data.
[**CustomerPrototypeGetQoeSeconds**](CustomerApi.md#CustomerPrototypeGetQoeSeconds) | **Get** /Customers/{id}/locations/{locationId}/flex/devices/{mac}/qoe/superLiveModeStream | Device or pod QoE super live mode data.
[**CustomerPrototypeGetRemoteConnectionsConfig**](CustomerApi.md#CustomerPrototypeGetRemoteConnectionsConfig) | **Get** /Customers/{id}/locations/{locationId}/securityPolicy/remoteConnections | Get the Unauthorized Remote Connections config for a Location ID.
[**CustomerPrototypeGetRoles**](CustomerApi.md#CustomerPrototypeGetRoles) | **Get** /Customers/{id}/roles | Queries roles of Customer.
[**CustomerPrototypeGetSecondaryNetworkInvitation**](CustomerApi.md#CustomerPrototypeGetSecondaryNetworkInvitation) | **Post** /Customers/{id}/locations/{locationId}/primarySecondaryNetworks/wpa3ssid/invitations | Update home devices visible to guests.
[**CustomerPrototypeGetSecurityRealizedStates**](CustomerApi.md#CustomerPrototypeGetSecurityRealizedStates) | **Get** /Customers/{id}/locations/{locationId}/securityPolicy/realizedState | Retrieve all securityStates for a Location ID.
[**CustomerPrototypeGetServiceLevel**](CustomerApi.md#CustomerPrototypeGetServiceLevel) | **Get** /Customers/{id}/locations/{locationId}/serviceLevel | Get the service level for this location
[**CustomerPrototypeGetSniffing**](CustomerApi.md#CustomerPrototypeGetSniffing) | **Get** /Customers/{id}/locations/{locationId}/sniffing | Get DNS, HTTP, UPnP and mDNS sniffing toggles for a Location ID.
[**CustomerPrototypeGetSpeedTestResults**](CustomerApi.md#CustomerPrototypeGetSpeedTestResults) | **Get** /Customers/{id}/locations/{locationId}/nodes/{nodeId}/speedTestResults | retrieve the speed test result for a node.
[**CustomerPrototypeGetSpeedTestResultsByRequestId**](CustomerApi.md#CustomerPrototypeGetSpeedTestResultsByRequestId) | **Get** /Customers/{id}/locations/{locationId}/nodes/{nodeId}/speedTestResults/{requestId} | retrieve single speed test result by request id for a node.
[**CustomerPrototypeGetSpeedTestResultsForApp**](CustomerApi.md#CustomerPrototypeGetSpeedTestResultsForApp) | **Get** /Customers/{id}/locations/{locationId}/appFacade/dashboard | Get the current speed test aggregation result for a Location ID.
[**CustomerPrototypeGetSsid**](CustomerApi.md#CustomerPrototypeGetSsid) | **Get** /Customers/{id}/locations/{locationId}/wifiNetwork/ssid | Get the current WiFi SSID for a Location ID.
[**CustomerPrototypeGetSubscription**](CustomerApi.md#CustomerPrototypeGetSubscription) | **Get** /Customers/{id}/locations/{locationId}/subscription | Get Subscription details for this location
[**CustomerPrototypeGetSummary**](CustomerApi.md#CustomerPrototypeGetSummary) | **Get** /Customers/{id}/locations/{locationId}/summary | DEPRECATED: The system summary for a location including topology, optimizations, and firmware upgrades.
[**CustomerPrototypeGetTaskStatuses**](CustomerApi.md#CustomerPrototypeGetTaskStatuses) | **Get** /Customers/{id}/locations/{locationId}/taskStatuses | Retrieve all task statuses of nodes from controller
[**CustomerPrototypeGetTermsAndPrivacyAccepted**](CustomerApi.md#CustomerPrototypeGetTermsAndPrivacyAccepted) | **Get** /Customers/{id}/termsAndPrivacyAccepted | Fetches hasOne relation termsAndPrivacyAccepted.
[**CustomerPrototypeGetTopology**](CustomerApi.md#CustomerPrototypeGetTopology) | **Get** /Customers/{id}/locations/{locationId}/topology | DEPRECATED: The topology for a location including channels and devices.
[**CustomerPrototypeGetTos**](CustomerApi.md#CustomerPrototypeGetTos) | **Get** /Customers/{id}/locations/{locationId}/devices/{mac}/tos | Describes the current state of TOS for the given client.
[**CustomerPrototypeGetUpnp**](CustomerApi.md#CustomerPrototypeGetUpnp) | **Get** /Customers/{id}/locations/{locationId}/networkConfiguration/upnp | Get the current UPnP setting for a Location ID.
[**CustomerPrototypeGetVapStates**](CustomerApi.md#CustomerPrototypeGetVapStates) | **Get** /Customers/{id}/locations/{locationId}/vapStates | Retrieve all Vap State on a particular Node for a Location ID.
[**CustomerPrototypeGetVapsAndStaStatesFromBackhaul**](CustomerApi.md#CustomerPrototypeGetVapsAndStaStatesFromBackhaul) | **Get** /Customers/{id}/locations/{locationId}/backhauls | Retrieve all Vap State on a particular Node for a Location ID.
[**CustomerPrototypeGetWhitelistApprovalRequests**](CustomerApi.md#CustomerPrototypeGetWhitelistApprovalRequests) | **Get** /Customers/{id}/locations/{locationId}/securityPolicy/websites/whitelist/approvalRequests | Get a list of pending approval requests for this location.
[**CustomerPrototypeGetWifiDashboard**](CustomerApi.md#CustomerPrototypeGetWifiDashboard) | **Get** /Customers/{id}/locations/{locationId}/appFacade/wifiDashboard | WiFi Dashboard
[**CustomerPrototypeGetWifiInvitationById**](CustomerApi.md#CustomerPrototypeGetWifiInvitationById) | **Post** /Customers/{id}/locations/{locationId}/wifiNetwork/accessZones/{zoneId}/keys/{keyId}/invitations | Update home devices visible to guests.
[**CustomerPrototypeGetWifiMotion**](CustomerApi.md#CustomerPrototypeGetWifiMotion) | **Get** /Customers/{id}/locations/{locationId}/wifiMotion | Get WifiMotion config for this location
[**CustomerPrototypeGetWifiNetwork**](CustomerApi.md#CustomerPrototypeGetWifiNetwork) | **Get** /Customers/{id}/locations/{locationId}/wifiNetwork | Get the current WiFi SSID and PSK for a Location ID.
[**CustomerPrototypeGetWifiNetworkDpp**](CustomerApi.md#CustomerPrototypeGetWifiNetworkDpp) | **Get** /Customers/{id}/locations/{locationId}/wifiNetwork/dpp | Get the current DPP configurator for a Location ID.
[**CustomerPrototypeGetWifiNetworks**](CustomerApi.md#CustomerPrototypeGetWifiNetworks) | **Get** /Customers/{id}/locations/{locationId}/wifiNetworks | WiFi Networks
[**CustomerPrototypeHasLocationById**](CustomerApi.md#CustomerPrototypeHasLocationById) | **Head** /Customers/{id}/locations/{locationId} | Verify that a Customer Id has a Location Id.
[**CustomerPrototypeIosDeviceTokenExists**](CustomerApi.md#CustomerPrototypeIosDeviceTokenExists) | **Get** /Customers/{id}/iosDeviceTokens/{deviceToken}/exists | Provides feedback as to whether an iOS deviceToken was previously registered for push notifications.
[**CustomerPrototypeLinkAccount**](CustomerApi.md#CustomerPrototypeLinkAccount) | **Post** /Customers/{id}/linkedAccounts | link the outside account, such as Samsung user.
[**CustomerPrototypeListCustomSharedSchedules**](CustomerApi.md#CustomerPrototypeListCustomSharedSchedules) | **Get** /Customers/{id}/locations/{locationId}/schedules | Get custom shared schedules for a given Location ID.
[**CustomerPrototypeMarketingExport**](CustomerApi.md#CustomerPrototypeMarketingExport) | **Get** /Customers/{id}/locations/{locationId}/marketingExport | Get detailed information of a location for CRM campaigns.
[**CustomerPrototypeMigrationStatus**](CustomerApi.md#CustomerPrototypeMigrationStatus) | **Get** /Customers/{id}/migration | Returns cloud migration status for customer
[**CustomerPrototypeOptimize**](CustomerApi.md#CustomerPrototypeOptimize) | **Post** /Customers/{id}/locations/{locationId}/optimize | Manually initiate an Optimize request for a Location ID.
[**CustomerPrototypeOverlordDeleteAppQoeConfig**](CustomerApi.md#CustomerPrototypeOverlordDeleteAppQoeConfig) | **Delete** /Customers/{id}/locations/{locationId}/config/appQoe | Resets a appQoe config. AppQoe is to monitor the Quality of Experience of these Apps in the house, which is what this PRD covers. This QoE monitoring will allow CSPs understand likely issues with applications.
[**CustomerPrototypeOverlordDeleteFlowCacheConfig**](CustomerApi.md#CustomerPrototypeOverlordDeleteFlowCacheConfig) | **Delete** /Customers/{id}/locations/{locationId}/config/flowCache | Resets a flowCache config. Enable/disable Flow Cache to help support devQA to check influence on the first stage of the investigation.
[**CustomerPrototypeOverlordDeleteSamKnowsConfig**](CustomerApi.md#CustomerPrototypeOverlordDeleteSamKnowsConfig) | **Delete** /Customers/{id}/locations/{locationId}/config/samKnows | Resets a samKnows config. SamKnows is a provider of internet performance measurement services. They offer the SamKnows Router Agent, which supports a range of QoS and QoE performance measurements. These measurements can be executed both on an ad-hoc and scheduled basis.
[**CustomerPrototypeOverlordDeleteSipAlgConfig**](CustomerApi.md#CustomerPrototypeOverlordDeleteSipAlgConfig) | **Delete** /Customers/{id}/locations/{locationId}/config/sipAlg | Resets a sipAlg config. sipAlg is an application within many routers. It inspects any VoIP traffic to prevent problems caused by firewalls and if necessary modifies the VoIP packets.
[**CustomerPrototypeOverlordDeleteStatsConfig**](CustomerApi.md#CustomerPrototypeOverlordDeleteStatsConfig) | **Delete** /Customers/{id}/locations/{locationId}/config/stats | Resets a stats config. Location Stats configuration, used to toggle which stats should be collected.
[**CustomerPrototypeOverlordUpdateAppQoeConfig**](CustomerApi.md#CustomerPrototypeOverlordUpdateAppQoeConfig) | **Patch** /Customers/{id}/locations/{locationId}/config/appQoe | Updates a appQoe config. AppQoe is to monitor the Quality of Experience of these Apps in the house, which is what this PRD covers. This QoE monitoring will allow CSPs understand likely issues with applications.
[**CustomerPrototypeOverlordUpdateFlowCacheConfig**](CustomerApi.md#CustomerPrototypeOverlordUpdateFlowCacheConfig) | **Patch** /Customers/{id}/locations/{locationId}/config/flowCache | Updates a flowCache config. Enable/disable Flow Cache to help support devQA to check influence on the first stage of the investigation.
[**CustomerPrototypeOverlordUpdateSamKnowsConfig**](CustomerApi.md#CustomerPrototypeOverlordUpdateSamKnowsConfig) | **Patch** /Customers/{id}/locations/{locationId}/config/samKnows | Updates a samKnows config. SamKnows is a provider of internet performance measurement services. They offer the SamKnows Router Agent, which supports a range of QoS and QoE performance measurements. These measurements can be executed both on an ad-hoc and scheduled basis.
[**CustomerPrototypeOverlordUpdateSipAlgConfig**](CustomerApi.md#CustomerPrototypeOverlordUpdateSipAlgConfig) | **Patch** /Customers/{id}/locations/{locationId}/config/sipAlg | Updates a sipAlg config. sipAlg is an application within many routers. It inspects any VoIP traffic to prevent problems caused by firewalls and if necessary modifies the VoIP packets.
[**CustomerPrototypeOverlordUpdateStatsConfig**](CustomerApi.md#CustomerPrototypeOverlordUpdateStatsConfig) | **Patch** /Customers/{id}/locations/{locationId}/config/stats | Updates a stats config. Location Stats configuration, used to toggle which stats should be collected.
[**CustomerPrototypePatchAccessZone**](CustomerApi.md#CustomerPrototypePatchAccessZone) | **Patch** /Customers/{id}/locations/{locationId}/wifiNetwork/accessZones/{zoneId} | Update an access zone
[**CustomerPrototypePatchAppPrioritizationLocationConfig**](CustomerApi.md#CustomerPrototypePatchAppPrioritizationLocationConfig) | **Patch** /Customers/{id}/locations/{locationId}/qos/appPrioritization | Update app prioritization config.
[**CustomerPrototypePatchAppTimeIpFlow**](CustomerApi.md#CustomerPrototypePatchAppTimeIpFlow) | **Patch** /Customers/{id}/locations/{locationId}/appTime/ipFlows | Patch IP flows config
[**CustomerPrototypePatchAttributesPatchCustomersid**](CustomerApi.md#CustomerPrototypePatchAttributesPatchCustomersid) | **Patch** /Customers/{id} | Patch attributes for a model instance and persist it into the data source.
[**CustomerPrototypePatchAttributesPutCustomersid**](CustomerApi.md#CustomerPrototypePatchAttributesPutCustomersid) | **Put** /Customers/{id} | Patch attributes for a model instance and persist it into the data source.
[**CustomerPrototypePatchCampaignCaptivePortalBranding**](CustomerApi.md#CustomerPrototypePatchCampaignCaptivePortalBranding) | **Patch** /Customers/{id}/locations/{locationId}/secondaryNetworks/captivePortals/{networkId}/campaign/branding | Patch the Captive Portal campaign branding properties for a given Location ID/NetworkId.
[**CustomerPrototypePatchCampaignCaptivePortalNetwork**](CustomerApi.md#CustomerPrototypePatchCampaignCaptivePortalNetwork) | **Patch** /Customers/{id}/locations/{locationId}/secondaryNetworks/captivePortals/{networkId}/campaign | Patch the Captive Portal campaign for a given Location ID/NetworkId.
[**CustomerPrototypePatchCaptivePortal**](CustomerApi.md#CustomerPrototypePatchCaptivePortal) | **Patch** /Customers/{id}/locations/{locationId}/secondaryNetworks/captivePortals/{networkId} | Update a Captive Portal for a given Location ID/NetworkId.
[**CustomerPrototypePatchCaptivePortalAuthorizedClients**](CustomerApi.md#CustomerPrototypePatchCaptivePortalAuthorizedClients) | **Patch** /Customers/{id}/locations/{locationId}/secondaryNetworks/captivePortals/{networkId}/authorizedClients/{mac} | Post CaptivePortal authorized clients
[**CustomerPrototypePatchCommand**](CustomerApi.md#CustomerPrototypePatchCommand) | **Patch** /Customers/{id}/locations/{locationId}/command | Returns cloud migration status for customer
[**CustomerPrototypePatchCustomDeviceType**](CustomerApi.md#CustomerPrototypePatchCustomDeviceType) | **Patch** /Customers/{id}/locations/{locationId}/devices/{mac}/customType | Update a Customer&#39;s device type configuration (user feedback).
[**CustomerPrototypePatchCustomSharedSchedule**](CustomerApi.md#CustomerPrototypePatchCustomSharedSchedule) | **Patch** /Customers/{id}/locations/{locationId}/schedules/{templateId} | Patch a custom shared schedule freeze template for a Location ID.
[**CustomerPrototypePatchDevice**](CustomerApi.md#CustomerPrototypePatchDevice) | **Patch** /Customers/{id}/locations/{locationId}/devices/{mac} | Patches a single Device to mark it favorite for a Location ID.
[**CustomerPrototypePatchDeviceAppTime**](CustomerApi.md#CustomerPrototypePatchDeviceAppTime) | **Patch** /Customers/{id}/locations/{locationId}/devices/{mac}/appTime | Update a Device&#39;s AppTime config by location ID.
[**CustomerPrototypePatchDeviceClientSteering**](CustomerApi.md#CustomerPrototypePatchDeviceClientSteering) | **Patch** /Customers/{id}/locations/{locationId}/devices/{mac}/clientSteering | Toggle auto:on/off client steering for a device.
[**CustomerPrototypePatchDeviceGroup**](CustomerApi.md#CustomerPrototypePatchDeviceGroup) | **Patch** /Customers/{id}/locations/{locationId}/networkAccess/networks/{networkId}/deviceGroups/{groupId} | Change a device group name or device members.
[**CustomerPrototypePatchDeviceOHPConfiguration**](CustomerApi.md#CustomerPrototypePatchDeviceOHPConfiguration) | **Patch** /Customers/{id}/locations/{locationId}/devices/{mac}/securityPolicy/ohp | Update the Device UUID Mapping for Out of Home Protection.
[**CustomerPrototypePatchDeviceQos**](CustomerApi.md#CustomerPrototypePatchDeviceQos) | **Patch** /Customers/{id}/locations/{locationId}/devices/{mac}/qos | Update QoS of a single device
[**CustomerPrototypePatchDeviceSecurityPolicy**](CustomerApi.md#CustomerPrototypePatchDeviceSecurityPolicy) | **Patch** /Customers/{id}/locations/{locationId}/devices/{mac}/securityPolicy | Update a Device&#39;s Security Policy for a location ID.
[**CustomerPrototypePatchDeviceSoundingState**](CustomerApi.md#CustomerPrototypePatchDeviceSoundingState) | **Patch** /Customers/{id}/locations/{locationId}/homeSecurity/devices/sounding | Patch the sounding states for the given devices
[**CustomerPrototypePatchDpp**](CustomerApi.md#CustomerPrototypePatchDpp) | **Patch** /Customers/{id}/locations/{locationId}/dpp | Patch the DPP configuration mode for a Location ID.
[**CustomerPrototypePatchEthernetLan**](CustomerApi.md#CustomerPrototypePatchEthernetLan) | **Patch** /Customers/{id}/locations/{locationId}/networkConfiguration/ethernetLan | Update the ethernetLan setting for a Location ID.
[**CustomerPrototypePatchFlowStats**](CustomerApi.md#CustomerPrototypePatchFlowStats) | **Patch** /Customers/{id}/locations/{locationId}/flowStats | Patches the flow stats configuration
[**CustomerPrototypePatchFrontHaul**](CustomerApi.md#CustomerPrototypePatchFrontHaul) | **Patch** /Customers/{id}/locations/{locationId}/secondaryNetworks/fronthauls/{networkId} | Update a Front Haul for a given Location ID/NetworkId.
[**CustomerPrototypePatchGroupOfUnassignedDevicesSecurityPolicy**](CustomerApi.md#CustomerPrototypePatchGroupOfUnassignedDevicesSecurityPolicy) | **Patch** /Customers/{id}/locations/{locationId}/groupOfUnassignedDevices/securityPolicy | Update a Location&#39;s Default Device Group Security Policy by location ID.
[**CustomerPrototypePatchHomeAwayActive**](CustomerApi.md#CustomerPrototypePatchHomeAwayActive) | **Patch** /Customers/{id}/locations/{locationId}/homeSecurity/homeAway | Enable/disable homeAway wifiMotionEvents activation for this location
[**CustomerPrototypePatchHomeSecurity**](CustomerApi.md#CustomerPrototypePatchHomeSecurity) | **Patch** /Customers/{id}/locations/{locationId}/homeSecurity | Enable/disable live motion streaming and/or motion events for this location
[**CustomerPrototypePatchHomeSecuritySensitivity**](CustomerApi.md#CustomerPrototypePatchHomeSecuritySensitivity) | **Patch** /Customers/{id}/locations/{locationId}/homeSecurity/sensitivity | Configure motion event configuration for this location
[**CustomerPrototypePatchIPv6**](CustomerApi.md#CustomerPrototypePatchIPv6) | **Patch** /Customers/{id}/locations/{locationId}/ipv6 | 
[**CustomerPrototypePatchKvConfigs**](CustomerApi.md#CustomerPrototypePatchKvConfigs) | **Patch** /Customers/{id}/locations/{locationId}/nodes/{nodeId}/kvConfigs | Retrieve all kvConfigs on a particular Node for a Location ID.
[**CustomerPrototypePatchLocation**](CustomerApi.md#CustomerPrototypePatchLocation) | **Patch** /Customers/{id}/locations/{locationId} | Update a Location&#39;s serviceId.
[**CustomerPrototypePatchLocationAppTime**](CustomerApi.md#CustomerPrototypePatchLocationAppTime) | **Patch** /Customers/{id}/locations/{locationId}/appTime | Update a Location&#39;s AppTime config by location ID.
[**CustomerPrototypePatchLocationBandSteering**](CustomerApi.md#CustomerPrototypePatchLocationBandSteering) | **Patch** /Customers/{id}/locations/{locationId}/bandSteering | Set mode for band steering
[**CustomerPrototypePatchLocationFreezeAutoExpire**](CustomerApi.md#CustomerPrototypePatchLocationFreezeAutoExpire) | **Patch** /Customers/{id}/locations/{locationId}/freeze/autoExpire | Put all devices except some to be frozen for a Location ID.
[**CustomerPrototypePatchLocationManager**](CustomerApi.md#CustomerPrototypePatchLocationManager) | **Patch** /Customers/{id}/locations/{locationId}/managers/{managerId} | Update type of access of manager on location.
[**CustomerPrototypePatchLocationQoeLiveMode**](CustomerApi.md#CustomerPrototypePatchLocationQoeLiveMode) | **Patch** /Customers/{id}/locations/{locationId}/qoe/liveMode | Update the location qoe liveMode by api call and Kafka message
[**CustomerPrototypePatchLocationSecurityPolicy**](CustomerApi.md#CustomerPrototypePatchLocationSecurityPolicy) | **Patch** /Customers/{id}/locations/{locationId}/securityPolicy | Update a Location&#39;s Security Policy by location ID.
[**CustomerPrototypePatchMulticast**](CustomerApi.md#CustomerPrototypePatchMulticast) | **Patch** /Customers/{id}/locations/{locationId}/networkConfiguration/multicast | Update the multicast settings for a Location ID.
[**CustomerPrototypePatchNetworkAccessNetwork**](CustomerApi.md#CustomerPrototypePatchNetworkAccessNetwork) | **Patch** /Customers/{id}/locations/{locationId}/networkAccess/networks/{networkId} | Enable or disable purgatory in the network
[**CustomerPrototypePatchOptimizations**](CustomerApi.md#CustomerPrototypePatchOptimizations) | **Patch** /Customers/{id}/locations/{locationId}/optimizations | Enable/disable optimizations for a Location ID.
[**CustomerPrototypePatchPerson**](CustomerApi.md#CustomerPrototypePatchPerson) | **Patch** /Customers/{id}/locations/{locationId}/persons/{personId} | Update a Person for a location ID.
[**CustomerPrototypePatchPersonAppTime**](CustomerApi.md#CustomerPrototypePatchPersonAppTime) | **Patch** /Customers/{id}/locations/{locationId}/persons/{personId}/appTime | Update a Person&#39;s AppTime config by location ID.
[**CustomerPrototypePatchPersonProfile**](CustomerApi.md#CustomerPrototypePatchPersonProfile) | **Patch** /Customers/{id}/locations/{locationId}/persons/{personId}/profile | Update a Person&#39;s Profile for a location ID.
[**CustomerPrototypePatchPersonSecurityPolicy**](CustomerApi.md#CustomerPrototypePatchPersonSecurityPolicy) | **Patch** /Customers/{id}/locations/{locationId}/persons/{personId}/securityPolicy | Update a Person&#39;s Security Policy for a location ID.
[**CustomerPrototypePatchRemoteConnectionsConfig**](CustomerApi.md#CustomerPrototypePatchRemoteConnectionsConfig) | **Patch** /Customers/{id}/locations/{locationId}/securityPolicy/remoteConnections | Patch a Remote Connections Config for the given Location ID.
[**CustomerPrototypePatchRoom**](CustomerApi.md#CustomerPrototypePatchRoom) | **Patch** /Customers/{id}/locations/{locationId}/rooms/{roomId} | Patch a Room for a Location ID/Room ID.
[**CustomerPrototypePatchSecurityConfiguration**](CustomerApi.md#CustomerPrototypePatchSecurityConfiguration) | **Patch** /Customers/{id}/locations/{locationId}/securityConfiguration | Patch Security Configurations for location (preferredIntelligence, etc)
[**CustomerPrototypePatchServiceLevel**](CustomerApi.md#CustomerPrototypePatchServiceLevel) | **Patch** /Customers/{id}/locations/{locationId}/serviceLevel | Set the service level for this location
[**CustomerPrototypePatchSubscription**](CustomerApi.md#CustomerPrototypePatchSubscription) | **Patch** /Customers/{id}/locations/{locationId}/subscription | Patch Subscription details for this location
[**CustomerPrototypePatchWifiMotion**](CustomerApi.md#CustomerPrototypePatchWifiMotion) | **Patch** /Customers/{id}/locations/{locationId}/wifiMotion | Enable/disable WifiMotion feature for this location
[**CustomerPrototypePatchWifiNetwork**](CustomerApi.md#CustomerPrototypePatchWifiNetwork) | **Patch** /Customers/{id}/locations/{locationId}/wifiNetwork | Update the SSID of the WifiNetwork
[**CustomerPrototypePostCampaignPreviewCaptivePortalNetwork**](CustomerApi.md#CustomerPrototypePostCampaignPreviewCaptivePortalNetwork) | **Post** /Customers/{id}/locations/{locationId}/secondaryNetworks/captivePortals/{networkId}/campaign/preview | POST Captive Portal campaign preview for a given Location ID/NetworkId.
[**CustomerPrototypePostCaptivePortal**](CustomerApi.md#CustomerPrototypePostCaptivePortal) | **Post** /Customers/{id}/locations/{locationId}/secondaryNetworks/captivePortals | Create a Captive Portal Network for a Location ID.
[**CustomerPrototypePostCaptivePortalAuthorizedClients**](CustomerApi.md#CustomerPrototypePostCaptivePortalAuthorizedClients) | **Post** /Customers/{id}/locations/{locationId}/secondaryNetworks/captivePortals/{networkId}/authorizedClients | Post CaptivePortal authorized clients
[**CustomerPrototypePostCaptivePortalCampaignAsset**](CustomerApi.md#CustomerPrototypePostCaptivePortalCampaignAsset) | **Post** /Customers/{id}/locations/{locationId}/secondaryNetworks/captivePortals/uploadCampaignAsset | Upload campaign asset for given location.
[**CustomerPrototypePostCaptivePortalEnableGuestEmailCollection**](CustomerApi.md#CustomerPrototypePostCaptivePortalEnableGuestEmailCollection) | **Post** /Customers/{id}/locations/{locationId}/secondaryNetworks/captivePortals/{networkId}/enableGuestEmailCollection | Patch the Captive Portal Network to be compliant for guest email collection.
[**CustomerPrototypePostCaptivePortalNetworkUsageStats**](CustomerApi.md#CustomerPrototypePostCaptivePortalNetworkUsageStats) | **Post** /Customers/{id}/locations/{locationId}/secondaryNetworks/captivePortals/{networkId}/networkUsage | Fetch the Captive Portal Network Usage stats for the given network.
[**CustomerPrototypePostCustomSharedSchedule**](CustomerApi.md#CustomerPrototypePostCustomSharedSchedule) | **Post** /Customers/{id}/locations/{locationId}/schedules | Create \&quot;custom shared\&quot; schedules that shared by all persons and devices in a location.
[**CustomerPrototypePostDeviceFreeze**](CustomerApi.md#CustomerPrototypePostDeviceFreeze) | **Post** /Customers/{id}/locations/{locationId}/devices/{mac}/freeze/{freezeTemplateId} | Post a shared schedule uuid freeze for a device for a Location ID.
[**CustomerPrototypePostDeviceGroup**](CustomerApi.md#CustomerPrototypePostDeviceGroup) | **Post** /Customers/{id}/locations/{locationId}/networkAccess/networks/{networkId}/deviceGroups | Create a named device group within a network and optionally specify member devices.
[**CustomerPrototypePostDeviceQos**](CustomerApi.md#CustomerPrototypePostDeviceQos) | **Post** /Customers/{id}/locations/{locationId}/devices/{mac}/qos | Set QoS of a single device
[**CustomerPrototypePostDeviceSecurityPolicyAnomalyExperience**](CustomerApi.md#CustomerPrototypePostDeviceSecurityPolicyAnomalyExperience) | **Post** /Customers/{id}/locations/{locationId}/devices/{mac}/securityPolicy/anomaly/experience | Initiate an Anomaly Experience (demo) for a Device on a location.
[**CustomerPrototypePostDeviceSecurityPolicyAnomalyWhitelist**](CustomerApi.md#CustomerPrototypePostDeviceSecurityPolicyAnomalyWhitelist) | **Post** /Customers/{id}/locations/{locationId}/devices/{mac}/securityPolicy/anomaly/websites/whitelist | Approve a previously blacklisted anomalous dns for a Device on a location.
[**CustomerPrototypePostDeviceSecurityPolicyWebsitesBlacklist**](CustomerApi.md#CustomerPrototypePostDeviceSecurityPolicyWebsitesBlacklist) | **Post** /Customers/{id}/locations/{locationId}/devices/{mac}/securityPolicy/websites/blacklist | Update a Device&#39;s Security Policy for a location ID to include a blacklisted DNS entry.
[**CustomerPrototypePostDeviceSecurityPolicyWebsitesWhitelist**](CustomerApi.md#CustomerPrototypePostDeviceSecurityPolicyWebsitesWhitelist) | **Post** /Customers/{id}/locations/{locationId}/devices/{mac}/securityPolicy/websites/whitelist | Update a Device&#39;s Security Policy for a location ID to include a whitelisted DNS entry.
[**CustomerPrototypePostDeviceToAccessZone**](CustomerApi.md#CustomerPrototypePostDeviceToAccessZone) | **Post** /Customers/{id}/locations/{locationId}/wifiNetwork/accessZones/{zoneId}/accessibleDevices/{mac} | Add a device mac to a WiFi Access Zone
[**CustomerPrototypePostFrontHaul**](CustomerApi.md#CustomerPrototypePostFrontHaul) | **Post** /Customers/{id}/locations/{locationId}/secondaryNetworks/fronthauls | Create a Front Haul Network for a Location ID.
[**CustomerPrototypePostFrontHaulsDpp**](CustomerApi.md#CustomerPrototypePostFrontHaulsDpp) | **Post** /Customers/{id}/locations/{locationId}/secondaryNetworks/fronthauls/{networkId}/dpp | Create the DPP setting for a Fronthaul Network.
[**CustomerPrototypePostFrontHaulsDppBootstrap**](CustomerApi.md#CustomerPrototypePostFrontHaulsDppBootstrap) | **Post** /Customers/{id}/locations/{locationId}/secondaryNetworks/fronthauls/{networkId}/dpp/bootstrapUris | Create a bootstrap for DPP setting for a Fronthaul Network.
[**CustomerPrototypePostFrontHaulsDppEnrollment**](CustomerApi.md#CustomerPrototypePostFrontHaulsDppEnrollment) | **Post** /Customers/{id}/locations/{locationId}/secondaryNetworks/fronthauls/{networkId}/dpp/enrollments | Create an enrollment for DPP setting for a fronthaul secondary network.
[**CustomerPrototypePostGroupOfUnassignedDevicesFreezeTemplateId**](CustomerApi.md#CustomerPrototypePostGroupOfUnassignedDevicesFreezeTemplateId) | **Post** /Customers/{id}/locations/{locationId}/groupOfUnassignedDevices/freeze/{freezeTemplateId} | POST GroupOfUnassignedDevices to be frozen for a Location ID.
[**CustomerPrototypePostGroupOfUnassignedDevicesSecurityPolicyWebsitesBlacklist**](CustomerApi.md#CustomerPrototypePostGroupOfUnassignedDevicesSecurityPolicyWebsitesBlacklist) | **Post** /Customers/{id}/locations/{locationId}/groupOfUnassignedDevices/securityPolicy/websites/blacklist | Update a Location&#39;s Default Device Group Security Policy for a location ID to include a blacklisted DNS entry.
[**CustomerPrototypePostGroupOfUnassignedDevicesSecurityPolicyWebsitesWhitelist**](CustomerApi.md#CustomerPrototypePostGroupOfUnassignedDevicesSecurityPolicyWebsitesWhitelist) | **Post** /Customers/{id}/locations/{locationId}/groupOfUnassignedDevices/securityPolicy/websites/whitelist | Update a Location&#39;s Default Device Group Security Policy for a location ID to include a whitelisted DNS entry.
[**CustomerPrototypePostKvConfigs**](CustomerApi.md#CustomerPrototypePostKvConfigs) | **Post** /Customers/{id}/locations/{locationId}/nodes/{nodeId}/kvConfigs | Retrieve all kvConfigs on a particular Node for a Location ID.
[**CustomerPrototypePostLocationSecurityPolicyOHPDeviceSetup**](CustomerApi.md#CustomerPrototypePostLocationSecurityPolicyOHPDeviceSetup) | **Post** /Customers/{id}/locations/{locationId}/securityPolicy/ohp/deviceSetup | Setup a Mobile Device for Security Out of Home Protection (returns a Deeplink for use with Mobolize).
[**CustomerPrototypePostLocationSecurityPolicyWebsitesBlacklist**](CustomerApi.md#CustomerPrototypePostLocationSecurityPolicyWebsitesBlacklist) | **Post** /Customers/{id}/locations/{locationId}/securityPolicy/websites/blacklist | Update a Location&#39;s Security Policy for a location ID to include a blacklisted DNS entry.
[**CustomerPrototypePostLocationSecurityPolicyWebsitesWhitelist**](CustomerApi.md#CustomerPrototypePostLocationSecurityPolicyWebsitesWhitelist) | **Post** /Customers/{id}/locations/{locationId}/securityPolicy/websites/whitelist | Update a Location&#39;s Security Policy for a location ID to include a whitelisted DNS entry.
[**CustomerPrototypePostManager**](CustomerApi.md#CustomerPrototypePostManager) | **Post** /Customers/{id}/locations/{locationId}/managers | Assign a manager to your location 
[**CustomerPrototypePostOnboardingCheckpoint**](CustomerApi.md#CustomerPrototypePostOnboardingCheckpoint) | **Post** /Customers/{id}/locations/{locationId}/onboardingCheckpoint | Record the new Onboarding Checkpoint for the Location ID.
[**CustomerPrototypePostPersonFreeze**](CustomerApi.md#CustomerPrototypePostPersonFreeze) | **Post** /Customers/{id}/locations/{locationId}/persons/{personId}/freeze/{freezeTemplateId} | Post a shared schedule uuid freeze for a person for a Location ID.
[**CustomerPrototypePostPersonSecurityPolicyWebsitesBlacklist**](CustomerApi.md#CustomerPrototypePostPersonSecurityPolicyWebsitesBlacklist) | **Post** /Customers/{id}/locations/{locationId}/persons/{personId}/securityPolicy/websites/blacklist | Update a Person&#39;s Security Policy for a location ID to include a blacklisted DNS entry.
[**CustomerPrototypePostPersonSecurityPolicyWebsitesWhitelist**](CustomerApi.md#CustomerPrototypePostPersonSecurityPolicyWebsitesWhitelist) | **Post** /Customers/{id}/locations/{locationId}/persons/{personId}/securityPolicy/websites/whitelist | Update a Person&#39;s Security Policy for a location ID to include a whitelisted DNS entry.
[**CustomerPrototypePostPersons**](CustomerApi.md#CustomerPrototypePostPersons) | **Post** /Customers/{id}/locations/{locationId}/persons | Create  a Person for a Location ID.
[**CustomerPrototypePostPortForward**](CustomerApi.md#CustomerPrototypePostPortForward) | **Post** /Customers/{id}/locations/{locationId}/networkConfiguration/dhcpReservations/{mac}/portForward | Record a new Port Forwarding entry for an existing DHCP IP reservation tied to a MAC address at a Location ID.
[**CustomerPrototypePostRemoteConnectionsAllow**](CustomerApi.md#CustomerPrototypePostRemoteConnectionsAllow) | **Post** /Customers/{id}/locations/{locationId}/devices/{mac}/securityPolicy/remoteConnections/allow | Post a Remote Connection Allow IpAddress/ttl for the given device and Location ID.
[**CustomerPrototypePostRemoteConnectionsAllowAll**](CustomerApi.md#CustomerPrototypePostRemoteConnectionsAllowAll) | **Post** /Customers/{id}/locations/{locationId}/devices/{mac}/securityPolicy/remoteConnections/allowAll | Post a Remote Connection Allow All/ttl for the given device and Location ID.
[**CustomerPrototypePostRooms**](CustomerApi.md#CustomerPrototypePostRooms) | **Post** /Customers/{id}/locations/{locationId}/rooms | Create a Room for a Location ID.
[**CustomerPrototypePostRunMobileIspSpeedTest**](CustomerApi.md#CustomerPrototypePostRunMobileIspSpeedTest) | **Post** /Customers/{id}/locations/{locationId}/ispSpeedTest | Run ISP speed test for GW node on mobile request.
[**CustomerPrototypePostSpeedTest**](CustomerApi.md#CustomerPrototypePostSpeedTest) | **Post** /Customers/{id}/locations/{locationId}/nodes/{nodeId}/speedTest | Run speed test for a node.
[**CustomerPrototypePostWhitelistApprovalRequest**](CustomerApi.md#CustomerPrototypePostWhitelistApprovalRequest) | **Post** /Customers/{id}/locations/{locationId}/securityPolicy/websites/whitelist/approvalRequests | Post a request for a whitelist exception to be added to your person profile.
[**CustomerPrototypePostWifiAccessZone**](CustomerApi.md#CustomerPrototypePostWifiAccessZone) | **Post** /Customers/{id}/locations/{locationId}/wifiNetwork/accessZones | Create a new WiFi Access Zone
[**CustomerPrototypePostWifiKey**](CustomerApi.md#CustomerPrototypePostWifiKey) | **Post** /Customers/{id}/locations/{locationId}/wifiNetwork/accessZones/{accessZone}/keys | Create a new WiFi Password
[**CustomerPrototypePostWifiNetwork**](CustomerApi.md#CustomerPrototypePostWifiNetwork) | **Post** /Customers/{id}/locations/{locationId}/wifiNetwork | Update a WiFi SSID and/or PSK for a Location ID.
[**CustomerPrototypePostWifiNetworkDpp**](CustomerApi.md#CustomerPrototypePostWifiNetworkDpp) | **Post** /Customers/{id}/locations/{locationId}/wifiNetwork/dpp | Create the DPP setting for a Location ID.
[**CustomerPrototypePostWifiNetworkDppBootstrap**](CustomerApi.md#CustomerPrototypePostWifiNetworkDppBootstrap) | **Post** /Customers/{id}/locations/{locationId}/wifiNetwork/dpp/bootstrapUris | Create a bootstrap for DPP setting for a wifi network.
[**CustomerPrototypePostWifiNetworkDppEnrollment**](CustomerApi.md#CustomerPrototypePostWifiNetworkDppEnrollment) | **Post** /Customers/{id}/locations/{locationId}/wifiNetwork/dpp/enrollments | Create an enrollment for DPP setting for a wifi network.
[**CustomerPrototypePublishSlowChangingDimensionConfigs**](CustomerApi.md#CustomerPrototypePublishSlowChangingDimensionConfigs) | **Post** /Customers/{id}/publishSlowChangingDimensionConfigs | Publish all slow changing dimension Kafka messages
[**CustomerPrototypePutAuthorizationsPutCustomersidLocationslocationIdAuthorizations**](CustomerApi.md#CustomerPrototypePutAuthorizationsPutCustomersidLocationslocationIdAuthorizations) | **Put** /Customers/{id}/locations/{locationId}/authorizations | Configure number of authorized leaf pods for a Location ID.
[**CustomerPrototypePutAuthorizationsPutCustomersidLocationslocationIdNodeAuthorizations**](CustomerApi.md#CustomerPrototypePutAuthorizationsPutCustomersidLocationslocationIdNodeAuthorizations) | **Put** /Customers/{id}/locations/{locationId}/nodeAuthorizations | Configure number of authorized leaf pods grouped by model id for a Location ID.
[**CustomerPrototypePutBackhaul**](CustomerApi.md#CustomerPrototypePutBackhaul) | **Put** /Customers/{id}/locations/{locationId}/backhaul | Toggle secure backhaul for a Location ID.
[**CustomerPrototypePutBandSteering**](CustomerApi.md#CustomerPrototypePutBandSteering) | **Put** /Customers/{id}/locations/{locationId}/bandSteering | Enable/disable band steering for a Location ID (deprecated)
[**CustomerPrototypePutBleMode**](CustomerApi.md#CustomerPrototypePutBleMode) | **Put** /Customers/{id}/locations/{locationId}/bleMode | Enable or Disable BLE beaconing for all Pods at a location for Pod location services (e.g. for Pods Naming).
[**CustomerPrototypePutBleModeForNode**](CustomerApi.md#CustomerPrototypePutBleModeForNode) | **Put** /Customers/{id}/locations/{locationId}/nodes/{nodeId}/bleMode | Enable or Disable BLE beaconing for the specific Pod at a location.
[**CustomerPrototypePutCaptivePortalSendDetails**](CustomerApi.md#CustomerPrototypePutCaptivePortalSendDetails) | **Put** /Customers/{id}/locations/{locationId}/secondaryNetworks/captivePortals/{networkId}/sendGuestDetails | Send Captive Portal Guest details to requesters email for a given Location ID/NetworkId.
[**CustomerPrototypePutCommand**](CustomerApi.md#CustomerPrototypePutCommand) | **Put** /Customers/{id}/locations/{locationId}/command | Returns cloud migration status for customer
[**CustomerPrototypePutControlMode**](CustomerApi.md#CustomerPrototypePutControlMode) | **Put** /Customers/{id}/locations/{locationId}/controlMode | Set control mode for a Location ID.
[**CustomerPrototypePutCouncilmanResync**](CustomerApi.md#CustomerPrototypePutCouncilmanResync) | **Put** /Customers/{id}/locations/{locationId}/councilman/resync | Push Security Configurations to Councilman.
[**CustomerPrototypePutDeviceFreeze**](CustomerApi.md#CustomerPrototypePutDeviceFreeze) | **Put** /Customers/{id}/locations/{locationId}/devices/{mac}/freeze/{freezeTemplateId} | Put a device to be frozen for a Location ID.
[**CustomerPrototypePutDeviceFreezeAutoExpire**](CustomerApi.md#CustomerPrototypePutDeviceFreezeAutoExpire) | **Put** /Customers/{id}/locations/{locationId}/devices/{mac}/freeze/autoExpire | Put a device to be frozen for a Location ID.
[**CustomerPrototypePutDeviceFreezeForever**](CustomerApi.md#CustomerPrototypePutDeviceFreezeForever) | **Put** /Customers/{id}/locations/{locationId}/devices/{mac}/freeze/forever | Put a device forever freeze for a Location ID.
[**CustomerPrototypePutDeviceFreezeResidentialGwManaged**](CustomerApi.md#CustomerPrototypePutDeviceFreezeResidentialGwManaged) | **Put** /Customers/{id}/locations/{locationId}/devices/{mac}/freeze/residentialGwManaged | Put a device residentialGwManaged freeze for a Location ID.
[**CustomerPrototypePutDeviceFreezeSuspend**](CustomerApi.md#CustomerPrototypePutDeviceFreezeSuspend) | **Put** /Customers/{id}/locations/{locationId}/devices/{mac}/freeze/suspend | Put a device suspend for a Location ID.
[**CustomerPrototypePutDeviceNickname**](CustomerApi.md#CustomerPrototypePutDeviceNickname) | **Put** /Customers/{id}/devices/{mac} | Nickname a Customer&#39;s device for all locations.
[**CustomerPrototypePutDevicesFreezeCommand**](CustomerApi.md#CustomerPrototypePutDevicesFreezeCommand) | **Put** /Customers/{id}/locations/{locationId}/command/devices/freeze | Sets autoExpire or freezeTemplateId suspend for specified mac addresses
[**CustomerPrototypePutDevicesVisibleToGuests**](CustomerApi.md#CustomerPrototypePutDevicesVisibleToGuests) | **Put** /Customers/{id}/locations/{locationId}/wifiNetwork/accessZones/home/devicesVisibleToGuests | DEPRECATED: Update home devices visible to guests.
[**CustomerPrototypePutDhcp**](CustomerApi.md#CustomerPrototypePutDhcp) | **Put** /Customers/{id}/locations/{locationId}/networkConfiguration/dhcp | Record or update a new DHCP subnet/subnetMask for a Location ID.
[**CustomerPrototypePutDhcpReservation**](CustomerApi.md#CustomerPrototypePutDhcpReservation) | **Put** /Customers/{id}/locations/{locationId}/networkConfiguration/dhcpReservations/{mac} | Record or update a new DHCP IP Reservation for a particular MAC address at a Location ID.
[**CustomerPrototypePutDnsServers**](CustomerApi.md#CustomerPrototypePutDnsServers) | **Put** /Customers/{id}/locations/{locationId}/networkConfiguration/dnsServers | Update the DNS IPv4 server addresses for a Location ID.
[**CustomerPrototypePutDppEnrollments**](CustomerApi.md#CustomerPrototypePutDppEnrollments) | **Put** /Customers/{id}/locations/{locationId}/dpp/enrollments | Create and persist a list of DPP enrollments
[**CustomerPrototypePutEthernetLan**](CustomerApi.md#CustomerPrototypePutEthernetLan) | **Put** /Customers/{id}/locations/{locationId}/nodes/{nodeId}/ethernetLan | Updates location nodes with ethernetLan modes
[**CustomerPrototypePutFirmwareUpgradeRequest**](CustomerApi.md#CustomerPrototypePutFirmwareUpgradeRequest) | **Put** /Customers/{id}/locations/{locationId}/firmware | Request Firmware Upgrade for a Location ID
[**CustomerPrototypePutFrontlineStorage**](CustomerApi.md#CustomerPrototypePutFrontlineStorage) | **Put** /Customers/{id}/locations/{locationId}/frontline/storage | Create or Update the frontline storage for a Location ID
[**CustomerPrototypePutGroupOfUnassignedDevicesFreezeAutoExpire**](CustomerApi.md#CustomerPrototypePutGroupOfUnassignedDevicesFreezeAutoExpire) | **Put** /Customers/{id}/locations/{locationId}/groupOfUnassignedDevices/freeze/autoExpire | Put GroupOfUnassignedDevices autoExpire freeze for a Location ID.
[**CustomerPrototypePutGroupOfUnassignedDevicesFreezeForever**](CustomerApi.md#CustomerPrototypePutGroupOfUnassignedDevicesFreezeForever) | **Put** /Customers/{id}/locations/{locationId}/groupOfUnassignedDevices/freeze/forever | PUT GroupOfUnassignedDevices forever for a Location ID.
[**CustomerPrototypePutGroupOfUnassignedDevicesFreezeSuspend**](CustomerApi.md#CustomerPrototypePutGroupOfUnassignedDevicesFreezeSuspend) | **Put** /Customers/{id}/locations/{locationId}/groupOfUnassignedDevices/freeze/suspend | PUT GroupOfUnassignedDevices suspend for a Location ID.
[**CustomerPrototypePutIosDeviceToken**](CustomerApi.md#CustomerPrototypePutIosDeviceToken) | **Put** /Customers/{id}/iosDeviceToken/{deviceToken} | Inserts the iOS device token for the Customer ID, which may be used for notification services.
[**CustomerPrototypePutIspSpeedTestConfiguration**](CustomerApi.md#CustomerPrototypePutIspSpeedTestConfiguration) | **Put** /Customers/{id}/locations/{locationId}/ispSpeedTestConfiguration | Enable|Disable ispSpeedTestConfiguration to schedule speed tests.
[**CustomerPrototypePutLedMode**](CustomerApi.md#CustomerPrototypePutLedMode) | **Put** /Customers/{id}/locations/{locationId}/nodes/{nodeId}/ledMode | Update the LED mode on a particular Node for a Location ID.
[**CustomerPrototypePutLocale**](CustomerApi.md#CustomerPrototypePutLocale) | **Put** /Customers/{id}/locations/{locationId}/locale | Configure locale values for a Location ID.
[**CustomerPrototypePutLocationSecurityPolicyOHPDeviceUuidMapping**](CustomerApi.md#CustomerPrototypePutLocationSecurityPolicyOHPDeviceUuidMapping) | **Put** /Customers/{id}/locations/{locationId}/securityPolicy/ohp/deviceUuid | Update the Device UUID Mapping for Out of Home Protection.
[**CustomerPrototypePutLocationSecurityPolicyOHPProtectionState**](CustomerApi.md#CustomerPrototypePutLocationSecurityPolicyOHPProtectionState) | **Put** /Customers/{id}/locations/{locationId}/securityPolicy/ohp/protectionState | Update the Device Protection State for Out of Home Protection.
[**CustomerPrototypePutLocationWanConfiguration**](CustomerApi.md#CustomerPrototypePutLocationWanConfiguration) | **Put** /Customers/{id}/locations/{locationId}/wanConfiguration | Persist WAN Configuration for a Location ID.
[**CustomerPrototypePutLocationWanSettings**](CustomerApi.md#CustomerPrototypePutLocationWanSettings) | **Put** /Customers/{id}/locations/{locationId}/wanSettings | DEPRECATED: Persist WAN Settings for a Location ID.
[**CustomerPrototypePutMonitorMode**](CustomerApi.md#CustomerPrototypePutMonitorMode) | **Put** /Customers/{id}/locations/{locationId}/monitorMode | Enable/disable monitor mode for a Location ID.
[**CustomerPrototypePutNetworkMode**](CustomerApi.md#CustomerPrototypePutNetworkMode) | **Put** /Customers/{id}/locations/{locationId}/networkMode | Update the Network Mode for a Location ID.
[**CustomerPrototypePutOptimizations**](CustomerApi.md#CustomerPrototypePutOptimizations) | **Put** /Customers/{id}/locations/{locationId}/optimizations | Enable/disable optimizations for a Location ID.
[**CustomerPrototypePutOverlordResync**](CustomerApi.md#CustomerPrototypePutOverlordResync) | **Put** /Customers/{id}/locations/{locationId}/overlord/resync | Push Secondary Network Configurations to Overlord.
[**CustomerPrototypePutPersistConfigurationOnGateway**](CustomerApi.md#CustomerPrototypePutPersistConfigurationOnGateway) | **Put** /Customers/{id}/locations/{locationId}/networkConfiguration/persistConfigurationOnGateway | Update settings for persistConfigurationOnGateway.
[**CustomerPrototypePutPersonFreeze**](CustomerApi.md#CustomerPrototypePutPersonFreeze) | **Put** /Customers/{id}/locations/{locationId}/persons/{personId}/freeze/{freezeTemplateId} | Put a person to be frozen for a Location ID.
[**CustomerPrototypePutPersonFreezeAutoExpire**](CustomerApi.md#CustomerPrototypePutPersonFreezeAutoExpire) | **Put** /Customers/{id}/locations/{locationId}/persons/{personId}/freeze/autoExpire | Put all devices from a person to be frozen for a Location ID.
[**CustomerPrototypePutPersonFreezeForever**](CustomerApi.md#CustomerPrototypePutPersonFreezeForever) | **Put** /Customers/{id}/locations/{locationId}/persons/{personId}/freeze/forever | Put a person forever freeze for a Location ID.
[**CustomerPrototypePutPersonFreezeSuspend**](CustomerApi.md#CustomerPrototypePutPersonFreezeSuspend) | **Put** /Customers/{id}/locations/{locationId}/persons/{personId}/freeze/suspend | Put a person suspend for a Location ID.
[**CustomerPrototypePutPortForward**](CustomerApi.md#CustomerPrototypePutPortForward) | **Put** /Customers/{id}/locations/{locationId}/networkConfiguration/dhcpReservations/{mac}/portForward/{externalPort} | Update an existing Port Forwarding entry for an existing DHCP IP reservation tied to a MAC address at a Location ID.
[**CustomerPrototypePutSniffing**](CustomerApi.md#CustomerPrototypePutSniffing) | **Put** /Customers/{id}/locations/{locationId}/sniffing | Updates location sniffing toggle modes
[**CustomerPrototypePutSnoozeOnDeviceAlert**](CustomerApi.md#CustomerPrototypePutSnoozeOnDeviceAlert) | **Put** /Customers/{id}/locations/{locationId}/devices/{mac}/alerts/{type} | Snooze an alert on a device.
[**CustomerPrototypePutSnoozeOnNodeAlert**](CustomerApi.md#CustomerPrototypePutSnoozeOnNodeAlert) | **Put** /Customers/{id}/locations/{locationId}/nodes/{nodeId}/alerts/{type} | Snooze an alert on a node.
[**CustomerPrototypePutSubscription**](CustomerApi.md#CustomerPrototypePutSubscription) | **Put** /Customers/{id}/locations/{locationId}/subscription | Put Subscription details for this location
[**CustomerPrototypePutUpnp**](CustomerApi.md#CustomerPrototypePutUpnp) | **Put** /Customers/{id}/locations/{locationId}/networkConfiguration/upnp | Update the UPnP setting for a Location ID.
[**CustomerPrototypePutWifiKey**](CustomerApi.md#CustomerPrototypePutWifiKey) | **Put** /Customers/{id}/locations/{locationId}/wifiNetwork/accessZones/{accessZone}/keys/{keyId} | Update a WiFi Password
[**CustomerPrototypePutWifiNetwork**](CustomerApi.md#CustomerPrototypePutWifiNetwork) | **Put** /Customers/{id}/locations/{locationId}/wifiNetwork | Update a WiFi SSID and/or PSK for a Location ID.
[**CustomerPrototypeRebootLocation**](CustomerApi.md#CustomerPrototypeRebootLocation) | **Put** /Customers/{id}/locations/{locationId}/reboot | Reboots a particular on-line Node for a Location ID.
[**CustomerPrototypeRebootNode**](CustomerApi.md#CustomerPrototypeRebootNode) | **Put** /Customers/{id}/locations/{locationId}/nodes/{nodeId}/reboot | Reboots a single on-line Node for a Location ID.
[**CustomerPrototypeRejectWhitelistRequest**](CustomerApi.md#CustomerPrototypeRejectWhitelistRequest) | **Delete** /Customers/{id}/locations/{locationId}/securityPolicy/websites/whitelist/approvalRequests/{requestId} | Reject an approval request for a website whitelist
[**CustomerPrototypeRemoveDeviceByMac**](CustomerApi.md#CustomerPrototypeRemoveDeviceByMac) | **Delete** /Customers/{id}/locations/{locationId}/devices/{mac} | Removes a device for a customer&#39;s location id, wiping config and setting a hidden flag.
[**CustomerPrototypeRemoveDeviceVisibleToGuests**](CustomerApi.md#CustomerPrototypeRemoveDeviceVisibleToGuests) | **Delete** /Customers/{id}/locations/{locationId}/wifiNetwork/accessZones/home/devicesVisibleToGuests/{mac} | DEPRECATED: Update home devices visible to guests.
[**CustomerPrototypeRenameNode**](CustomerApi.md#CustomerPrototypeRenameNode) | **Put** /Customers/{id}/locations/{locationId}/nodes/{nodeId} | Rename a particular Node for a Location ID with the option to disable the blinking LED.
[**CustomerPrototypeResendManagerInvite**](CustomerApi.md#CustomerPrototypeResendManagerInvite) | **Post** /Customers/{id}/locations/{locationId}/managers/{managerId}/resendInvite | Resend invite to a manager that has status \&quot;pending\&quot;.
[**CustomerPrototypeResetTos**](CustomerApi.md#CustomerPrototypeResetTos) | **Post** /Customers/{id}/locations/{locationId}/devices/{mac}/tos/reset | Resets the back-off and thresholds for the given client.
[**CustomerPrototypeSetForcedSteer**](CustomerApi.md#CustomerPrototypeSetForcedSteer) | **Put** /Customers/{id}/locations/{locationId}/devices/{mac}/forcedSteer | Force a device to use the 2.4Ghz band with auto expire.
[**CustomerPrototypeSetPrimarySecondaryNetworks**](CustomerApi.md#CustomerPrototypeSetPrimarySecondaryNetworks) | **Put** /Customers/{id}/locations/{locationId}/primarySecondaryNetworks | Set networks at wpa3 transition flow
[**CustomerPrototypeShareDevice**](CustomerApi.md#CustomerPrototypeShareDevice) | **Put** /Customers/{id}/locations/{locationId}/networkAccess/networks/{networkId}/devices/{mac}/groupShares | Share access to individual device. 
[**CustomerPrototypeShareDeviceGroup**](CustomerApi.md#CustomerPrototypeShareDeviceGroup) | **Put** /Customers/{id}/locations/{locationId}/networkAccess/networks/{networkId}/deviceGroups/{groupId}/groupShares | Share access for a group or employee.
[**CustomerPrototypeStartWps**](CustomerApi.md#CustomerPrototypeStartWps) | **Post** /Customers/{id}/locations/{locationId}/startWps | Start a WPS session
[**CustomerPrototypeStitchDevice**](CustomerApi.md#CustomerPrototypeStitchDevice) | **Post** /Customers/{id}/locations/{locationId}/devices/stitch | Delete prioritization of a single device
[**CustomerPrototypeUnapproveDevice**](CustomerApi.md#CustomerPrototypeUnapproveDevice) | **Delete** /Customers/{id}/locations/{locationId}/networkAccess/networks/{networkId}/approved/{mac} | Unapprove approved devices in the network
[**CustomerPrototypeUnblockDevice**](CustomerApi.md#CustomerPrototypeUnblockDevice) | **Delete** /Customers/{id}/locations/{locationId}/networkAccess/blocked/{mac} | Unblock blocked devices
[**CustomerPrototypeUnclaimAllNodes**](CustomerApi.md#CustomerPrototypeUnclaimAllNodes) | **Delete** /Customers/{id}/locations/{locationId}/nodes | Unclaim all Nodes from a Location ID with the option of preserving the original Package ID.
[**CustomerPrototypeUnclaimNode**](CustomerApi.md#CustomerPrototypeUnclaimNode) | **Delete** /Customers/{id}/locations/{locationId}/nodes/{nodeId} | Unclaim a particular Node from a Location ID with the option of preserving the original Package ID.
[**CustomerPrototypeUpdateLocationName**](CustomerApi.md#CustomerPrototypeUpdateLocationName) | **Put** /Customers/{id}/locations/{locationId} | Update the location name.
[**CustomerPrototypeUpdateMigration**](CustomerApi.md#CustomerPrototypeUpdateMigration) | **Put** /Customers/{id}/_migration | Update _migration of this model.
[**CustomerPrototypeUpdateTermsAndPrivacy**](CustomerApi.md#CustomerPrototypeUpdateTermsAndPrivacy) | **Post** /Customers/{id}/termsAndPrivacyAccepted | Update a terms and privacy acceptance for customer.
[**CustomerPrototypeUserInfo**](CustomerApi.md#CustomerPrototypeUserInfo) | **Get** /Customers/{id}/userInfo | Get customer details with userInfo access token.
[**CustomerPrototypeVerifyEmailPasswordlessToken**](CustomerApi.md#CustomerPrototypeVerifyEmailPasswordlessToken) | **Get** /Customers/{id}/passwordLessToken | Verifies the email token and activates tokens related to it. Returns verified text with redirect to \&quot;signup complete deep link\&quot;
[**CustomerPrototypeVlanServices**](CustomerApi.md#CustomerPrototypeVlanServices) | **Get** /Customers/{id}/locations/{locationId}/vlanServices | Returns vlanServices from Customer location state
[**CustomerPrototypeWpsState**](CustomerApi.md#CustomerPrototypeWpsState) | **Get** /Customers/{id}/locations/{locationId}/wpsState | Get WPS state
[**CustomerRefreshOauthAccessToken**](CustomerApi.md#CustomerRefreshOauthAccessToken) | **Post** /Customers/refreshOauthAccessToken | Refresh access and refresh tokens
[**CustomerRegister**](CustomerApi.md#CustomerRegister) | **Post** /Customers/register | Register/create an anonymous account with an accountId instead of with email/password.
[**CustomerRegisterWithGroups**](CustomerApi.md#CustomerRegisterWithGroups) | **Post** /Customers/registerWithGroups | Register/create an account with an accountId plus email/password/groups.
[**CustomerResendEmailVerification**](CustomerApi.md#CustomerResendEmailVerification) | **Post** /Customers/resendEmailVerification | Resend the verification email.
[**CustomerResetPassword**](CustomerApi.md#CustomerResetPassword) | **Post** /Customers/reset | Reset password for a user with email.
[**CustomerSearchFields**](CustomerApi.md#CustomerSearchFields) | **Get** /Customers/search/{keyword} | Search the keyword on a particular field such as \&quot;accountId\&quot;, \&quot;name\&quot;, \&quot;email\&quot;.


# **CustomerConfirm**
> CustomerConfirm(ctx, uid, token, optional)
Confirm a user registration with identity verification token.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **uid** | **string**|  | 
  **token** | **string**|  | 
 **optional** | ***CustomerApiCustomerConfirmOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CustomerApiCustomerConfirmOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **redirect** | **optional.String**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomerCount**
> InlineResponse2001 CustomerCount(ctx, optional)
Count instances of the model matched by where from the data source.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CustomerApiCustomerCountOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CustomerApiCustomerCountOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **where** | **optional.String**| Criteria to match model instances | 

### Return type

[**InlineResponse2001**](inline_response_200_1.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomerCreateOrUpdateUser**
> Customer CustomerCreateOrUpdateUser(ctx, email, name, roles, optional)
Create or update a Plume NOC user.

<div><strong>200</strong>: Success, user created.</div> <div><strong>400</strong>: Required fields missing.</div> <div><strong>422</strong>: Input validation failed.</div> <div><strong>500</strong>: Internal server error.</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **email** | **string**|  | 
  **name** | **string**|  | 
  **roles** | **string**|  | 
 **optional** | ***CustomerApiCustomerCreateOrUpdateUserOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CustomerApiCustomerCreateOrUpdateUserOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **groups** | **optional.String**|  | 

### Return type

[**Customer**](Customer.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomerCustomCreate**
> Customer CustomerCustomCreate(ctx, optional)
Create a Plume customer.

<div><strong>200</strong>: Success, customer created.</div> <div><strong>422</strong>: Input validation failed.</div> <div><strong>500</strong>: Internal server error.</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CustomerApiCustomerCustomCreateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CustomerApiCustomerCustomCreateOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **email** | **optional.String**|  | 
 **password** | **optional.String**|  | 
 **name** | **optional.String**|  | 
 **partnerId** | **optional.String**|  | 
 **person** | **optional.String**| Person object should contain field &#39;imageId&#39;  and object profile with field type (String) | 
 **location** | **optional.String**| Location object should contain field &#39;name&#39; (String) | 
 **notificationOptions** | **optional.String**|  | 
 **passwordLessToken** | **optional.Bool**|  | [default to false]
 **source** | **optional.String**|  | [default to customCreate]

### Return type

[**Customer**](Customer.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomerEmailExists**
> bool CustomerEmailExists(ctx, optional)


Check if customer email exists and is known to Plume, pass email as parameter to /Customers/exists?email=xxx@yyy.com <div><strong>200</strong>: customer email exists and is known to Plume, emailVerified returned</div> <div><strong>400</strong>: email is required</div> <div><strong>404</strong>: customer email does not exist and is not known to Plume</div> <div><strong>422</strong>: email is not valid</div> <div><strong>500</strong>: internal server error</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CustomerApiCustomerEmailExistsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CustomerApiCustomerEmailExistsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **email** | **optional.String**|  | 

### Return type

**bool**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomerEmailPasswordlessToken**
> AccessToken CustomerEmailPasswordlessToken(ctx, email, optional)
Generate two accessTokens with special scopes for the account with the email address and send a verification email.

<div><strong>200</strong>: Success, return new appToken, refreshToken and send out the email with emailToken.</div> <div><strong>401</strong>: Authorization required or customer id not found</div> <div><strong>422</strong>: Email must be defined and valid.</div> <div><strong>500</strong>: internal server error</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **email** | **string**|  | 
 **optional** | ***CustomerApiCustomerEmailPasswordlessTokenOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CustomerApiCustomerEmailPasswordlessTokenOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **notificationOptions** | **optional.String**|  | 

### Return type

[**AccessToken**](AccessToken.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomerFind**
> []Customer CustomerFind(ctx, optional)
Find all instances of the model matched by filter from the data source.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CustomerApiCustomerFindOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CustomerApiCustomerFindOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filter** | **optional.String**| Filter defining fields, where, include, order, offset, and limit - must be a JSON-encoded string (&#x60;{\&quot;where\&quot;:{\&quot;something\&quot;:\&quot;value\&quot;}}&#x60;).  See https://loopback.io/doc/en/lb3/Querying-data.html#using-stringified-json-in-rest-queries for more details. | 

### Return type

[**[]Customer**](Customer.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomerFindById**
> Customer CustomerFindById(ctx, id, optional)
Find a model instance by {{id}} from the data source.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Model id | 
 **optional** | ***CustomerApiCustomerFindByIdOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CustomerApiCustomerFindByIdOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **filter** | **optional.String**| Filter defining fields and include - must be a JSON-encoded string ({\&quot;something\&quot;:\&quot;value\&quot;}) | 

### Return type

[**Customer**](Customer.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomerImportData**
> interface{} CustomerImportData(ctx, data, migratedFrom, reason)
Import customer data

<div><strong>204</strong>: Success.</div> <div><strong>400</strong>: Nothing to import.</div> <div><strong>422</strong>: Import data is invalid.</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | **string**|  | 
  **migratedFrom** | **string**|  | 
  **reason** | **string**|  | 

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomerLogin**
> interface{} CustomerLogin(ctx, credentials, optional)
Login a user with username/email and password.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **credentials** | **interface{}**|  | 
 **optional** | ***CustomerApiCustomerLoginOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CustomerApiCustomerLoginOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **include** | **optional.String**| Related objects to include in the response. See the description of return value for more details. | 

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomerLogoutPostCustomersLogout**
> CustomerLogoutPostCustomersLogout(ctx, )
Logout a user with access token.

### Required Parameters
This endpoint does not need any parameter.

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomerPrototypeAddDeviceVisibleToGuests**
> []string CustomerPrototypeAddDeviceVisibleToGuests(ctx, id, locationId, mac)
DEPRECATED: Update home devices visible to guests.

<div><strong>200</strong>: Success, devicesVisibleToGuests returned.</div> <div><strong>400</strong>: Required fields missing.</div> <div><strong>401</strong>: Authorization required or customer id not found.</div> <div><strong>404</strong>: Customer id, location id, or WifiNetwork does not exist.</div> <div><strong>422</strong>: Device mac validation failed.</div> <div><strong>500</strong>: Internal server error</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Customer id | 
  **locationId** | **string**|  | 
  **mac** | **string**| mac to be added | 

### Return type

**[]string**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomerPrototypeAppFacadeHome**
> AppFacadeHomeResponse CustomerPrototypeAppFacadeHome(ctx, id, locationId, optional)
Retrieve timezone, capabilities, summary, ... for this location.

<div><strong>200</strong>: Success, an array of properties returned.</div> <div><strong>401</strong>: Authorization required or customer id not found</div> <div><strong>404</strong>: customer id or location id does not exist and is not known to Plume</div> <div><strong>500</strong>: internal server error</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Customer id | 
  **locationId** | **string**|  | 
 **optional** | ***CustomerApiCustomerPrototypeAppFacadeHomeOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CustomerApiCustomerPrototypeAppFacadeHomeOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **filters** | **optional.String**|  | 
 **daysOffline** | **optional.Float64**|  | 

### Return type

[**AppFacadeHomeResponse**](AppFacadeHomeResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomerPrototypeApproveDevices**
> []string CustomerPrototypeApproveDevices(ctx, id, locationId, networkId, macs)
Approve devices in the network

<div><strong>204</strong>: Success.</div> <div><strong>404</strong>: Location does not exist.</div> <div><strong>404</strong>: Network does not exist.</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Customer id | 
  **locationId** | **string**|  | 
  **networkId** | **string**|  | 
  **macs** | [**[]XAny**](x-any.md)|  | 

### Return type

**[]string**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomerPrototypeApproveWhitelistRequest**
> CustomerPrototypeApproveWhitelistRequest(ctx, id, locationId, requestId, persons)
Approve a persons whitelist request and add it to the security policy.

<div><strong>204</strong>: No content.</div> <div><strong>400</strong>: Required fields missing or field type is incorrect.</div> <div><strong>401</strong>: Authorization required or customer id not found</div> <div><strong>404</strong>: Location id, CustomerId or requst id does not exist and is not known to Plume</div> <div><strong>500</strong>: Internal server error</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Customer id | 
  **locationId** | **string**|  | 
  **requestId** | **string**|  | 
  **persons** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomerPrototypeBlockDevices**
> []string CustomerPrototypeBlockDevices(ctx, id, locationId, macs)
Block devices

<div><strong>200</strong>: Success.</div> <div><strong>404</strong>: Location does not exist.</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Customer id | 
  **locationId** | **string**|  | 
  **macs** | [**[]XAny**](x-any.md)|  | 

### Return type

**[]string**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomerPrototypeClaimNode**
> NodeClaimingResponse CustomerPrototypeClaimNode(ctx, id, locationId, optional)
Claim a node and all nodes still associated to its Package ID for a Location ID.

<div><strong>200</strong>: King node claimed and all related claimed nodes are returned.</div> <div><strong>204</strong>: Valid serial number but zero new claimed nodes.</div> <div><strong>404</strong>: Unable to find Node with serial number, customer id, or location id.</div> <div><strong>409</strong>: Node is owned by another customer.</div> <div><strong>422</strong>: Claiming request exceeded numPodsAuthorized (=leaf pods), accountId+partnerId not unique, and/or monitorMode=true.</div> <div><strong>500</strong>: internal server error</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Customer id | 
  **locationId** | **string**| location id | 
 **optional** | ***CustomerApiCustomerPrototypeClaimNodeOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CustomerApiCustomerPrototypeClaimNodeOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **serialNumber** | **optional.String**| unique serial number or ID of Node | 
 **radioMac24** | **optional.String**| optional but required for auto-importing, must be a valid mac address | 
 **radioMac50** | **optional.String**| optional but required for auto-importing, must be a valid mac address | 
 **radioMac60** | **optional.String**| optional but required for auto-importing, must be a valid mac address | 
 **ethernetMac** | **optional.String**| optional but required for auto-importing, must be a valid mac address | 
 **ethernet1Mac** | **optional.String**| optional but required for auto-importing, must be a valid mac address | 
 **claimKey** | **optional.String**| optional but required for auto-importing, must be a valid claimKey | 
 **model** | **optional.String**| optional when auto-importing, ignored otherwise | 
 **hybridCheck** | **optional.Bool**| optional when auto-importing, ignored otherwise | 
 **nickname** | **optional.String**| optional node nickname | 
 **skipSubscription** | **optional.Bool**| skip subscription update | [default to false]
 **backhaulDhcpPoolIdx** | **optional.Float64**| optional node backhaulDhcpPoolIdx | 
 **room** | **optional.String**| optional room identifier | 

### Return type

[**NodeClaimingResponse**](NodeClaimingResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomerPrototypeCreateAccessTokens**
> AccessToken CustomerPrototypeCreateAccessTokens(ctx, id, optional)
Creates a new instance in accessTokens of this model.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Customer id | 
 **optional** | ***CustomerApiCustomerPrototypeCreateAccessTokensOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CustomerApiCustomerPrototypeCreateAccessTokensOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **data** | [**optional.Interface of AccessToken**](AccessToken.md)|  | 

### Return type

[**AccessToken**](AccessToken.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomerPrototypeCreateGetMarketingExportDataAccessToken**
> AccessToken CustomerPrototypeCreateGetMarketingExportDataAccessToken(ctx, id, optional)
Create access token to get marketing data by CRM for campaigns

<div><strong>200</strong>: Success, accessToken returned.</div> <div><strong>401</strong>: Authorization required or customer id not found</div> <div><strong>404</strong>: customer id does not exist and is not known to Plume</div> <div><strong>500</strong>: internal server error</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Customer id | 
 **optional** | ***CustomerApiCustomerPrototypeCreateGetMarketingExportDataAccessTokenOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CustomerApiCustomerPrototypeCreateGetMarketingExportDataAccessTokenOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **ttl** | **optional.Float64**|  | 

### Return type

[**AccessToken**](AccessToken.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomerPrototypeCreateIpLimitedAccessToken**
> AccessToken CustomerPrototypeCreateIpLimitedAccessToken(ctx, id, optional)
Create access token with limited privileges as defined for IP authenticated customers

<div><strong>200</strong>: Success, response object returned.</div> <div><strong>401</strong>: Authorization required or customer id not found</div> <div><strong>404</strong>: customer id does not exist and is not known to Plume</div> <div><strong>500</strong>: internal server error</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Customer id | 
 **optional** | ***CustomerApiCustomerPrototypeCreateIpLimitedAccessTokenOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CustomerApiCustomerPrototypeCreateIpLimitedAccessTokenOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **ttl** | **optional.Float64**|  | 

### Return type

[**AccessToken**](AccessToken.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomerPrototypeCreateLocation**
> XAny CustomerPrototypeCreateLocation(ctx, id, name, optional)
Create a new location.

<div><strong>200</strong>: Success, updated.</div> <div><strong>400</strong>: Required field(the location name) missing.</div> <div><strong>401</strong>: Authorization required or customer id not found.</div> <div><strong>500</strong>: Internal server error</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Customer id | 
  **name** | **string**|  | 
 **optional** | ***CustomerApiCustomerPrototypeCreateLocationOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CustomerApiCustomerPrototypeCreateLocationOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **profile** | **optional.String**|  | 

### Return type

[**XAny**](x-any.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomerPrototypeCreateMigration**
> Migration CustomerPrototypeCreateMigration(ctx, id, optional)
Creates a new instance in _migration of this model.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Customer id | 
 **optional** | ***CustomerApiCustomerPrototypeCreateMigrationOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CustomerApiCustomerPrototypeCreateMigrationOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **data** | [**optional.Interface of Migration**](Migration.md)|  | 

### Return type

[**Migration**](Migration.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomerPrototypeCreateNewPasswordlessToken**
> AccessToken CustomerPrototypeCreateNewPasswordlessToken(ctx, id)
Generates usable passwordless accessToken for the account with the email address.

<div><strong>204</strong>: Success, return new appToken.</div> <div><strong>401</strong>: Authorization required or customer id not found</div> <div><strong>500</strong>: internal server error</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Customer id | 

### Return type

[**AccessToken**](AccessToken.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomerPrototypeCreateOauthAccessToken**
> AccessToken CustomerPrototypeCreateOauthAccessToken(ctx, id, optional)
Create access token with ouath scope.

<div><strong>200</strong>: Success, access token created and returned.</div> <div><strong>401</strong>: Authorization Required.</div> <div><strong>500</strong>: Internal server error.</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Customer id | 
 **optional** | ***CustomerApiCustomerPrototypeCreateOauthAccessTokenOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CustomerApiCustomerPrototypeCreateOauthAccessTokenOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **scope** | **optional.String**|  | 
 **ttlSeconds** | **optional.Float64**|  | 
 **singleToken** | **optional.Bool**|  | 

### Return type

[**AccessToken**](AccessToken.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomerPrototypeCreateOauthRefreshToken**
> AccessToken CustomerPrototypeCreateOauthRefreshToken(ctx, id, optional)
Create refresh token for a specific access token with ouath scope.

<div><strong>200</strong>: Success, access token created and returned.</div> <div><strong>401</strong>: Authorization Required.</div> <div><strong>500</strong>: Internal server error.</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Customer id | 
 **optional** | ***CustomerApiCustomerPrototypeCreateOauthRefreshTokenOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CustomerApiCustomerPrototypeCreateOauthRefreshTokenOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **linkedAccessToken** | **optional.String**|  | 
 **prevLinkedRefreshTokenId** | **optional.String**|  | 
 **clientId** | **optional.String**|  | 

### Return type

[**AccessToken**](AccessToken.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomerPrototypeCreatePatchServiceLevelAccessToken**
> AccessToken CustomerPrototypeCreatePatchServiceLevelAccessToken(ctx, id, optional)
Create access token to patch customer serviceLevel used by ZUORA

<div><strong>200</strong>: Success, accessToken returned.</div> <div><strong>401</strong>: Authorization required or customer id not found</div> <div><strong>404</strong>: customer id does not exist and is not known to Plume</div> <div><strong>500</strong>: internal server error</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Customer id | 
 **optional** | ***CustomerApiCustomerPrototypeCreatePatchServiceLevelAccessTokenOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CustomerApiCustomerPrototypeCreatePatchServiceLevelAccessTokenOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **ttl** | **optional.Float64**|  | 

### Return type

[**AccessToken**](AccessToken.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomerPrototypeCreateReadDnsAccessToken**
> AccessToken CustomerPrototypeCreateReadDnsAccessToken(ctx, id)
Create access token to read data related to DNS security policies

<div><strong>200</strong>: Success, accessToken returned.</div> <div><strong>401</strong>: Authorization required or customer id not found</div> <div><strong>404</strong>: customer id does not exist and is not known to Plume</div> <div><strong>500</strong>: internal server error</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Customer id | 

### Return type

[**AccessToken**](AccessToken.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomerPrototypeDeleteAllDeviceFreezes**
> CustomerPrototypeDeleteAllDeviceFreezes(ctx, id, locationId, mac)
Delete/clear all device freezes templateIds for a mac.

<div><strong>204</strong>: Success, updated.</div> <div><strong>401</strong>: Authorization required or customer id not found.</div> <div><strong>404</strong>: Location id, does not exist.</div> <div><strong>500</strong>: Internal server error</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Customer id | 
  **locationId** | **string**|  | 
  **mac** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomerPrototypeDeleteAppPrioritizationLocationConfig**
> interface{} CustomerPrototypeDeleteAppPrioritizationLocationConfig(ctx, id, locationId)
Set custom setting to default for app prioritization.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Customer id | 
  **locationId** | **string**|  | 

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomerPrototypeDeleteCaptivePortal**
> CustomerPrototypeDeleteCaptivePortal(ctx, id, locationId, networkId)
Delete a CaptivePortal for a Location ID.

<div><strong>200</strong>: Success.</div> <div><strong>401</strong>: Authorization required or customer id not found</div> <div><strong>404</strong>: Location id/NetworkId does not exist and is not known to Plume</div> <div><strong>500</strong>: Internal server error</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Customer id | 
  **locationId** | **string**|  | 
  **networkId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomerPrototypeDeleteCaptivePortalAuthorizedClients**
> CustomerPrototypeDeleteCaptivePortalAuthorizedClients(ctx, id, locationId, networkId, mac)
Delete Authorized Client

<div><strong>200</strong>: Success.</div> <div><strong>401</strong>: Authorization required or customer id not found</div> <div><strong>404</strong>: Location id/NetworkId does not exist and is not known to Plume</div> <div><strong>500</strong>: Internal server error</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Customer id | 
  **locationId** | **string**|  | 
  **networkId** | **string**|  | 
  **mac** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomerPrototypeDeleteConfigs**
> CustomerPrototypeDeleteConfigs(ctx, id, locationId, optional)
Delete specified location settings, while keeping claimed nodes intact

<div><strong>204</strong>: Success, a job well done.</div> <div><strong>401</strong>: Authorization required </div> <div><strong>404</strong>: location id not found or nodeId missing from URL <div><strong>500</strong>: internal server error</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Customer id | 
  **locationId** | **string**|  | 
 **optional** | ***CustomerApiCustomerPrototypeDeleteConfigsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CustomerApiCustomerPrototypeDeleteConfigsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **persons** | **optional.Bool**| Whether or not to delete person information | [default to false]
 **onboardingCheckpoints** | **optional.Bool**| Whether or not to delete onboarding checkpoints | [default to false]
 **devices** | **optional.Bool**| Whether or not to delete devices related information | [default to false]
 **networkConfiguration** | **optional.Bool**| Whether or not to delete network configuration | [default to false]
 **wifiNetwork** | **optional.Bool**| Whether or not to delete wifi network | [default to false]
 **deviceFreeze** | **optional.Bool**| Whether or not to delete device freeze templates | [default to false]
 **deviceNicknames** | **optional.Bool**| Whether or not to delete device nicknames | [default to false]
 **managers** | **optional.Bool**| Whether or not to delete managers of the location | [default to false]
 **wanConfiguration** | **optional.Bool**| Whether or not to delete wanConfiguration | [default to false]

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomerPrototypeDeleteCustomSharedSchedule**
> CustomerPrototypeDeleteCustomSharedSchedule(ctx, id, locationId, templateId)
Delete \"custom shared\" schedule shared by all persons and devices in a location.

<div><strong>204</strong>: Success, the custom shared schedule deleted.</div> <div><strong>401</strong>: Authorization required or customer id not found</div> <div><strong>404</strong>: Location id does not exist or is not known to Plume</div> <div><strong>500</strong>: internal server error</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Customer id | 
  **locationId** | **string**|  | 
  **templateId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomerPrototypeDeleteCustomer**
> interface{} CustomerPrototypeDeleteCustomer(ctx, id)
Delete a model instance by {{id}} from the data source.

<div><strong>200</strong>: Success, customer details returned.</div> <div><strong>401</strong>: Authorization Required.</div> <div><strong>500</strong>: Internal server error.</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Customer id | 

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomerPrototypeDeleteDeviceFreeze**
> CustomerPrototypeDeleteDeviceFreeze(ctx, id, locationId, mac, freezeTemplateId)
Delete a device to be frozen for a Location ID.

<div><strong>204</strong>: Success, updated.</div> <div><strong>401</strong>: Authorization required or customer id not found.</div> <div><strong>404</strong>: Location id, does not exist.</div> <div><strong>422</strong>: MAC address does not exist or is invalid.</div> <div><strong>500</strong>: Internal server error</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Customer id | 
  **locationId** | **string**|  | 
  **mac** | **string**|  | 
  **freezeTemplateId** | **string**| Valid templates are &#39;untilMidinight&#39;, &#39;schoolNights&#39;, etc. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomerPrototypeDeleteDeviceFreezeAutoExpire**
> CustomerPrototypeDeleteDeviceFreezeAutoExpire(ctx, id, locationId, mac)
Delete a device to be frozen for a Location ID.

<div><strong>204</strong>: Success, updated.</div> <div><strong>401</strong>: Authorization required or customer id not found.</div> <div><strong>404</strong>: Location id, does not exist.</div> <div><strong>500</strong>: Internal server error</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Customer id | 
  **locationId** | **string**|  | 
  **mac** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomerPrototypeDeleteDeviceFreezeForever**
> CustomerPrototypeDeleteDeviceFreezeForever(ctx, id, locationId, mac)
Delete a device forever freeze for a Location ID.

<div><strong>204</strong>: Success, updated.</div> <div><strong>401</strong>: Authorization required or customer id not found.</div> <div><strong>404</strong>: Location id, does not exist.</div> <div><strong>422</strong>: MAC address does not exist or is invalid.</div> <div><strong>500</strong>: Internal server error</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Customer id | 
  **locationId** | **string**|  | 
  **mac** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomerPrototypeDeleteDeviceFreezeResidentialGwManaged**
> CustomerPrototypeDeleteDeviceFreezeResidentialGwManaged(ctx, id, locationId, mac)
Delete a device residentialGwManaged freeze for a Location ID.

<div><strong>204</strong>: Success, updated.</div> <div><strong>401</strong>: Authorization required or customer id not found.</div> <div><strong>404</strong>: Location id, does not exist.</div> <div><strong>422</strong>: MAC address does not exist or is invalid.</div> <div><strong>500</strong>: Internal server error</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Customer id | 
  **locationId** | **string**|  | 
  **mac** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomerPrototypeDeleteDeviceFreezeSuspend**
> CustomerPrototypeDeleteDeviceFreezeSuspend(ctx, id, locationId, mac)
Delete a device suspend for a Location ID.

<div><strong>204</strong>: Success, updated.</div> <div><strong>401</strong>: Authorization required or customer id not found.</div> <div><strong>404</strong>: Location id, does not exist.</div> <div><strong>422</strong>: MAC address does not exist or is invalid.</div> <div><strong>500</strong>: Internal server error</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Customer id | 
  **locationId** | **string**|  | 
  **mac** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomerPrototypeDeleteDeviceFromAccessZone**
> XAny CustomerPrototypeDeleteDeviceFromAccessZone(ctx, id, locationId, zoneId, mac)
Delete a device mac from a WiFi Access Zone

<div><strong>200</strong>: Success, all access zones returned</div> <div><strong>400</strong>: Required fields missing</div> <div><strong>401</strong>: Authorization required or customer id not found</div> <div><strong>404</strong>: Customer id, location id, or WifiNetwork does not exist</div> <div><strong>422</strong>: Validation failed</div> <div><strong>500</strong>: Internal server error</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Customer id | 
  **locationId** | **string**|  | 
  **zoneId** | **float64**| id of access zone | 
  **mac** | **string**| the device mac to be added to the access zone | 

### Return type

[**XAny**](x-any.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomerPrototypeDeleteDeviceFromPerson**
> CustomerPrototypeDeleteDeviceFromPerson(ctx, id, locationId, personId, mac)
Unassign a device from Person for a location ID.

<div><strong>204</strong>: Success.</div> <div><strong>401</strong>: Authorization required or customer id not found</div> <div><strong>404</strong>: Location id, Person id, or mac does not exist and is not known to Plume</div> <div><strong>500</strong>: Internal server error</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Customer id | 
  **locationId** | **string**|  | 
  **personId** | **string**|  | 
  **mac** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomerPrototypeDeleteDeviceGroup**
> CustomerPrototypeDeleteDeviceGroup(ctx, id, locationId, networkId, groupId)
Delete a device group from a network.

<div><strong>200</strong>: Success.</div> <div><strong>422</strong>: Schema validation failed.</div> <div><strong>404</strong>: Location does not exist.</div> <div><strong>404</strong>: Network does not exist.</div> <div><strong>403</strong>: Not allowed to delete standalone groups or groups in unsupported networks.</div> <div><strong>401</strong>: Unauthorized.</div> <div><strong>400</strong>: Invalid JSON or missing arguments.</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Customer id | 
  **locationId** | **string**|  | 
  **networkId** | **string**|  | 
  **groupId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomerPrototypeDeleteDeviceQosPrioritization**
> []XAny CustomerPrototypeDeleteDeviceQosPrioritization(ctx, id, locationId, mac)
Delete prioritization of a single device

<div><strong>202</strong>: Success.</div> <div><strong>404</strong>: Location does not exist.</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Customer id | 
  **locationId** | **string**|  | 
  **mac** | **string**|  | 

### Return type

[**[]XAny**](x-any.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomerPrototypeDeleteDeviceSecurityPolicyAnomalyExperience**
> CustomerPrototypeDeleteDeviceSecurityPolicyAnomalyExperience(ctx, id, locationId, mac)
Delete an Anomaly Experience (demo) for a Device on a location.

<div><strong>204</strong>: Success.</div> <div><strong>401</strong>: Authorization required or customer id not found</div> <div><strong>404</strong>: Location id, WifiNetwork, Device does not exist and is not known to Plume</div> <div><strong>500</strong>: Internal server error</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Customer id | 
  **locationId** | **string**|  | 
  **mac** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomerPrototypeDeleteDeviceSecurityPolicyAnomalyWhitelist**
> CustomerPrototypeDeleteDeviceSecurityPolicyAnomalyWhitelist(ctx, id, locationId, mac, fqdn)
Update a Location's Anomaly Security Policy for a location ID to remove a whitelisted DNS entry.

<div><strong>204</strong>: Success.</div> <div><strong>401</strong>: Authorization required or customer id not found</div> <div><strong>404</strong>: Location id, WifiNetwork, Device or DNS does not exist and is not known to Plume</div> <div><strong>500</strong>: Internal server error</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Customer id | 
  **locationId** | **string**|  | 
  **mac** | **string**|  | 
  **fqdn** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomerPrototypeDeleteDhcpReservation**
> []DhcpReservation CustomerPrototypeDeleteDhcpReservation(ctx, id, locationId, mac)
Delete a current DHCP IP reservation and the associated port forwarding entries for a particular MAC address at a Location ID.

<div><strong>200</strong>: Success, remaining DhcpReservations are returned.</div> <div><strong>401</strong>: Authorization required or customer id not found</div> <div><strong>404</strong>: NetworkConfiguration or DhcpReservation is empty.</div> <div><strong>422</strong>: mac is empty or invalid.</div> <div><strong>500</strong>: internal server error</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Customer id | 
  **locationId** | **string**|  | 
  **mac** | **string**|  | 

### Return type

[**[]DhcpReservation**](DhcpReservation.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomerPrototypeDeleteForcedSteer**
> interface{} CustomerPrototypeDeleteForcedSteer(ctx, id, locationId, mac)
Disable 2.4Ghz band enforcement early.

<div><strong>204</strong>: Success, forced steer ended early.</div> <div><strong>404</strong>: Location ID or Device mac not found or the device has not been online in the last 31 days</div> <div><strong>500</strong>: Internal server error</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Customer id | 
  **locationId** | **string**| locationId | 
  **mac** | **string**| MAC address of the target device. Must have been online in the last 31 days. | 

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomerPrototypeDeleteFromDeviceSecurityPolicyWebsitesBlacklist**
> CustomerPrototypeDeleteFromDeviceSecurityPolicyWebsitesBlacklist(ctx, id, locationId, mac, dns)
Update a Device's Security Policy for a location ID to remove a blacklisted DNS entry.

<div><strong>204</strong>: Success.</div> <div><strong>401</strong>: Authorization required or customer id not found</div> <div><strong>404</strong>: Location id, WifiNetwork, Device, or DNS does not exist and is not known to Plume</div> <div><strong>500</strong>: Internal server error</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Customer id | 
  **locationId** | **string**|  | 
  **mac** | **string**|  | 
  **dns** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomerPrototypeDeleteFromDeviceSecurityPolicyWebsitesWhitelist**
> CustomerPrototypeDeleteFromDeviceSecurityPolicyWebsitesWhitelist(ctx, id, locationId, mac, dns)
Update a Device's Security Policy for a location ID to remove a whitelisted DNS entry.

<div><strong>204</strong>: Success.</div> <div><strong>401</strong>: Authorization required or customer id not found</div> <div><strong>404</strong>: Location id, WifiNetwork, Device, or DNS does not exist and is not known to Plume</div> <div><strong>500</strong>: Internal server error</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Customer id | 
  **locationId** | **string**|  | 
  **mac** | **string**|  | 
  **dns** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomerPrototypeDeleteFromGroupOfUnassignedDevicesSecurityPolicyWebsitesBlacklist**
> CustomerPrototypeDeleteFromGroupOfUnassignedDevicesSecurityPolicyWebsitesBlacklist(ctx, id, locationId, dns)
Update a Location's Default Device Group Security Policy for a location ID to remove a blacklisted DNS entry.

<div><strong>204</strong>: Success.</div> <div><strong>401</strong>: Authorization required or customer id not found</div> <div><strong>404</strong>: Location id, WifiNetwork, or DNS does not exist and is not known to Plume</div> <div><strong>500</strong>: Internal server error</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Customer id | 
  **locationId** | **string**|  | 
  **dns** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomerPrototypeDeleteFromGroupOfUnassignedDevicesSecurityPolicyWebsitesWhitelist**
> CustomerPrototypeDeleteFromGroupOfUnassignedDevicesSecurityPolicyWebsitesWhitelist(ctx, id, locationId, dns)
Update a Location's Default Device Group Security Policy for a location ID to remove a whitelisted DNS entry.

<div><strong>204</strong>: Success.</div> <div><strong>401</strong>: Authorization required or customer id not found</div> <div><strong>404</strong>: Location id, WifiNetwork, or DNS does not exist and is not known to Plume</div> <div><strong>500</strong>: Internal server error</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Customer id | 
  **locationId** | **string**|  | 
  **dns** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomerPrototypeDeleteFromLocationSecurityPolicyWebsitesBlacklist**
> CustomerPrototypeDeleteFromLocationSecurityPolicyWebsitesBlacklist(ctx, id, locationId, dns)
Update a Location's Security Policy for a location ID to remove a blacklisted DNS entry.

<div><strong>204</strong>: Success.</div> <div><strong>401</strong>: Authorization required or customer id not found</div> <div><strong>404</strong>: Location id, WifiNetwork, or DNS does not exist and is not known to Plume</div> <div><strong>500</strong>: Internal server error</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Customer id | 
  **locationId** | **string**|  | 
  **dns** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomerPrototypeDeleteFromLocationSecurityPolicyWebsitesWhitelist**
> CustomerPrototypeDeleteFromLocationSecurityPolicyWebsitesWhitelist(ctx, id, locationId, dns)
Update a Locations's Security Policy for a location ID to remove a whitelisted DNS entry.

<div><strong>204</strong>: Success.</div> <div><strong>401</strong>: Authorization required or customer id not found</div> <div><strong>404</strong>: Location id, WifiNetwork, or DNS does not exist and is not known to Plume</div> <div><strong>500</strong>: Internal server error</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Customer id | 
  **locationId** | **string**|  | 
  **dns** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomerPrototypeDeleteFromPersonSecurityPolicyWebsitesBlacklist**
> CustomerPrototypeDeleteFromPersonSecurityPolicyWebsitesBlacklist(ctx, id, locationId, personId, dns)
Update a Person's Security Policy for a location ID to remove a blacklisted DNS entry.

<div><strong>204</strong>: Success.</div> <div><strong>401</strong>: Authorization required or customer id not found</div> <div><strong>404</strong>: Location id, WifiNetwork, Person id, or DNS does not exist and is not known to Plume</div> <div><strong>500</strong>: Internal server error</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Customer id | 
  **locationId** | **string**|  | 
  **personId** | **string**|  | 
  **dns** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomerPrototypeDeleteFromPersonSecurityPolicyWebsitesWhitelist**
> CustomerPrototypeDeleteFromPersonSecurityPolicyWebsitesWhitelist(ctx, id, locationId, personId, dns)
Update a Person's Security Policy for a location ID to remove a whitelisted DNS entry.

<div><strong>204</strong>: Success.</div> <div><strong>401</strong>: Authorization required or customer id not found</div> <div><strong>404</strong>: Location id, WifiNetwork, Person id, or DNS does not exist and is not known to Plume</div> <div><strong>500</strong>: Internal server error</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Customer id | 
  **locationId** | **string**|  | 
  **personId** | **string**|  | 
  **dns** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomerPrototypeDeleteFrontHaul**
> CustomerPrototypeDeleteFrontHaul(ctx, id, locationId, networkId)
Delete a Front Haul for a Location ID.

<div><strong>200</strong>: Success.</div> <div><strong>401</strong>: Authorization required or customer id not found</div> <div><strong>404</strong>: Location id/NetworkId does not exist and is not known to Plume</div> <div><strong>500</strong>: Internal server error</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Customer id | 
  **locationId** | **string**|  | 
  **networkId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomerPrototypeDeleteGdprCaptivePortalsData**
> interface{} CustomerPrototypeDeleteGdprCaptivePortalsData(ctx, id, locationId, networkId, email)
Delete the Gdpr Captive Portals data for a guest.

<div><strong>200</strong>: Success, GDPR Captive Portals data deleted.</div> <div><strong>401</strong>: Authorization required or customer id not found</div> <div><strong>404</strong>: location id or secondary networks does not exist.</div> <div><strong>500</strong>: internal server error</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Customer id | 
  **locationId** | **string**|  | 
  **networkId** | **string**|  | 
  **email** | **string**|  | 

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomerPrototypeDeleteGroupOfUnassignedDevicesFreezeAutoExpire**
> CustomerPrototypeDeleteGroupOfUnassignedDevicesFreezeAutoExpire(ctx, id, locationId)
Delete GroupOfUnassignedDevices autoExpire freeze for a Location ID.

<div><strong>204</strong>: Success, updated.</div> <div><strong>401</strong>: Authorization required or customer id not found.</div> <div><strong>404</strong>: Location id, does not exist.</div> <div><strong>500</strong>: Internal server error</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Customer id | 
  **locationId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomerPrototypeDeleteGroupOfUnassignedDevicesFreezeForever**
> CustomerPrototypeDeleteGroupOfUnassignedDevicesFreezeForever(ctx, id, locationId)
Delete GroupOfUnassignedDevices forever freeze for a Location ID.

<div><strong>204</strong>: Success, updated.</div> <div><strong>401</strong>: Authorization required or customer id not found.</div> <div><strong>404</strong>: Location id, does not exist.</div> <div><strong>500</strong>: Internal server error</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Customer id | 
  **locationId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomerPrototypeDeleteGroupOfUnassignedDevicesFreezeSuspend**
> CustomerPrototypeDeleteGroupOfUnassignedDevicesFreezeSuspend(ctx, id, locationId)
Delete GroupOfUnassignedDevices suspend for a Location ID.

<div><strong>204</strong>: Success, updated.</div> <div><strong>401</strong>: Authorization required or customer id not found.</div> <div><strong>404</strong>: Location id, does not exist.</div> <div><strong>500</strong>: Internal server error</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Customer id | 
  **locationId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomerPrototypeDeleteGroupOfUnassignedDevicesFreezeTemplateId**
> CustomerPrototypeDeleteGroupOfUnassignedDevicesFreezeTemplateId(ctx, id, locationId, freezeTemplateId)
Delete GroupOfUnassignedDevices uuid freeze for a Location ID.

<div><strong>204</strong>: Success, updated.</div> <div><strong>401</strong>: Authorization required or customer id not found.</div> <div><strong>404</strong>: Location id, does not exist.</div> <div><strong>500</strong>: Internal server error</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Customer id | 
  **locationId** | **string**|  | 
  **freezeTemplateId** | **string**| Valid templates are uuids | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomerPrototypeDeleteGroupOfUnassignedDevicesFreezes**
> CustomerPrototypeDeleteGroupOfUnassignedDevicesFreezes(ctx, id, locationId)
Delete All GroupOfUnassignedDevices freeze except autoExpire for a Location ID.

<div><strong>204</strong>: Success, updated.</div> <div><strong>401</strong>: Authorization required or customer id not found.</div> <div><strong>404</strong>: Location id, does not exist.</div> <div><strong>500</strong>: Internal server error</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Customer id | 
  **locationId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomerPrototypeDeleteKvConfigs**
> KvConfig CustomerPrototypeDeleteKvConfigs(ctx, id, locationId, nodeId, module, key)
Delete kvConfigs with selected module and key on a particular Node for a Location ID.

<div><strong>200</strong>: Success, your new info looks good.</div> <div><strong>401</strong>: Authorization required or customer id not found</div> <div><strong>404</strong>: location id does not exist.</div> <div><strong>422</strong>: nodeId must be defined.</div> <div><strong>425</strong>: nodeId must belong to the location.</div> <div><strong>500</strong>: internal server error</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Customer id | 
  **locationId** | **string**|  | 
  **nodeId** | **string**|  | 
  **module** | **string**|  | 
  **key** | **string**|  | 

### Return type

[**KvConfig**](KvConfig.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomerPrototypeDeleteLinkedAccount**
> interface{} CustomerPrototypeDeleteLinkedAccount(ctx, id, provider, userId)
link the outside account, such as Samsung user.

<div><strong>200</strong>: Success, the outside account inserted into the customer info/object.</div> <div><strong>401</strong>: Authorization required or customer id not found</div> <div><strong>500</strong>: internal server error</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Customer id | 
  **provider** | **string**|  | 
  **userId** | **string**|  | 

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomerPrototypeDeleteLocation**
> CustomerPrototypeDeleteLocation(ctx, id, locationId)
Archive a location.

<div><strong>204</strong>: Success, location archived.</div> <div><strong>401</strong>: Authorization required or customer id not found.</div> <div><strong>404</strong>: Location id, does not exist.</div> <div><strong>409</strong>: Location already archived.</div> <div><strong>500</strong>: Internal server error</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Customer id | 
  **locationId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomerPrototypeDeleteLocationEventsHistory**
> CustomerPrototypeDeleteLocationEventsHistory(ctx, id, locationId, categories, optional)
Delete a Location's Security Events history for a location ID.

<div><strong>204</strong>: Success.</div> <div><strong>401</strong>: Authorization required or customer id not found</div> <div><strong>404</strong>: Location id, WifiNetwork, Device or DNS does not exist and is not known to Plume</div> <div><strong>500</strong>: Internal server error</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Customer id | 
  **locationId** | **string**|  | 
  **categories** | **string**|  | 
 **optional** | ***CustomerApiCustomerPrototypeDeleteLocationEventsHistoryOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CustomerApiCustomerPrototypeDeleteLocationEventsHistoryOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **reason** | **optional.String**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomerPrototypeDeleteLocationFreezeAutoExpire**
> XAny CustomerPrototypeDeleteLocationFreezeAutoExpire(ctx, id, locationId)
Delete the location freeze/autoExpire for a Location ID.

<div><strong>200</strong>: Success, updated.</div> <div><strong>401</strong>: Authorization required or customer id not found.</div> <div><strong>404</strong>: Location id, does not exist.</div> <div><strong>500</strong>: Internal server error</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Customer id | 
  **locationId** | **string**|  | 

### Return type

[**XAny**](x-any.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomerPrototypeDeleteManagerAccess**
> interface{} CustomerPrototypeDeleteManagerAccess(ctx, id, locationId, managerId)
Delete manager access for location and destroy access tokens for that manager\".

<div><strong>204</strong>: Success.</div> <div><strong>400</strong>: Required fields missing or field type is incorrect.</div> <div><strong>404</strong>: Location or Manager does not exist.</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Customer id | 
  **locationId** | **string**|  | 
  **managerId** | **string**|  | 

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomerPrototypeDeleteNodeLocked**
> interface{} CustomerPrototypeDeleteNodeLocked(ctx, id, nodeId)
Delete a node model based on its id.

<div><strong>204</strong>: The node was successfully deleted.</div> <div><strong>401</strong>: Authorization Required.</div> <div><strong>404</strong>: Node or customer not found.</div> <div><strong>422</strong>: Node deletion could not be completed.</div> <div><strong>500</strong>: Internal server error.</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Customer id | 
  **nodeId** | **string**|  | 

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomerPrototypeDeleteNodePersistentConfigs**
> CustomerPrototypeDeleteNodePersistentConfigs(ctx, id, locationId, nodeId, optional)
Delete persistent data/configs from node in runtime.

<div><strong>204</strong>: Success.</div> <div><strong>400</strong>: Required fields missing.</div> <div><strong>401</strong>: Authorization required or customer id not found.</div> <div><strong>404</strong>: Location or Node, does not exist.</div> <div><strong>422</strong>: Required fields are not valid.</div> <div><strong>500</strong>: Internal server error</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Customer id | 
  **locationId** | **string**| location id | 
  **nodeId** | **string**| node id | 
 **optional** | ***CustomerApiCustomerPrototypeDeleteNodePersistentConfigsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CustomerApiCustomerPrototypeDeleteNodePersistentConfigsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **deleteAllPersistentConfigs** | **optional.Bool**| whether all persistent config data or just GW-offline data will be deleted | [default to false]

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomerPrototypeDeletePerson**
> CustomerPrototypeDeletePerson(ctx, id, locationId, personId, optional)
Delete a Person for a location ID.

<div><strong>204</strong>: Success.</div> <div><strong>401</strong>: Authorization required or customer id not found</div> <div><strong>404</strong>: Location id or Person id does not exist and is not known to Plume</div> <div><strong>500</strong>: Internal server error</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Customer id | 
  **locationId** | **string**|  | 
  **personId** | **string**|  | 
 **optional** | ***CustomerApiCustomerPrototypeDeletePersonOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CustomerApiCustomerPrototypeDeletePersonOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **blockUnassignedDevices** | **optional.Bool**| block any devices previously assigned to Person (false by default) | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomerPrototypeDeletePersonAllFreeze**
> CustomerPrototypeDeletePersonAllFreeze(ctx, id, locationId, personId)
Delete a person to be frozen for a Location ID.

<div><strong>204</strong>: Success, updated.</div> <div><strong>401</strong>: Authorization required or customer id not found.</div> <div><strong>404</strong>: Location id, does not exist.</div> <div><strong>500</strong>: Internal server error</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Customer id | 
  **locationId** | **string**|  | 
  **personId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomerPrototypeDeletePersonFreeze**
> CustomerPrototypeDeletePersonFreeze(ctx, id, locationId, personId, freezeTemplateId)
Delete a person to be frozen for a Location ID.

<div><strong>204</strong>: Success, updated.</div> <div><strong>401</strong>: Authorization required or customer id not found.</div> <div><strong>404</strong>: Location id, does not exist.</div> <div><strong>500</strong>: Internal server error</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Customer id | 
  **locationId** | **string**|  | 
  **personId** | **string**|  | 
  **freezeTemplateId** | **string**| Valid templates are &#39;untilMidinight&#39;, &#39;schoolNights&#39;, etc. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomerPrototypeDeletePersonFreezeAutoExpire**
> CustomerPrototypeDeletePersonFreezeAutoExpire(ctx, id, locationId, personId)
Delete all devices from a person to be frozen for a Location ID.

<div><strong>204</strong>: Success, updated.</div> <div><strong>401</strong>: Authorization required or customer id not found.</div> <div><strong>404</strong>: Location id, does not exist.</div> <div><strong>500</strong>: Internal server error</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Customer id | 
  **locationId** | **string**|  | 
  **personId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomerPrototypeDeletePersonFreezeForever**
> CustomerPrototypeDeletePersonFreezeForever(ctx, id, locationId, personId)
Delete a person forever freeze for a Location ID.

<div><strong>204</strong>: Success, updated.</div> <div><strong>401</strong>: Authorization required or customer id not found.</div> <div><strong>404</strong>: Location id, does not exist.</div> <div><strong>500</strong>: Internal server error</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Customer id | 
  **locationId** | **string**|  | 
  **personId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomerPrototypeDeletePersonFreezeSuspend**
> CustomerPrototypeDeletePersonFreezeSuspend(ctx, id, locationId, personId)
Delete person suspend for a Location ID.

<div><strong>204</strong>: Success, updated.</div> <div><strong>401</strong>: Authorization required or customer id not found.</div> <div><strong>404</strong>: Location id, does not exist.</div> <div><strong>500</strong>: Internal server error</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Customer id | 
  **locationId** | **string**|  | 
  **personId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomerPrototypeDeletePersonProfile**
> CustomerPrototypeDeletePersonProfile(ctx, id, locationId, personId)
Delete a Person's Profile for a location ID.

<div><strong>204</strong>: Success.</div> <div><strong>401</strong>: Authorization required or customer id not found</div> <div><strong>404</strong>: Location id or Person id does not exist and is not known to Plume</div> <div><strong>500</strong>: Internal server error</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Customer id | 
  **locationId** | **string**|  | 
  **personId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomerPrototypeDeletePortForward**
> []PortForward CustomerPrototypeDeletePortForward(ctx, id, locationId, mac, externalPort)
Delete an existing Port Forwarding entry for an existing DHCP IP reservation tied to a MAC address at a Location ID.

<div><strong>200</strong>: Success, returns list of remaining port forwards.</div> <div><strong>401</strong>: Authorization required or customer id not found</div> <div><strong>404</strong>: NetworkConfiguration, DhcpReservation or PortForward does not exist.</div> <div><strong>422</strong>: mac does not exist, or is invalid, or externalPort is empty.</div> <div><strong>500</strong>: internal server error</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Customer id | 
  **locationId** | **string**|  | 
  **mac** | **string**|  | 
  **externalPort** | **string**|  | 

### Return type

[**[]PortForward**](PortForward.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomerPrototypeDeleteRemoteConnectionsAllow**
> interface{} CustomerPrototypeDeleteRemoteConnectionsAllow(ctx, id, locationId, mac, ipaddr)
Delete a Remote Connection Allow IpAddress/ttl for the given device and Location ID.

<div><strong>200</strong>: Success.</div> <div><strong>401</strong>: Authorization required or customer id not found</div> <div><strong>404</strong>: Location id or Device mac does not exist and is not known to Plume</div> <div><strong>500</strong>: Internal server error</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Customer id | 
  **locationId** | **string**|  | 
  **mac** | **string**|  | 
  **ipaddr** | **string**| ipaddress | 

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomerPrototypeDeleteRemoteConnectionsAllowAll**
> interface{} CustomerPrototypeDeleteRemoteConnectionsAllowAll(ctx, id, locationId, mac)
Delete a Remote Connection Allow All for the given device and Location ID.

<div><strong>200</strong>: Success.</div> <div><strong>401</strong>: Authorization required or customer id not found</div> <div><strong>404</strong>: Location id or Device mac does not exist and is not known to Plume</div> <div><strong>500</strong>: Internal server error</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Customer id | 
  **locationId** | **string**|  | 
  **mac** | **string**|  | 

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomerPrototypeDeleteRoom**
> CustomerPrototypeDeleteRoom(ctx, id, locationId, roomId)
Delete a Room for a location ID.

<div><strong>204</strong>: Success.</div> <div><strong>401</strong>: Authorization required or customer id not found</div> <div><strong>404</strong>: Location id or Room id does not exist and is not known to Plume</div> <div><strong>500</strong>: Internal server error</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Customer id | 
  **locationId** | **string**|  | 
  **roomId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomerPrototypeDeleteWifiAccessZone**
> []WifiAccessZone CustomerPrototypeDeleteWifiAccessZone(ctx, id, locationId, accessZoneId)
Delete a WiFi Access Zone

<div><strong>200</strong>: Success, remaining access zones returned</div> <div><strong>400</strong>: Required fields missing</div> <div><strong>401</strong>: Authorization required or customer id not found</div> <div><strong>404</strong>: Customer id, location id, or WifiNetwork does not exist</div> <div><strong>405</strong>: Cannot delete a read-only access zone</div> <div><strong>500</strong>: Internal server error</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Customer id | 
  **locationId** | **string**|  | 
  **accessZoneId** | **float64**| access zone id | 

### Return type

[**[]WifiAccessZone**](WifiAccessZone.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomerPrototypeDeleteWifiKey**
> []WifiNetworkKey CustomerPrototypeDeleteWifiKey(ctx, id, locationId, accessZone, keyId)
Delete a WiFi Password

<div><strong>200</strong>: Success, all passwords returned</div> <div><strong>400</strong>: Required fields missing</div> <div><strong>401</strong>: Authorization required or customer id not found</div> <div><strong>404</strong>: Customer id, location id, or WifiNetwork does not exist</div> <div><strong>405</strong>: Cannot delete a read-only key</div> <div><strong>500</strong>: Internal server error</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Customer id | 
  **locationId** | **string**|  | 
  **accessZone** | **string**| home | guests | internetAccessOnly | 
  **keyId** | **float64**| Unique password id: 0-9 | 

### Return type

[**[]WifiNetworkKey**](WifiNetworkKey.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomerPrototypeDestroyMigration**
> CustomerPrototypeDestroyMigration(ctx, id)
Deletes _migration of this model.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Customer id | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomerPrototypeDisableLogin**
> CustomerPrototypeDisableLogin(ctx, id, optional)
Disable customer from logging in until their account is reactivated.

<div><strong>204</strong>: Customer has been disabled.</div> <div><strong>401</strong>: Authorization required.</div> <div><strong>404</strong>: Customer does not exist.</div> <div><strong>500</strong>: Internal server error.</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Customer id | 
 **optional** | ***CustomerApiCustomerPrototypeDisableLoginOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CustomerApiCustomerPrototypeDisableLoginOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **triggerReset** | **optional.Bool**|  | [default to false]

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomerPrototypeEnableDeviceTypeSniffing**
> CustomerPrototypeEnableDeviceTypeSniffing(ctx, id, locationId, mac)
Re-enables deviceType sniffing for a particular device.

<div><strong>204</strong>: Success, your new info looks good.</div> <div><strong>401</strong>: Authorization required or customer id not found.</div> <div><strong>404</strong>: location id, does not exist.</div> <div><strong>404</strong>: No device found with provided mac address</div> <div><strong>500</strong>: internal server error</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Customer id | 
  **locationId** | **string**|  | 
  **mac** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomerPrototypeEnableLogin**
> CustomerPrototypeEnableLogin(ctx, id)
Enable customer log in, after it has been disabled.

<div><strong>204</strong>: Customer has been enabled.</div> <div><strong>401</strong>: Authorization required.</div> <div><strong>404</strong>: Customer does not exist.</div> <div><strong>500</strong>: Internal server error.</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Customer id | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomerPrototypeFactoryReset**
> CustomerPrototypeFactoryReset(ctx, id, locationId, optional)
Reset specified location settings to default, while keeping claimed nodes intact. Some of the flags can cause a node to be reeboted.

<div><strong>204</strong>: Success, a job well done.</div> <div><strong>401</strong>: Authorization required </div> <div><strong>404</strong>: location id not found or nodeId missing from URL <div><strong>500</strong>: internal server error</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Customer id | 
  **locationId** | **string**|  | 
 **optional** | ***CustomerApiCustomerPrototypeFactoryResetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CustomerApiCustomerPrototypeFactoryResetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **persons** | **optional.Bool**| Whether or not to delete person information | [default to false]
 **onboardingCheckpoints** | **optional.Bool**| Whether or not to reset onboarding checkpoints | [default to false]
 **devices** | **optional.Bool**| Whether or not to delete devices related information | [default to false]
 **networkConfiguration** | **optional.Bool**| Whether or not to reset network configuration (triggers node reboot) | [default to false]
 **wifiNetwork** | **optional.Bool**| Whether or not to reset wifi network (triggers node reboot) | [default to false]
 **deviceFreeze** | **optional.Bool**| Whether or not to reset device freeze templates | [default to false]
 **deviceNicknames** | **optional.Bool**| Whether or not to reset device nicknames | [default to false]
 **managers** | **optional.Bool**| Whether or not to reset managers of the location | [default to false]
 **wanConfiguration** | **optional.Bool**| Whether or not to reset wanConfiguration | [default to false]

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomerPrototypeFindLocationById**
> Location CustomerPrototypeFindLocationById(ctx, id, locationId, optional)
Get a Location's combined State and Config by LocationId.

<div><strong>200</strong>: Success, full object returned.</div> <div><strong>401</strong>: Authorization required.</div> <div><strong>404</strong>: LocationId not found.</div> <div><strong>500</strong>: Internal server error.</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Customer id | 
  **locationId** | **string**|  | 
 **optional** | ***CustomerApiCustomerPrototypeFindLocationByIdOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CustomerApiCustomerPrototypeFindLocationByIdOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **include** | **optional.String**| CSV value of objects to add to the response: summary (is the only option for now) | 

### Return type

[**Location**](Location.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomerPrototypeGetAccessTokenForManagedLocation**
> interface{} CustomerPrototypeGetAccessTokenForManagedLocation(ctx, id, locationId)
Get an access token for a location where you are assigned as a manager

<div><strong>200</strong>: Success.</div> <div><strong>400</strong>: Required fields missing or field type is incorrect.</div> <div><strong>404</strong>: Location does not exist.</div> <div><strong>422</strong>: Invalid email, name, access type or manager is already assigned to this location </div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Customer id | 
  **locationId** | **string**|  | 

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomerPrototypeGetAlerts**
> interface{} CustomerPrototypeGetAlerts(ctx, id, locationId)
Retrieve active alerts for this location.

<div><strong>200</strong>: Success, an array of Nodes and an array of Devices returned.</div> <div><strong>404</strong>: customer id or location id does not exist and is not known to Plume</div> <div><strong>500</strong>: internal server error</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Customer id | 
  **locationId** | **string**|  | 

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomerPrototypeGetAppEngagementTimer**
> interface{} CustomerPrototypeGetAppEngagementTimer(ctx, id, locationId)
Get information about app engagement timer details for a location

<div><strong>204</strong>: Success.</div> <div><strong>404</strong>: Location does not exist.</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Customer id | 
  **locationId** | **string**|  | 

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomerPrototypeGetAppIdInfoCaptivePortalNetwork**
> NetworkConfig CustomerPrototypeGetAppIdInfoCaptivePortalNetwork(ctx, id, locationId)
Get the AppId info for the given location.

<div><strong>200</strong>: Success, appId info returned.</div> <div><strong>401</strong>: Authorization required or customer id not found</div> <div><strong>404</strong>: location id or url does not exist.</div> <div><strong>500</strong>: internal server error</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Customer id | 
  **locationId** | **string**|  | 

### Return type

[**NetworkConfig**](NetworkConfig.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomerPrototypeGetAppPrioritizationLocationConfig**
> interface{} CustomerPrototypeGetAppPrioritizationLocationConfig(ctx, id, locationId)
Get status for app prioritization.

<div><strong>200</strong>: Success.</div> <div><strong>401</strong>: Authorization required or customer id not found</div> <div><strong>404</strong>: Location id or WifiNetwork does not exist and is not known to Plume</div> <div><strong>500</strong>: Internal server error</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Customer id | 
  **locationId** | **string**|  | 

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomerPrototypeGetAppPrioritizationMonitoring**
> interface{} CustomerPrototypeGetAppPrioritizationMonitoring(ctx, id, locationId, startTime, endTime, optional)
Get monitoring metrics for app prioritization.

<div><strong>200</strong>: Success.</div> <div><strong>401</strong>: Authorization required or customer id not found</div> <div><strong>404</strong>: Location id or WifiNetwork does not exist and is not known to Plume</div> <div><strong>500</strong>: Internal server error</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Customer id | 
  **locationId** | **string**|  | 
  **startTime** | **time.Time**| format yyyy-mm-ddThh:MM:ss.nnnZ, 24 hours time specified in UTC | 
  **endTime** | **time.Time**| format yyyy-mm-ddThh:MM:ss.nnnZ, 24 hours time specified in UTC | 
 **optional** | ***CustomerApiCustomerPrototypeGetAppPrioritizationMonitoringOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CustomerApiCustomerPrototypeGetAppPrioritizationMonitoringOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **granularity** | **optional.String**| any of the values - total/15 minutes/1 hour/1 day | 
 **macs** | **optional.String**| array of macs[] | 
 **trafficClasses** | **optional.String**| array of trafficClasses[] | 
 **sortOrder** | **optional.String**| TxBytes\&quot;|| \&quot;RxBytes | 
 **limit** | **optional.Float64**| Maximum number of devices to return. | 

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomerPrototypeGetAppPrioritizationTemplateConfig**
> interface{} CustomerPrototypeGetAppPrioritizationTemplateConfig(ctx, id, locationId)
Get AppPrioritization template configs

<div><strong>200</strong>: Success.</div> <div><strong>401</strong>: Authorization required or customer id not found</div> <div><strong>404</strong>: Location id or WifiNetwork does not exist and is not known to Plume</div> <div><strong>500</strong>: Internal server error</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Customer id | 
  **locationId** | **string**|  | 

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomerPrototypeGetAppQoeStatsByTrafficClass**
> interface{} CustomerPrototypeGetAppQoeStatsByTrafficClass(ctx, id, locationId, timePeriod, optional)
Get App QoE metrics by traffic classes / devices / apps.

<div><strong>200</strong>: Success.</div> <div><strong>401</strong>: Authorization required or customer id not found</div> <div><strong>404</strong>: Location id or WifiNetwork does not exist and is not known to Plume</div> <div><strong>500</strong>: Internal server error</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Customer id | 
  **locationId** | **string**|  | 
  **timePeriod** | **string**| Any of \&quot;last24Hours\&quot;,\&quot;last7Days\&quot;,\&quot;last30Days\&quot; | 
 **optional** | ***CustomerApiCustomerPrototypeGetAppQoeStatsByTrafficClassOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CustomerApiCustomerPrototypeGetAppQoeStatsByTrafficClassOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **includeApps** | **optional.Bool**| Default false, to include app stats in the response | 
 **trafficClassNames** | **optional.String**| array of traffic classes - default list - av_streaming, gaming, video_conferencing | 

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomerPrototypeGetAppQoeTrafficClassMetricsGetCustomersidLocationslocationIdAppqoeTrafficClassStats**
> interface{} CustomerPrototypeGetAppQoeTrafficClassMetricsGetCustomersidLocationslocationIdAppqoeTrafficClassStats(ctx, id, locationId, granularity, startTime, endTime, optional)
Get App QoE metrics for traffic classes.

<div><strong>200</strong>: Success.</div> <div><strong>401</strong>: Authorization required or customer id not found</div> <div><strong>404</strong>: Location id or WifiNetwork does not exist and is not known to Plume</div> <div><strong>500</strong>: Internal server error</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Customer id | 
  **locationId** | **string**|  | 
  **granularity** | **string**| any of the values - total/1 minute/15 minutes/1 hour/1 day | 
  **startTime** | **time.Time**| format yyyy-mm-ddThh:MM:ss.nnnZ, 24 hours time specified in UTC | 
  **endTime** | **time.Time**| format yyyy-mm-ddThh:MM:ss.nnnZ, 24 hours time specified in UTC | 
 **optional** | ***CustomerApiCustomerPrototypeGetAppQoeTrafficClassMetricsGetCustomersidLocationslocationIdAppqoeTrafficClassStatsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CustomerApiCustomerPrototypeGetAppQoeTrafficClassMetricsGetCustomersidLocationslocationIdAppqoeTrafficClassStatsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





 **macs** | **optional.String**| array of macs[] | 
 **trafficClasses** | **optional.String**| array of trafficClasses | 
 **limit** | **optional.Float64**| Maximum number of devices to return. | 

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomerPrototypeGetAppQoeTrafficClassMetricsPostCustomersidLocationslocationIdAppqoeTrafficClassStats**
> interface{} CustomerPrototypeGetAppQoeTrafficClassMetricsPostCustomersidLocationslocationIdAppqoeTrafficClassStats(ctx, id, locationId, granularity, startTime, endTime, optional)
Get App QoE metrics for traffic classes.

<div><strong>200</strong>: Success.</div> <div><strong>401</strong>: Authorization required or customer id not found</div> <div><strong>404</strong>: Location id or WifiNetwork does not exist and is not known to Plume</div> <div><strong>500</strong>: Internal server error</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Customer id | 
  **locationId** | **string**|  | 
  **granularity** | **string**| any of the values - total/1 minute/15 minutes/1 hour/1 day | 
  **startTime** | **time.Time**| format yyyy-mm-ddThh:MM:ss.nnnZ, 24 hours time specified in UTC | 
  **endTime** | **time.Time**| format yyyy-mm-ddThh:MM:ss.nnnZ, 24 hours time specified in UTC | 
 **optional** | ***CustomerApiCustomerPrototypeGetAppQoeTrafficClassMetricsPostCustomersidLocationslocationIdAppqoeTrafficClassStatsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CustomerApiCustomerPrototypeGetAppQoeTrafficClassMetricsPostCustomersidLocationslocationIdAppqoeTrafficClassStatsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





 **macs** | **optional.String**| array of macs | 
 **trafficClasses** | **optional.String**| array of trafficClasses | 
 **limit** | **optional.Float64**| Maximum number of devices to return. | 

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomerPrototypeGetAppTimeIpFlows**
> interface{} CustomerPrototypeGetAppTimeIpFlows(ctx, id, locationId)
Get IP flows config

<div><strong>200</strong>: Success.</div> <div><strong>400</strong>: Required fields missing or field type is incorrect.</div> <div><strong>401</strong>: Authorization required or customer id not found</div> <div><strong>404</strong>: Location id or device does not exist and is not known to Plume</div> <div><strong>500</strong>: Internal server error</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Customer id | 
  **locationId** | **string**|  | 

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomerPrototypeGetAuditTrailForCustomer**
> XAny CustomerPrototypeGetAuditTrailForCustomer(ctx, id)
Get audit trail for a customer.

<div><strong>200</strong>: Ok.</div> <div><strong>401</strong>: Authorization required.</div> <div><strong>404</strong>: Customer id, does not exist.</div> <div><strong>500</strong>: Internal server error.</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Customer id | 

### Return type

[**XAny**](x-any.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomerPrototypeGetAuditTrailForLocation**
> XAny CustomerPrototypeGetAuditTrailForLocation(ctx, id, locationId)
Get audit trail for location.

<div><strong>200</strong>: Ok.</div> <div><strong>401</strong>: Authorization required.</div> <div><strong>404</strong>: Location id, does not exist.</div> <div><strong>500</strong>: Internal server error.</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Customer id | 
  **locationId** | **string**|  | 

### Return type

[**XAny**](x-any.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomerPrototypeGetAuthorizations**
> Authorizations CustomerPrototypeGetAuthorizations(ctx, id, locationId)
Get the number of authorized leaf pods for a Location ID.

<div>Number of leaf pods that are authorized to be claimed and be a part of the Plume network</div> <div><strong>200</strong>: Success, numPodsAuthorized returned.</div> <div><strong>401</strong>: Authorization required or customer id not found.</div> <div><strong>404</strong>: Location id, does not exist.</div> <div><strong>500</strong>: Internal server error</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Customer id | 
  **locationId** | **string**|  | 

### Return type

[**Authorizations**](Authorizations.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomerPrototypeGetBackhaul**
> interface{} CustomerPrototypeGetBackhaul(ctx, id, locationId)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Customer id | 
  **locationId** | **string**|  | 

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomerPrototypeGetCampaignCaptivePortalNetwork**
> NetworkConfig CustomerPrototypeGetCampaignCaptivePortalNetwork(ctx, id, locationId, networkId)
Get the Captive Portal campaign for a given Location ID/NetworkId.

<div><strong>200</strong>: Success, campaign returned.</div> <div><strong>401</strong>: Authorization required or customer id not found</div> <div><strong>404</strong>: location id or CaptivePortal Network does not exist.</div> <div><strong>500</strong>: internal server error</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Customer id | 
  **locationId** | **string**|  | 
  **networkId** | **string**|  | 

### Return type

[**NetworkConfig**](NetworkConfig.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomerPrototypeGetCaptivePortalAuthorizedClients**
> interface{} CustomerPrototypeGetCaptivePortalAuthorizedClients(ctx, id, locationId, networkId)
Get CaptivePortal authorized clients

<div><strong>200</strong>: Success.</div> <div><strong>401</strong>: Authorization required or customer id not found</div> <div><strong>404</strong>: Location id/NetworkId does not exist and is not known to Plume</div> <div><strong>500</strong>: Internal server error</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Customer id | 
  **locationId** | **string**|  | 
  **networkId** | **string**|  | 

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomerPrototypeGetCaptivePortalGuestEmailCollectionInfo**
> InlineResponse2008 CustomerPrototypeGetCaptivePortalGuestEmailCollectionInfo(ctx, id, locationId, networkId, optional)
Fetch the Captive Portal Network guest info download availability for the given network.

<div><strong>200</strong>: Success, CaptivePortal Networks guest info download availability returned.</div> <div><strong>401</strong>: Authorization required or customer id not found</div> <div><strong>404</strong>: location id or secondary networks does not exist.</div> <div><strong>500</strong>: internal server error</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Customer id | 
  **locationId** | **string**|  | 
  **networkId** | **string**|  | 
 **optional** | ***CustomerApiCustomerPrototypeGetCaptivePortalGuestEmailCollectionInfoOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CustomerApiCustomerPrototypeGetCaptivePortalGuestEmailCollectionInfoOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **duration** | **optional.Float64**| number of days for how far back in history for data | 
 **limit** | **optional.Float64**| limit how many emails we wish to return | 

### Return type

[**InlineResponse2008**](inline_response_200_8.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomerPrototypeGetCaptivePortalGuests**
> []interface{} CustomerPrototypeGetCaptivePortalGuests(ctx, id, locationId, networkId, optional)
Fetch the list of Guests which were logged into the given captivePortal network during the current day.

<div><strong>200</strong>: Success, CaptivePortal Networks returned.</div> <div><strong>401</strong>: Authorization required or customer id not found</div> <div><strong>404</strong>: location id or secondary networks does not exist.</div> <div><strong>500</strong>: internal server error</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Customer id | 
  **locationId** | **string**|  | 
  **networkId** | **string**|  | 
 **optional** | ***CustomerApiCustomerPrototypeGetCaptivePortalGuestsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CustomerApiCustomerPrototypeGetCaptivePortalGuestsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **orderBy** | **optional.String**| Order by: &lt;connectionTime&gt; | [default to connectionTime]
 **limit** | **optional.Float64**|  | [default to 20]

### Return type

**[]interface{}**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomerPrototypeGetCaptivePortalNetworks**
> NetworkConfig CustomerPrototypeGetCaptivePortalNetworks(ctx, id, locationId)
Get the Captive Portal configs for a given Location ID.

<div><strong>200</strong>: Success, CaptivePortal Networks returned.</div> <div><strong>401</strong>: Authorization required or customer id not found</div> <div><strong>404</strong>: location id or secondary networks does not exist.</div> <div><strong>500</strong>: internal server error</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Customer id | 
  **locationId** | **string**|  | 

### Return type

[**NetworkConfig**](NetworkConfig.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomerPrototypeGetCaptivePortalSendDetails**
> *os.File CustomerPrototypeGetCaptivePortalSendDetails(ctx, id, locationId, networkId)
Download Captive Portal Guest details for a given Location ID/NetworkId.

<div><strong>204</strong>: Success, updated.</div> <div><strong>401</strong>: Authorization required or customer id not found</div> <div><strong>404</strong>: location id or CaptivePortal Network does not exist.</div> <div><strong>500</strong>: internal server error</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Customer id | 
  **locationId** | **string**|  | 
  **networkId** | **string**|  | 

### Return type

[***os.File**](*os.File.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomerPrototypeGetCaptivePortalSendDetailsDirect**
> *os.File CustomerPrototypeGetCaptivePortalSendDetailsDirect(ctx, id, locationId, networkId, optional)
Download Captive Portal Guest details for a given Location ID/NetworkId without accessing Amazon S3.

<div><strong>204</strong>: Success, updated.</div> <div><strong>401</strong>: Authorization required or customer id not found</div> <div><strong>404</strong>: location id or CaptivePortal Network does not exist.</div> <div><strong>500</strong>: internal server error</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Customer id | 
  **locationId** | **string**|  | 
  **networkId** | **string**|  | 
 **optional** | ***CustomerApiCustomerPrototypeGetCaptivePortalSendDetailsDirectOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CustomerApiCustomerPrototypeGetCaptivePortalSendDetailsDirectOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **duration** | **optional.Float64**|  | 
 **limit** | **optional.Float64**|  | 

### Return type

[***os.File**](*os.File.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomerPrototypeGetCommands**
> []interface{} CustomerPrototypeGetCommands(ctx, id, locationId)
Returns list of linked command accounts for the location

<div><strong>200</strong>: Success, return the  list of linked command accounts.</div> <div><strong>401</strong>: Authorization required</div> <div><strong>404</strong>: location id does not exist.</div> <div><strong>500</strong>: internal server error</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Customer id | 
  **locationId** | **string**|  | 

### Return type

**[]interface{}**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomerPrototypeGetCompanyInfoCaptivePortalNetwork**
> NetworkConfig CustomerPrototypeGetCompanyInfoCaptivePortalNetwork(ctx, id, locationId, optional)
Get the companyInfo for the given url (domain).

<div><strong>200</strong>: Success, company info returned.</div> <div><strong>401</strong>: Authorization required or customer id not found</div> <div><strong>404</strong>: location id or url does not exist.</div> <div><strong>500</strong>: internal server error</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Customer id | 
  **locationId** | **string**|  | 
 **optional** | ***CustomerApiCustomerPrototypeGetCompanyInfoCaptivePortalNetworkOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CustomerApiCustomerPrototypeGetCompanyInfoCaptivePortalNetworkOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **url** | **optional.String**|  | 
 **urlType** | **optional.String**| only &#39;domain&#39; currently supported | 

### Return type

[**NetworkConfig**](NetworkConfig.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomerPrototypeGetControlMode**
> interface{} CustomerPrototypeGetControlMode(ctx, id, locationId)
Get control mode for a Location ID.

<div><strong>200</strong>: Success.</div> <div><strong>401</strong>: Authorization required or customer id not found.</div> <div><strong>404</strong>: Location id, does not exist.</div> <div><strong>500</strong>: Internal server error</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Customer id | 
  **locationId** | **string**|  | 

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomerPrototypeGetCustomerNodeById**
> NodeResponse CustomerPrototypeGetCustomerNodeById(ctx, id, nodeId)
Returns a single Node for a Customer ID.

<div><strong>200</strong>: Success, node returned with locationId field.</div> <div><strong>401</strong>: Authorization required or customer id not found.</div> <div><strong>404</strong>: customer id or location id does not exist. Or, nodeId not claimed to this account.</div> <div><strong>500</strong>: internal server error</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Customer id | 
  **nodeId** | **string**| id of node | 

### Return type

[**NodeResponse**](NodeResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomerPrototypeGetCustomerSupportConfigurations**
> XAny CustomerPrototypeGetCustomerSupportConfigurations(ctx, id, locationId)
Returns partner customer support configuration.

<div><strong>200</strong>: Success.</div> <div><strong>404</strong>: customer id or location id does not exist.</div> <div><strong>500</strong>: internal server error</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Customer id | 
  **locationId** | **string**|  | 

### Return type

[**XAny**](x-any.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomerPrototypeGetDashboard**
> XAny CustomerPrototypeGetDashboard(ctx, id, locationId, optional)
Daily/Weekly/Monthly device usage summary report based on location

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Customer id | 
  **locationId** | **string**|  | 
 **optional** | ***CustomerApiCustomerPrototypeGetDashboardOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CustomerApiCustomerPrototypeGetDashboardOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **macs** | **optional.String**| mac list of all devices in the location | 

### Return type

[**XAny**](x-any.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomerPrototypeGetDefaultNetworkAppTimeAppsDataUsage**
> []interface{} CustomerPrototypeGetDefaultNetworkAppTimeAppsDataUsage(ctx, id, locationId, timePeriod, optional)
Fetch the AppTime Apps Data Usage Stats for captivePortal network.

<div><strong>200</strong>: Success, AppTime Stats returned.</div> <div><strong>401</strong>: Authorization required or customer id not found</div> <div><strong>404</strong>: location id or secondary networks does not exist.</div> <div><strong>500</strong>: internal server error</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Customer id | 
  **locationId** | **string**|  | 
  **timePeriod** | **string**| Any of \&quot;weekly\&quot;,\&quot;dailyToday\&quot;,\&quot;dailyYesterday\&quot;,\&quot;daily2DaysAgo\&quot;,\&quot;daily3DaysAgo\&quot;,\&quot;daily4DaysAgo\&quot;,\&quot;daily5DaysAgo\&quot;,\&quot;daily6DaysAgo\&quot;,\&quot;last30Days\&quot;,\&quot;last12Months\&quot; | 
 **optional** | ***CustomerApiCustomerPrototypeGetDefaultNetworkAppTimeAppsDataUsageOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CustomerApiCustomerPrototypeGetDefaultNetworkAppTimeAppsDataUsageOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **limit** | **optional.Float64**| Maximum number of apps to return. Defaults to 20 | 
 **grouping** | **optional.String**| typing of Grouping for the purposes of applying the limit. Can be: &#39;perTimeSlot&#39; ONLY | [default to perTimeSlot]

### Return type

**[]interface{}**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomerPrototypeGetDefaultNetworkAppTimeAppsOnlineTime**
> []interface{} CustomerPrototypeGetDefaultNetworkAppTimeAppsOnlineTime(ctx, id, locationId, timePeriod, optional)
Fetch the AppTime Apps Online Time Stats for default network.

<div><strong>200</strong>: Success, AppTime Stats returned.</div> <div><strong>401</strong>: Authorization required or customer id not found</div> <div><strong>404</strong>: location id or secondary networks does not exist.</div> <div><strong>500</strong>: internal server error</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Customer id | 
  **locationId** | **string**|  | 
  **timePeriod** | **string**| Any of \&quot;weekly\&quot;,\&quot;dailyToday\&quot;,\&quot;dailyYesterday\&quot;,\&quot;daily2DaysAgo\&quot;,\&quot;daily3DaysAgo\&quot;,\&quot;daily4DaysAgo\&quot;,\&quot;daily5DaysAgo\&quot;,\&quot;daily6DaysAgo\&quot;,\&quot;last30Days\&quot;,\&quot;last12Months\&quot; | 
 **optional** | ***CustomerApiCustomerPrototypeGetDefaultNetworkAppTimeAppsOnlineTimeOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CustomerApiCustomerPrototypeGetDefaultNetworkAppTimeAppsOnlineTimeOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **limit** | **optional.Float64**| Maximum number of apps to return. Defaults to 20 | 
 **grouping** | **optional.String**| typing of Grouping for the purposes of applying the limit. Can be: &#39;perTimeSlot&#39; ONLY | [default to perTimeSlot]

### Return type

**[]interface{}**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomerPrototypeGetDefaultNetworkAppTimeCategoriesDataUsage**
> []interface{} CustomerPrototypeGetDefaultNetworkAppTimeCategoriesDataUsage(ctx, id, locationId, timePeriod, optional)
Fetch the AppTime Categories Data Usage Stats for captivePortal network.

<div><strong>200</strong>: Success, AppTime Stats returned.</div> <div><strong>401</strong>: Authorization required or customer id not found</div> <div><strong>404</strong>: location id or secondary networks does not exist.</div> <div><strong>500</strong>: internal server error</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Customer id | 
  **locationId** | **string**|  | 
  **timePeriod** | **string**| Any of \&quot;weekly\&quot;,\&quot;dailyToday\&quot;,\&quot;dailyYesterday\&quot;,\&quot;daily2DaysAgo\&quot;,\&quot;daily3DaysAgo\&quot;,\&quot;daily4DaysAgo\&quot;,\&quot;daily5DaysAgo\&quot;,\&quot;daily6DaysAgo\&quot;,\&quot;last30Days\&quot;,\&quot;last12Months\&quot; | 
 **optional** | ***CustomerApiCustomerPrototypeGetDefaultNetworkAppTimeCategoriesDataUsageOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CustomerApiCustomerPrototypeGetDefaultNetworkAppTimeCategoriesDataUsageOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **limit** | **optional.Float64**| Maximum number of categories to return. Defaults to 20 | 
 **grouping** | **optional.String**| typing of Grouping for the purposes of applying the limit. Can be: &#39;overall&#39;|&#39;perTimeSlot&#39; | [default to perTimeSlot]

### Return type

**[]interface{}**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomerPrototypeGetDefaultNetworkAppTimeCategoriesOnlineTime**
> []interface{} CustomerPrototypeGetDefaultNetworkAppTimeCategoriesOnlineTime(ctx, id, locationId, timePeriod, optional)
Fetch the AppTime Categories Online Time Stats for captivePortal network.

<div><strong>200</strong>: Success, AppTime Stats returned.</div> <div><strong>401</strong>: Authorization required or customer id not found</div> <div><strong>404</strong>: location id or secondary networks does not exist.</div> <div><strong>500</strong>: internal server error</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Customer id | 
  **locationId** | **string**|  | 
  **timePeriod** | **string**| Any of \&quot;weekly\&quot;,\&quot;dailyToday\&quot;,\&quot;dailyYesterday\&quot;,\&quot;daily2DaysAgo\&quot;,\&quot;daily3DaysAgo\&quot;,\&quot;daily4DaysAgo\&quot;,\&quot;daily5DaysAgo\&quot;,\&quot;daily6DaysAgo\&quot;,\&quot;last30Days\&quot;,\&quot;last12Months\&quot; | 
 **optional** | ***CustomerApiCustomerPrototypeGetDefaultNetworkAppTimeCategoriesOnlineTimeOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CustomerApiCustomerPrototypeGetDefaultNetworkAppTimeCategoriesOnlineTimeOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **limit** | **optional.Float64**| Maximum number of categories to return. Defaults to 20 | 
 **grouping** | **optional.String**| typing of Grouping for the purposes of applying the limit. Can be: &#39;overall&#39;|&#39;perTimeSlot&#39; | [default to perTimeSlot]

### Return type

**[]interface{}**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomerPrototypeGetDeviceAlarms**
> interface{} CustomerPrototypeGetDeviceAlarms(ctx, id, locationId, mac, optional)
Device alarm history graph array for a particular MAC address.

<div><strong>200</strong>: Success.</div> <div><strong>401</strong>: Authorization required.</div> <div><strong>404</strong>: Location id does not exist.</div> <div><strong>500</strong>: Internal server error</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Customer id | 
  **locationId** | **string**|  | 
  **mac** | **string**| mac id of device | 
 **optional** | ***CustomerApiCustomerPrototypeGetDeviceAlarmsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CustomerApiCustomerPrototypeGetDeviceAlarmsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **coverageAlarmThreshold** | **optional.String**| a coverage alarm will be returned (value&#x3D;1) when rssi_alarm_penalty_count &gt;&#x3D; this value | [default to 1]
 **granularity** | **optional.String**| days/hours/minutes | [default to days]
 **limit** | **optional.Float64**| X # of days/hours/minutes | [default to 7]
 **start** | **optional.Float64**| number of milliseconds elapsed since 1 January 1970 00:00:00 UTC. Defaults to now - (limit * granularity) | 

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomerPrototypeGetDeviceAppTimeAppsDataUsage**
> []interface{} CustomerPrototypeGetDeviceAppTimeAppsDataUsage(ctx, id, locationId, mac, timePeriod, optional)
Fetch the AppTime Apps Data Usage Stats for a Device.

<div><strong>200</strong>: Success, AppTime Stats returned.</div> <div><strong>401</strong>: Authorization required or customer id not found</div> <div><strong>404</strong>: location id or device does not exist.</div> <div><strong>500</strong>: internal server error</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Customer id | 
  **locationId** | **string**|  | 
  **mac** | **string**|  | 
  **timePeriod** | **string**| Any of \&quot;weekly\&quot;,\&quot;dailyToday\&quot;,\&quot;dailyYesterday\&quot;,\&quot;daily2DaysAgo\&quot;,\&quot;daily3DaysAgo\&quot;,\&quot;daily4DaysAgo\&quot;,\&quot;daily5DaysAgo\&quot;,\&quot;daily6DaysAgo\&quot;,\&quot;last30Days\&quot;,\&quot;last12Months\&quot; | 
 **optional** | ***CustomerApiCustomerPrototypeGetDeviceAppTimeAppsDataUsageOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CustomerApiCustomerPrototypeGetDeviceAppTimeAppsDataUsageOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **limit** | **optional.Float64**| Maximum number of apps to return. Defaults to 20 | 
 **grouping** | **optional.String**| typing of Grouping for the purposes of applying the limit. Can be: &#39;perTimeSlot&#39; ONLY | [default to perTimeSlot]

### Return type

**[]interface{}**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomerPrototypeGetDeviceAppTimeAppsOnlineTime**
> []interface{} CustomerPrototypeGetDeviceAppTimeAppsOnlineTime(ctx, id, locationId, mac, timePeriod, optional)
Fetch the AppTime Apps Online Time Stats for a Device.

<div><strong>200</strong>: Success, AppTime Stats returned.</div> <div><strong>401</strong>: Authorization required or customer id not found</div> <div><strong>404</strong>: location id or device does not exist.</div> <div><strong>500</strong>: internal server error</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Customer id | 
  **locationId** | **string**|  | 
  **mac** | **string**|  | 
  **timePeriod** | **string**| Any of \&quot;weekly\&quot;,\&quot;dailyToday\&quot;,\&quot;dailyYesterday\&quot;,\&quot;daily2DaysAgo\&quot;,\&quot;daily3DaysAgo\&quot;,\&quot;daily4DaysAgo\&quot;,\&quot;daily5DaysAgo\&quot;,\&quot;daily6DaysAgo\&quot;,\&quot;last30Days\&quot;,\&quot;last12Months\&quot; | 
 **optional** | ***CustomerApiCustomerPrototypeGetDeviceAppTimeAppsOnlineTimeOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CustomerApiCustomerPrototypeGetDeviceAppTimeAppsOnlineTimeOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **limit** | **optional.Float64**| Maximum number of apps to return. Defaults to 20 | 
 **grouping** | **optional.String**| typing of Grouping for the purposes of applying the limit. Can be: &#39;perTimeSlot&#39; ONLY | [default to perTimeSlot]

### Return type

**[]interface{}**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomerPrototypeGetDeviceAppTimeCategoriesDataUsage**
> []interface{} CustomerPrototypeGetDeviceAppTimeCategoriesDataUsage(ctx, id, locationId, mac, timePeriod, optional)
Fetch the AppTime Categories Data Usage Stats for a Device.

<div><strong>200</strong>: Success, AppTime Stats returned.</div> <div><strong>401</strong>: Authorization required or customer id not found</div> <div><strong>404</strong>: location id or device does not exist.</div> <div><strong>500</strong>: internal server error</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Customer id | 
  **locationId** | **string**|  | 
  **mac** | **string**|  | 
  **timePeriod** | **string**| Any of \&quot;weekly\&quot;,\&quot;dailyToday\&quot;,\&quot;dailyYesterday\&quot;,\&quot;daily2DaysAgo\&quot;,\&quot;daily3DaysAgo\&quot;,\&quot;daily4DaysAgo\&quot;,\&quot;daily5DaysAgo\&quot;,\&quot;daily6DaysAgo\&quot;,\&quot;last30Days\&quot;,\&quot;last12Months\&quot; | 
 **optional** | ***CustomerApiCustomerPrototypeGetDeviceAppTimeCategoriesDataUsageOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CustomerApiCustomerPrototypeGetDeviceAppTimeCategoriesDataUsageOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **limit** | **optional.Float64**| Maximum number of categories to return. Defaults to 20 | 
 **grouping** | **optional.String**| typing of Grouping for the purposes of applying the limit. Can be: &#39;overall&#39;|&#39;perTimeSlot&#39; | [default to perTimeSlot]

### Return type

**[]interface{}**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomerPrototypeGetDeviceAppTimeCategoriesOnlineTime**
> []interface{} CustomerPrototypeGetDeviceAppTimeCategoriesOnlineTime(ctx, id, locationId, mac, timePeriod, optional)
Fetch the AppTime Categories Online Time Stats for a Device.

<div><strong>200</strong>: Success, AppTime Stats returned.</div> <div><strong>401</strong>: Authorization required or customer id not found</div> <div><strong>404</strong>: location id or device does not exist.</div> <div><strong>500</strong>: internal server error</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Customer id | 
  **locationId** | **string**|  | 
  **mac** | **string**|  | 
  **timePeriod** | **string**| Any of \&quot;weekly\&quot;,\&quot;dailyToday\&quot;,\&quot;dailyYesterday\&quot;,\&quot;daily2DaysAgo\&quot;,\&quot;daily3DaysAgo\&quot;,\&quot;daily4DaysAgo\&quot;,\&quot;daily5DaysAgo\&quot;,\&quot;daily6DaysAgo\&quot;,\&quot;last30Days\&quot;,\&quot;last12Months\&quot; | 
 **optional** | ***CustomerApiCustomerPrototypeGetDeviceAppTimeCategoriesOnlineTimeOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CustomerApiCustomerPrototypeGetDeviceAppTimeCategoriesOnlineTimeOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **limit** | **optional.Float64**| Maximum number of categories to return. Defaults to 20 | 
 **grouping** | **optional.String**| typing of Grouping for the purposes of applying the limit. Can be: &#39;overall&#39;|&#39;perTimeSlot&#39; | [default to perTimeSlot]

### Return type

**[]interface{}**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomerPrototypeGetDeviceBandSteeringStats**
> XAny CustomerPrototypeGetDeviceBandSteeringStats(ctx, id, locationId, mac, optional)
Device band steering stats with all nodes for a particular MAC address.

<div><strong>200</strong>: Success.</div> <div><strong>401</strong>: Authorization required.</div> <div><strong>404</strong>: Location id does not exist.</div> <div><strong>500</strong>: Internal server error</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Customer id | 
  **locationId** | **string**| location id of devices and nodes | 
  **mac** | **string**| mac id of device | 
 **optional** | ***CustomerApiCustomerPrototypeGetDeviceBandSteeringStatsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CustomerApiCustomerPrototypeGetDeviceBandSteeringStatsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **granularity** | **optional.String**| days/hours/minutes | [default to days]
 **limit** | **optional.Float64**| X # of days/hours/minutes | [default to 7]
 **start** | **optional.Float64**| number of milliseconds elapsed since 1 January 1970 00:00:00 UTC. Defaults to now - (limit * granularity) | 

### Return type

[**XAny**](x-any.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomerPrototypeGetDeviceByMac**
> DeviceResponse CustomerPrototypeGetDeviceByMac(ctx, id, locationId, mac, optional)
Returns a single Device for a Location ID.

<div><strong>200</strong>: Success, device returned.</div> <div><strong>404</strong>: customer id or location id does not exist. Or, device not found in this network 's history.</div> <div><strong>500</strong>: internal server error</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Customer id | 
  **locationId** | **string**|  | 
  **mac** | **string**| mac of device | 
 **optional** | ***CustomerApiCustomerPrototypeGetDeviceByMacOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CustomerApiCustomerPrototypeGetDeviceByMacOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **include** | **optional.String**| can be &#39;bandwidthData&#39;, &#39;chartsData&#39; or both. None means &#39;bandwidthData&#39; only. | 
 **daysOffline** | **optional.Float64**| exclude devices disconnected longer than daysOffline. | 

### Return type

[**DeviceResponse**](DeviceResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomerPrototypeGetDeviceClientSteeringStats**
> XAny CustomerPrototypeGetDeviceClientSteeringStats(ctx, id, locationId, mac, optional)
Device client steering stats with all nodes for a particular MAC address.

<div><strong>200</strong>: Success.</div> <div><strong>401</strong>: Authorization required.</div> <div><strong>404</strong>: Location id does not exist.</div> <div><strong>500</strong>: Internal server error</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Customer id | 
  **locationId** | **string**| location id of devices and nodes | 
  **mac** | **string**| mac id of device | 
 **optional** | ***CustomerApiCustomerPrototypeGetDeviceClientSteeringStatsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CustomerApiCustomerPrototypeGetDeviceClientSteeringStatsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **granularity** | **optional.String**| days/hours/minutes | [default to days]
 **limit** | **optional.Float64**| X # of days/hours/minutes | [default to 7]
 **start** | **optional.Float64**| number of milliseconds elapsed since 1 January 1970 00:00:00 UTC. Defaults to now - (limit * granularity) | 

### Return type

[**XAny**](x-any.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomerPrototypeGetDeviceGroups**
> []XAny CustomerPrototypeGetDeviceGroups(ctx, id, locationId, networkId)
Get a list of device groups in a network, along with a list of member devices and group shares.

<div><strong>200</strong>: Success.</div> <div><strong>404</strong>: Location does not exist.</div> <div><strong>404</strong>: Network does not exist.</div> <div><strong>401</strong>: Unauthorized.</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Customer id | 
  **locationId** | **string**|  | 
  **networkId** | **string**|  | 

### Return type

[**[]XAny**](x-any.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomerPrototypeGetDeviceQoeMetrics**
> interface{} CustomerPrototypeGetDeviceQoeMetrics(ctx, id, locationId, mac, optional)
Device or pod QoE 15 minutes data.

<div><strong>200</strong>: Success.</div> <div><strong>401</strong>: Authorization required.</div> <div><strong>404</strong>: Location id does not exist.</div> <div><strong>500</strong>: Internal server error</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Customer id | 
  **locationId** | **string**|  | 
  **mac** | **string**| device mac address | 
 **optional** | ***CustomerApiCustomerPrototypeGetDeviceQoeMetricsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CustomerApiCustomerPrototypeGetDeviceQoeMetricsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **granularity** | **optional.String**| days/hours/minutes | [default to days]
 **limit** | **optional.Float64**| X # of days/hours/minutes | [default to 7]
 **timestampISOFormat** | **optional.Bool**| either timestamp utc number or ISO string | [default to false]

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomerPrototypeGetDeviceSecurity**
> DeviceResponse CustomerPrototypeGetDeviceSecurity(ctx, id, locationId, mac)
Returns the security policy Device for a Location ID.

<div><strong>200</strong>: Success, device returned.</div> <div><strong>404</strong>: customer id or location id does not exist. Or, device not found in this network 's history.</div> <div><strong>500</strong>: internal server error</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Customer id | 
  **locationId** | **string**|  | 
  **mac** | **string**| mac of device | 

### Return type

[**DeviceResponse**](DeviceResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomerPrototypeGetDeviceSecurityPolicyEvents**
> SecurityEventsResponse CustomerPrototypeGetDeviceSecurityPolicyEvents(ctx, id, locationId, mac, startTime, optional)
Get a Security Policy Events for Device for a Location ID.

<div><strong>200</strong>: Success.</div> <div><strong>401</strong>: Authorization required or customer id not found</div> <div><strong>404</strong>: Location id or WifiNetwork does not exist and is not known to Plume</div> <div><strong>500</strong>: Internal server error</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Customer id | 
  **locationId** | **string**|  | 
  **mac** | **string**| mac of device | 
  **startTime** | **time.Time**|  | 
 **optional** | ***CustomerApiCustomerPrototypeGetDeviceSecurityPolicyEventsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CustomerApiCustomerPrototypeGetDeviceSecurityPolicyEventsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **includes** | **optional.String**|  | 
 **limit** | **optional.Float64**|  | 
 **direction** | **optional.String**|  | 
 **protectionType** | **optional.String**|  | [default to ihp]

### Return type

[**SecurityEventsResponse**](SecurityEventsResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomerPrototypeGetDeviceSecurityPolicyHourlyCounts**
> Counts CustomerPrototypeGetDeviceSecurityPolicyHourlyCounts(ctx, id, locationId, mac)
Get a Security Policy Hourly Blocked Counts for a Device for a Location ID.

<div><strong>200</strong>: Success.</div> <div><strong>401</strong>: Authorization required or customer id not found</div> <div><strong>404</strong>: Location id or WifiNetwork does not exist and is not known to Plume</div> <div><strong>500</strong>: Internal server error</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Customer id | 
  **locationId** | **string**|  | 
  **mac** | **string**| mac of device | 

### Return type

[**Counts**](Counts.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomerPrototypeGetDeviceSoundingState**
> interface{} CustomerPrototypeGetDeviceSoundingState(ctx, id, locationId, optional)
Fetch the sounding states for eligible devices in this location

<div><strong>200</strong>: Success, device sounding states returned.</div> <div><strong>401</strong>: Authorization required or customer id not found</div> <div><strong>404</strong>: customer id or location id does not exist and is not known to Plume</div> <div><strong>500</strong>: internal server error</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Customer id | 
  **locationId** | **string**|  | 
 **optional** | ***CustomerApiCustomerPrototypeGetDeviceSoundingStateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CustomerApiCustomerPrototypeGetDeviceSoundingStateOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **mac** | **optional.String**| Optional mac address for single device lookup (fetches all devices by default) | 

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomerPrototypeGetDeviceSteeringWithAthena**
> interface{} CustomerPrototypeGetDeviceSteeringWithAthena(ctx, id, locationId, mac, optional)
Find all instances of the model.

<div><strong>200</strong>: Success.</div> <div><strong>401</strong>: Authorization required.</div> <div><strong>404</strong>: Location id does not exist.</div> <div><strong>500</strong>: Internal server error</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Customer id | 
  **locationId** | **string**|  | 
  **mac** | **string**|  | 
 **optional** | ***CustomerApiCustomerPrototypeGetDeviceSteeringWithAthenaOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CustomerApiCustomerPrototypeGetDeviceSteeringWithAthenaOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **order** | **optional.String**| desc || asc | [default to desc]
 **limit** | **optional.Float64**| 1000 max for deep:false and 10 max for deep:true | [default to 10]
 **startAt** | **optional.String**| find objects after this value | 
 **endAt** | **optional.String**| find objects before this value | 

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomerPrototypeGetDevices**
> InlineResponse20011 CustomerPrototypeGetDevices(ctx, id, locationId, optional)
Get all the devices for a Location ID, including the device name, icon to use, MAC and IP  address, connecting nodes and more.

All devices with 2g, 5g and 6g channel settings <div><strong>200</strong>: Success, array of Devices returned.</div> <div><strong>401</strong>: Authorization required or customer id not found</div> <div><strong>404</strong>: location id does not exist and is not known to Plume</div> <div><strong>500</strong>: internal server error</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Customer id | 
  **locationId** | **string**|  | 
 **optional** | ***CustomerApiCustomerPrototypeGetDevicesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CustomerApiCustomerPrototypeGetDevicesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **daysOffline** | **optional.Float64**|  | 
 **allSSIDs** | **optional.Bool**|  | 
 **showPartnerComponent** | **optional.Bool**|  | [default to true]

### Return type

[**InlineResponse20011**](inline_response_200_11.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomerPrototypeGetDhcp**
> interface{} CustomerPrototypeGetDhcp(ctx, id, locationId)
Get current DHCP Configuration details for a Location ID.

<div><strong>200</strong>: Success, current dhcp returned.</div> <div><strong>401</strong>: Authorization required or customer id not found</div> <div><strong>404</strong>: location id or dhcp does not exist.</div> <div><strong>500</strong>: internal server error</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Customer id | 
  **locationId** | **string**|  | 

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomerPrototypeGetDhcpReservation**
> DhcpReservation CustomerPrototypeGetDhcpReservation(ctx, id, locationId, mac)
Get current DHCP IP reservation details for a Location ID.

<div><strong>200</strong>: Success, current DhcpReservation returned.</div> <div><strong>401</strong>: Authorization required or customer id not found</div> <div><strong>404</strong>: location id or DhcpReservation does not exist.</div> <div><strong>422</strong>: mac is empty, or invalid.</div> <div><strong>500</strong>: internal server error</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Customer id | 
  **locationId** | **string**|  | 
  **mac** | **string**|  | 

### Return type

[**DhcpReservation**](DhcpReservation.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomerPrototypeGetDhcpReservations**
> []DhcpReservation CustomerPrototypeGetDhcpReservations(ctx, id, locationId)
Get current DHCP IP reservation details for a Location ID.

<div><strong>200</strong>: Success, current DhcpReservation returned.</div> <div><strong>401</strong>: Authorization required or customer id not found</div> <div><strong>404</strong>: location id or DhcpReservation does not exist.</div> <div><strong>422</strong>: mac is empty, or invalid.</div> <div><strong>500</strong>: internal server error</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Customer id | 
  **locationId** | **string**|  | 

### Return type

[**[]DhcpReservation**](DhcpReservation.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomerPrototypeGetDnsServers**
> DnsServers CustomerPrototypeGetDnsServers(ctx, id, locationId)
Get the current DNS IP addresses and settings for a Location ID.

<div><strong>200</strong>: Success, current DNS server settings returned.</div> <div><strong>401</strong>: Authorization required or customer id not found</div> <div><strong>404</strong>: NetworkConfiguration or DNS server settings does not exist.</div> <div><strong>500</strong>: internal server error</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Customer id | 
  **locationId** | **string**|  | 

### Return type

[**DnsServers**](DnsServers.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomerPrototypeGetDpp**
> interface{} CustomerPrototypeGetDpp(ctx, id, locationId)
Get the current DPP configuration for a Location ID.

<div><strong>200</strong>: Success, current DPP configuration returned.</div> <div><strong>401</strong>: Authorization required or customer id not found</div> <div><strong>500</strong>: internal server error</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Customer id | 
  **locationId** | **string**|  | 

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomerPrototypeGetDppAnnouncementsFromController**
> XAny CustomerPrototypeGetDppAnnouncementsFromController(ctx, id, locationId)
Returns DPP announcements from controller

<div><strong>200</strong>: Success.</div> <div><strong>401</strong>: Authorization required or customer id not found</div> <div><strong>404</strong>: Location id or WifiNetwork does not exist and is not known to Plume</div> <div><strong>500</strong>: Internal server error</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Customer id | 
  **locationId** | **string**|  | 

### Return type

[**XAny**](x-any.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomerPrototypeGetEmployeeNetworkAppTimeAppsDataUsage**
> []interface{} CustomerPrototypeGetEmployeeNetworkAppTimeAppsDataUsage(ctx, id, locationId, networkId, timePeriod, optional)
Fetch the AppTime Apps Data Usage Stats for fronthaul network.

<div><strong>200</strong>: Success, AppTime Stats returned.</div> <div><strong>401</strong>: Authorization required or customer id not found</div> <div><strong>404</strong>: location id or secondary networks does not exist.</div> <div><strong>500</strong>: internal server error</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Customer id | 
  **locationId** | **string**|  | 
  **networkId** | **string**|  | 
  **timePeriod** | **string**| Any of \&quot;weekly\&quot;,\&quot;dailyToday\&quot;,\&quot;dailyYesterday\&quot;,\&quot;daily2DaysAgo\&quot;,\&quot;daily3DaysAgo\&quot;,\&quot;daily4DaysAgo\&quot;,\&quot;daily5DaysAgo\&quot;,\&quot;daily6DaysAgo\&quot;,\&quot;last30Days\&quot;,\&quot;last12Months\&quot; | 
 **optional** | ***CustomerApiCustomerPrototypeGetEmployeeNetworkAppTimeAppsDataUsageOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CustomerApiCustomerPrototypeGetEmployeeNetworkAppTimeAppsDataUsageOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **limit** | **optional.Float64**| Maximum number of apps to return. Defaults to 20 | 
 **grouping** | **optional.String**| typing of Grouping for the purposes of applying the limit. Can be: &#39;perTimeSlot&#39; ONLY | [default to perTimeSlot]

### Return type

**[]interface{}**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomerPrototypeGetEmployeeNetworkAppTimeAppsOnlineTime**
> []interface{} CustomerPrototypeGetEmployeeNetworkAppTimeAppsOnlineTime(ctx, id, locationId, networkId, timePeriod, optional)
Fetch the AppTime Apps Online Time Stats for fronthaul network.

<div><strong>200</strong>: Success, AppTime Stats returned.</div> <div><strong>401</strong>: Authorization required or customer id not found</div> <div><strong>404</strong>: location id or secondary networks does not exist.</div> <div><strong>500</strong>: internal server error</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Customer id | 
  **locationId** | **string**|  | 
  **networkId** | **string**|  | 
  **timePeriod** | **string**| Any of \&quot;weekly\&quot;,\&quot;dailyToday\&quot;,\&quot;dailyYesterday\&quot;,\&quot;daily2DaysAgo\&quot;,\&quot;daily3DaysAgo\&quot;,\&quot;daily4DaysAgo\&quot;,\&quot;daily5DaysAgo\&quot;,\&quot;daily6DaysAgo\&quot;,\&quot;last30Days\&quot;,\&quot;last12Months\&quot; | 
 **optional** | ***CustomerApiCustomerPrototypeGetEmployeeNetworkAppTimeAppsOnlineTimeOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CustomerApiCustomerPrototypeGetEmployeeNetworkAppTimeAppsOnlineTimeOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **limit** | **optional.Float64**| Maximum number of apps to return. Defaults to 20 | 
 **grouping** | **optional.String**| typing of Grouping for the purposes of applying the limit. Can be: &#39;perTimeSlot&#39; ONLY | [default to perTimeSlot]

### Return type

**[]interface{}**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomerPrototypeGetEmployeeNetworkAppTimeCategoriesDataUsage**
> []interface{} CustomerPrototypeGetEmployeeNetworkAppTimeCategoriesDataUsage(ctx, id, locationId, networkId, timePeriod, optional)
Fetch the AppTime Categories Data Usage Stats for fronthaul network.

<div><strong>200</strong>: Success, AppTime Stats returned.</div> <div><strong>401</strong>: Authorization required or customer id not found</div> <div><strong>404</strong>: location id or secondary networks does not exist.</div> <div><strong>500</strong>: internal server error</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Customer id | 
  **locationId** | **string**|  | 
  **networkId** | **string**|  | 
  **timePeriod** | **string**| Any of \&quot;weekly\&quot;,\&quot;dailyToday\&quot;,\&quot;dailyYesterday\&quot;,\&quot;daily2DaysAgo\&quot;,\&quot;daily3DaysAgo\&quot;,\&quot;daily4DaysAgo\&quot;,\&quot;daily5DaysAgo\&quot;,\&quot;daily6DaysAgo\&quot;,\&quot;last30Days\&quot;,\&quot;last12Months\&quot; | 
 **optional** | ***CustomerApiCustomerPrototypeGetEmployeeNetworkAppTimeCategoriesDataUsageOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CustomerApiCustomerPrototypeGetEmployeeNetworkAppTimeCategoriesDataUsageOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **limit** | **optional.Float64**| Maximum number of categories to return. Defaults to 20 | 
 **grouping** | **optional.String**| typing of Grouping for the purposes of applying the limit. Can be: &#39;overall&#39;|&#39;perTimeSlot&#39; | [default to perTimeSlot]

### Return type

**[]interface{}**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomerPrototypeGetEmployeeNetworkAppTimeCategoriesOnlineTime**
> []interface{} CustomerPrototypeGetEmployeeNetworkAppTimeCategoriesOnlineTime(ctx, id, locationId, networkId, timePeriod, optional)
Fetch the AppTime Categories Online Time Stats for fronthaul network.

<div><strong>200</strong>: Success, AppTime Stats returned.</div> <div><strong>401</strong>: Authorization required or customer id not found</div> <div><strong>404</strong>: location id or secondary networks does not exist.</div> <div><strong>500</strong>: internal server error</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Customer id | 
  **locationId** | **string**|  | 
  **networkId** | **string**|  | 
  **timePeriod** | **string**| Any of \&quot;weekly\&quot;,\&quot;dailyToday\&quot;,\&quot;dailyYesterday\&quot;,\&quot;daily2DaysAgo\&quot;,\&quot;daily3DaysAgo\&quot;,\&quot;daily4DaysAgo\&quot;,\&quot;daily5DaysAgo\&quot;,\&quot;daily6DaysAgo\&quot;,\&quot;last30Days\&quot;,\&quot;last12Months\&quot; | 
 **optional** | ***CustomerApiCustomerPrototypeGetEmployeeNetworkAppTimeCategoriesOnlineTimeOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CustomerApiCustomerPrototypeGetEmployeeNetworkAppTimeCategoriesOnlineTimeOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **limit** | **optional.Float64**| Maximum number of categories to return. Defaults to 20 | 
 **grouping** | **optional.String**| typing of Grouping for the purposes of applying the limit. Can be: &#39;overall&#39;|&#39;perTimeSlot&#39; | [default to perTimeSlot]

### Return type

**[]interface{}**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomerPrototypeGetEntitledAccessList**
> interface{} CustomerPrototypeGetEntitledAccessList(ctx, id)
Get a list of all locations on which you are assigned as a manager.

<div><strong>200</strong>: Success.</div> <div><strong>400</strong>: Required fields missing or field type is incorrect.</div> <div><strong>404</strong>: Location does not exist.</div> <div><strong>422</strong>: Invalid email, name, access type or manager is already assigned to this location </div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Customer id | 

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomerPrototypeGetEventHistory**
> []interface{} CustomerPrototypeGetEventHistory(ctx, id, locationId, optional)
Fetch the event history for this location

<div><strong>200</strong>: Success, event array returned.</div> <div><strong>401</strong>: Authorization required or customer id not found</div> <div><strong>404</strong>: customer id or location id does not exist and is not known to Plume</div> <div><strong>500</strong>: internal server error</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Customer id | 
  **locationId** | **string**|  | 
 **optional** | ***CustomerApiCustomerPrototypeGetEventHistoryOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CustomerApiCustomerPrototypeGetEventHistoryOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **from** | **optional.Float64**| UTC unix ts | 
 **to** | **optional.Float64**| UTC unix ts, defaults to now | 
 **category** | **optional.String**| Filter events by category (Motion or Plume [config changes]). Multiple categories can be passed as a comma-separated string. Default is both. | 
 **limit** | **optional.Float64**| Maximum number of events to return. Defaults to 10 | 
 **sort** | **optional.Bool**| whether the returned events will be post-sorted by timestamp | 

### Return type

**[]interface{}**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomerPrototypeGetFastInterference**
> XAny CustomerPrototypeGetFastInterference(ctx, id, locationId)
Get from Controller Fast interference status.

<div><strong>200</strong>: Ok.</div> <div><strong>401</strong>: Authorization required.</div> <div><strong>404</strong>: Location id, does not exist.</div> <div><strong>500</strong>: Internal server error.</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Customer id | 
  **locationId** | **string**|  | 

### Return type

[**XAny**](x-any.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomerPrototypeGetFirmwareUpgradeStatus**
> []interface{} CustomerPrototypeGetFirmwareUpgradeStatus(ctx, id, locationId)
Firmware Upgrade Status

<div><strong>200</strong>: Success, response object returned.</div> <div><strong>401</strong>: Authorization required or customer id not found</div> <div><strong>404</strong>: customer id or location id does not exist and is not known to Plume</div> <div><strong>500</strong>: internal server error</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Customer id | 
  **locationId** | **string**|  | 

### Return type

**[]interface{}**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomerPrototypeGetFlowStats**
> interface{} CustomerPrototypeGetFlowStats(ctx, id, locationId)
GET the flow stats configuration

<div><strong>200</strong>: Success, current flow stats configuration returned.</div> <div><strong>401</strong>: Authorization required or customer id not found.</div> <div><strong>404</strong>: location id does not exist.</div> <div><strong>500</strong>: internal server error</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Customer id | 
  **locationId** | **string**|  | 

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomerPrototypeGetForceGraph**
> string CustomerPrototypeGetForceGraph(ctx, id, locationId, optional)
HTML or JSON (vertices[] + edges[]) used to display a Network Topology.

<div>The HTML and JSON to initialize and dynamically display and update a Topology.</div> <div>The JSON can also be used to get a network's list of nodes + devices (a.k.a. vertices) and links (a.k.a., edges).</div><div>&nbsp;</div> <div><strong>200</strong>: Success, HTML or JSON returned depending on \"Accept\" HTTP header.</div> <div><strong>404</strong>: customer id or location id does not exist and is not known to Plume</div> <div><strong>500</strong>: internal server error</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Customer id | 
  **locationId** | **string**|  | 
 **optional** | ***CustomerApiCustomerPrototypeGetForceGraphOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CustomerApiCustomerPrototypeGetForceGraphOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **ip** | **optional.String**| deprecated and optional IP address of client displaying the topology | 
 **mac** | **optional.String**| optional mac address of client displaying the topology | 
 **authKey** | **optional.String**| PubNub authKey | 
 **subscribeKey** | **optional.String**| PubNub subscribeKey | 
 **view** | **optional.String**| view template override (e.g., iguana) | 
 **allSSIDs** | **optional.Bool**|  | 
 **showPartnerComponent** | **optional.Bool**|  | 

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomerPrototypeGetFrontHaulNetworks**
> NetworkConfig CustomerPrototypeGetFrontHaulNetworks(ctx, id, locationId)
Get the Front Haul Portal configs for a given Location ID.

<div><strong>200</strong>: Success, FrontHaul Networks returned.</div> <div><strong>401</strong>: Authorization required or customer id not found</div> <div><strong>404</strong>: location id or secondary network does not exist.</div> <div><strong>500</strong>: internal server error</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Customer id | 
  **locationId** | **string**|  | 

### Return type

[**NetworkConfig**](NetworkConfig.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomerPrototypeGetFrontHaulsDpp**
> interface{} CustomerPrototypeGetFrontHaulsDpp(ctx, id, locationId, networkId)
Get the current DPP configurator for a Location ID.

<div><strong>200</strong>: Success, current DPP configurator returned.</div> <div><strong>401</strong>: Authorization required or customer id not found</div> <div><strong>500</strong>: internal server error</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Customer id | 
  **locationId** | **string**|  | 
  **networkId** | **string**|  | 

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomerPrototypeGetFrontlineStorage**
> FrontlineStorage CustomerPrototypeGetFrontlineStorage(ctx, id, locationId)
Fetch the frontline storage for this location

<div><strong>200</strong>: Success, HomeSecurity object returned.</div> <div><strong>401</strong>: Authorization required or customer id not found</div> <div><strong>404</strong>: customer id or location id does not exist and is not known to Plume</div> <div><strong>500</strong>: internal server error</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Customer id | 
  **locationId** | **string**|  | 

### Return type

[**FrontlineStorage**](FrontlineStorage.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomerPrototypeGetGdprCaptivePortalsData**
> interface{} CustomerPrototypeGetGdprCaptivePortalsData(ctx, id, locationId, networkId, email, localEndDate, localStartDate)
Fetch the Gdpr Captive Portals data for a guest.

<div><strong>200</strong>: Success, GDPR Captive Portals data returned.</div> <div><strong>401</strong>: Authorization required or customer id not found</div> <div><strong>404</strong>: location id or secondary networks does not exist.</div> <div><strong>500</strong>: internal server error</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Customer id | 
  **locationId** | **string**|  | 
  **networkId** | **string**|  | 
  **email** | **string**|  | 
  **localEndDate** | **string**|  | 
  **localStartDate** | **string**|  | 

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomerPrototypeGetGroupOfUnassignedDevicesFreezePolicy**
> XAny CustomerPrototypeGetGroupOfUnassignedDevicesFreezePolicy(ctx, id, locationId)
Get GroupOfUnassignedDevices freeze policy for a Location ID.

<div><strong>200</strong>: Ok.</div> <div><strong>401</strong>: Authorization required or customer id not found.</div> <div><strong>404</strong>: Location id, does not exist.</div> <div><strong>500</strong>: Internal server error</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Customer id | 
  **locationId** | **string**|  | 

### Return type

[**XAny**](x-any.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomerPrototypeGetGroupOfUnassignedDevicesSecurityPolicy**
> Person CustomerPrototypeGetGroupOfUnassignedDevicesSecurityPolicy(ctx, id, locationId)
Get a Security Policy for a Location ID.

<div><strong>200</strong>: Success.</div> <div><strong>401</strong>: Authorization required or customer id not found</div> <div><strong>404</strong>: Location id or WifiNetwork does not exist and is not known to Plume</div> <div><strong>500</strong>: Internal server error</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Customer id | 
  **locationId** | **string**|  | 

### Return type

[**Person**](Person.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomerPrototypeGetGroupOfUnassignedDevicesSecurityPolicyEvents**
> SecurityEventsResponse CustomerPrototypeGetGroupOfUnassignedDevicesSecurityPolicyEvents(ctx, id, locationId, startTime, optional)
Get a Security Policy Events for groupOfUnassignedDevices for a Location ID.

<div><strong>200</strong>: Success.</div> <div><strong>401</strong>: Authorization required or customer id not found</div> <div><strong>404</strong>: Location id or WifiNetwork does not exist and is not known to Plume</div> <div><strong>500</strong>: Internal server error</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Customer id | 
  **locationId** | **string**|  | 
  **startTime** | **time.Time**|  | 
 **optional** | ***CustomerApiCustomerPrototypeGetGroupOfUnassignedDevicesSecurityPolicyEventsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CustomerApiCustomerPrototypeGetGroupOfUnassignedDevicesSecurityPolicyEventsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **includes** | **optional.String**|  | 
 **limit** | **optional.Float64**|  | 
 **direction** | **optional.String**|  | 
 **protectionType** | **optional.String**|  | [default to ihp]

### Return type

[**SecurityEventsResponse**](SecurityEventsResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomerPrototypeGetGroupOfUnassignedDevicesSecurityPolicyHourlyCounts**
> Counts CustomerPrototypeGetGroupOfUnassignedDevicesSecurityPolicyHourlyCounts(ctx, id, locationId)
Get a Security Policy Hourly Blocked Counts for group Of Unassigned Devices for a Location ID.

<div><strong>200</strong>: Success.</div> <div><strong>401</strong>: Authorization required or customer id not found</div> <div><strong>404</strong>: Location id or WifiNetwork does not exist and is not known to Plume</div> <div><strong>500</strong>: Internal server error</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Customer id | 
  **locationId** | **string**|  | 

### Return type

[**Counts**](Counts.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomerPrototypeGetGroups**
> []XAny CustomerPrototypeGetGroups(ctx, id)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Customer id | 

### Return type

[**[]XAny**](x-any.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomerPrototypeGetGuestNetworkAppTimeAppsDataUsage**
> []interface{} CustomerPrototypeGetGuestNetworkAppTimeAppsDataUsage(ctx, id, locationId, networkId, timePeriod, optional)
Fetch the AppTime Apps Data Usage Stats for captivePortal network.

<div><strong>200</strong>: Success, AppTime Stats returned.</div> <div><strong>401</strong>: Authorization required or customer id not found</div> <div><strong>404</strong>: location id or secondary networks does not exist.</div> <div><strong>500</strong>: internal server error</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Customer id | 
  **locationId** | **string**|  | 
  **networkId** | **string**|  | 
  **timePeriod** | **string**| Any of \&quot;weekly\&quot;,\&quot;dailyToday\&quot;,\&quot;dailyYesterday\&quot;,\&quot;daily2DaysAgo\&quot;,\&quot;daily3DaysAgo\&quot;,\&quot;daily4DaysAgo\&quot;,\&quot;daily5DaysAgo\&quot;,\&quot;daily6DaysAgo\&quot;,\&quot;last30Days\&quot;,\&quot;last12Months\&quot; | 
 **optional** | ***CustomerApiCustomerPrototypeGetGuestNetworkAppTimeAppsDataUsageOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CustomerApiCustomerPrototypeGetGuestNetworkAppTimeAppsDataUsageOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **limit** | **optional.Float64**| Maximum number of apps to return. Defaults to 20 | 
 **grouping** | **optional.String**| typing of Grouping for the purposes of applying the limit. Can be: &#39;perTimeSlot&#39; ONLY | [default to perTimeSlot]

### Return type

**[]interface{}**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomerPrototypeGetGuestNetworkAppTimeAppsOnlineTime**
> []interface{} CustomerPrototypeGetGuestNetworkAppTimeAppsOnlineTime(ctx, id, locationId, networkId, timePeriod, optional)
Fetch the AppTime Apps Online Time Stats for captivePortal network.

<div><strong>200</strong>: Success, AppTime Stats returned.</div> <div><strong>401</strong>: Authorization required or customer id not found</div> <div><strong>404</strong>: location id or secondary networks does not exist.</div> <div><strong>500</strong>: internal server error</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Customer id | 
  **locationId** | **string**|  | 
  **networkId** | **string**|  | 
  **timePeriod** | **string**| Any of \&quot;weekly\&quot;,\&quot;dailyToday\&quot;,\&quot;dailyYesterday\&quot;,\&quot;daily2DaysAgo\&quot;,\&quot;daily3DaysAgo\&quot;,\&quot;daily4DaysAgo\&quot;,\&quot;daily5DaysAgo\&quot;,\&quot;daily6DaysAgo\&quot;,\&quot;last30Days\&quot;,\&quot;last12Months\&quot; | 
 **optional** | ***CustomerApiCustomerPrototypeGetGuestNetworkAppTimeAppsOnlineTimeOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CustomerApiCustomerPrototypeGetGuestNetworkAppTimeAppsOnlineTimeOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **limit** | **optional.Float64**| Maximum number of apps to return. Defaults to 20 | 
 **grouping** | **optional.String**| typing of Grouping for the purposes of applying the limit. Can be: &#39;perTimeSlot&#39; ONLY | [default to perTimeSlot]

### Return type

**[]interface{}**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomerPrototypeGetGuestNetworkAppTimeCategoriesDataUsage**
> []interface{} CustomerPrototypeGetGuestNetworkAppTimeCategoriesDataUsage(ctx, id, locationId, networkId, timePeriod, optional)
Fetch the AppTime Categories Data Usage Stats for captivePortal network.

<div><strong>200</strong>: Success, AppTime Stats returned.</div> <div><strong>401</strong>: Authorization required or customer id not found</div> <div><strong>404</strong>: location id or secondary networks does not exist.</div> <div><strong>500</strong>: internal server error</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Customer id | 
  **locationId** | **string**|  | 
  **networkId** | **string**|  | 
  **timePeriod** | **string**| Any of \&quot;weekly\&quot;,\&quot;dailyToday\&quot;,\&quot;dailyYesterday\&quot;,\&quot;daily2DaysAgo\&quot;,\&quot;daily3DaysAgo\&quot;,\&quot;daily4DaysAgo\&quot;,\&quot;daily5DaysAgo\&quot;,\&quot;daily6DaysAgo\&quot;,\&quot;last30Days\&quot;,\&quot;last12Months\&quot; | 
 **optional** | ***CustomerApiCustomerPrototypeGetGuestNetworkAppTimeCategoriesDataUsageOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CustomerApiCustomerPrototypeGetGuestNetworkAppTimeCategoriesDataUsageOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **limit** | **optional.Float64**| Maximum number of categories to return. Defaults to 20 | 
 **grouping** | **optional.String**| typing of Grouping for the purposes of applying the limit. Can be: &#39;overall&#39;|&#39;perTimeSlot&#39; | [default to perTimeSlot]

### Return type

**[]interface{}**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomerPrototypeGetGuestNetworkAppTimeCategoriesOnlineTime**
> []interface{} CustomerPrototypeGetGuestNetworkAppTimeCategoriesOnlineTime(ctx, id, locationId, networkId, timePeriod, optional)
Fetch the AppTime Categories Online Time Stats for captivePortal network.

<div><strong>200</strong>: Success, AppTime Stats returned.</div> <div><strong>401</strong>: Authorization required or customer id not found</div> <div><strong>404</strong>: location id or secondary networks does not exist.</div> <div><strong>500</strong>: internal server error</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Customer id | 
  **locationId** | **string**|  | 
  **networkId** | **string**|  | 
  **timePeriod** | **string**| Any of \&quot;weekly\&quot;,\&quot;dailyToday\&quot;,\&quot;dailyYesterday\&quot;,\&quot;daily2DaysAgo\&quot;,\&quot;daily3DaysAgo\&quot;,\&quot;daily4DaysAgo\&quot;,\&quot;daily5DaysAgo\&quot;,\&quot;daily6DaysAgo\&quot;,\&quot;last30Days\&quot;,\&quot;last12Months\&quot; | 
 **optional** | ***CustomerApiCustomerPrototypeGetGuestNetworkAppTimeCategoriesOnlineTimeOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CustomerApiCustomerPrototypeGetGuestNetworkAppTimeCategoriesOnlineTimeOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **limit** | **optional.Float64**| Maximum number of categories to return. Defaults to 20 | 
 **grouping** | **optional.String**| typing of Grouping for the purposes of applying the limit. Can be: &#39;overall&#39;|&#39;perTimeSlot&#39; | [default to perTimeSlot]

### Return type

**[]interface{}**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomerPrototypeGetHomeAwayLocationEvents**
> []interface{} CustomerPrototypeGetHomeAwayLocationEvents(ctx, id, locationId, optional)
Fetch the all the Homeaway events history for this location

<div><strong>200</strong>: Success, event array returned.</div> <div><strong>401</strong>: Authorization required or customer id not found</div> <div><strong>404</strong>: customer id or location id does not exist and is not known to Plume</div> <div><strong>500</strong>: internal server error</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Customer id | 
  **locationId** | **string**|  | 
 **optional** | ***CustomerApiCustomerPrototypeGetHomeAwayLocationEventsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CustomerApiCustomerPrototypeGetHomeAwayLocationEventsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **from** | **optional.Float64**| UTC unix epoch ms, defaults to 1 week ago | 
 **to** | **optional.Float64**| UTC unix epoch ms, defaults to now | 
 **limit** | **optional.Float64**| Maximum number of events to return. Defaults to 100 | 

### Return type

**[]interface{}**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomerPrototypeGetHomeSecurity**
> HomeSecurity CustomerPrototypeGetHomeSecurity(ctx, id, locationId)
Fetch the home security configuration for this location

<div><strong>200</strong>: Success, HomeSecurity object returned.</div> <div><strong>401</strong>: Authorization required or customer id not found</div> <div><strong>404</strong>: customer id or location id does not exist and is not known to Plume</div> <div><strong>500</strong>: internal server error</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Customer id | 
  **locationId** | **string**|  | 

### Return type

[**HomeSecurity**](HomeSecurity.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomerPrototypeGetIPv6**
> interface{} CustomerPrototypeGetIPv6(ctx, id, locationId)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Customer id | 
  **locationId** | **string**|  | 

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomerPrototypeGetKvConfigs**
> []KvConfig CustomerPrototypeGetKvConfigs(ctx, id, locationId, nodeId)
Retrieve all kvConfigs on a particular Node for a Location ID.

<div><strong>200</strong>: Success, your new info looks good.</div> <div><strong>401</strong>: Authorization required or customer id not found</div> <div><strong>404</strong>: location id does not exist.</div> <div><strong>422</strong>: nodeId must be defined.</div> <div><strong>425</strong>: nodeId must belong to the location.</div> <div><strong>500</strong>: internal server error</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Customer id | 
  **locationId** | **string**|  | 
  **nodeId** | **string**|  | 

### Return type

[**[]KvConfig**](KvConfig.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomerPrototypeGetLocationAppTime**
> interface{} CustomerPrototypeGetLocationAppTime(ctx, id, locationId)
Get a Location's AppTime config by location ID.

<div><strong>200</strong>: Success.</div> <div><strong>400</strong>: Required fields missing or field type is incorrect.</div> <div><strong>401</strong>: Authorization required or customer id not found</div> <div><strong>404</strong>: Location id not exist and is not known to Plume</div> <div><strong>500</strong>: Internal server error</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Customer id | 
  **locationId** | **string**|  | 

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomerPrototypeGetLocationCapabilities**
> CapabilitiesResponse CustomerPrototypeGetLocationCapabilities(ctx, id, locationId)
Get the non-base feature capabilities supported by a particular Location ID.

<div>The controller will implement logic to determine the non-base features supported by the Pods in the location ID. The feature capability is determined on the system level, and not per individual Pod.</div> <div>The mobile apps or other WebUIs should only show the UI for a feature if the disabled value equals \"false\".</div> <div>&nbsp;</div> <div><strong>200</strong>: Success, current Capabilities returned.</div> <div><strong>401</strong>: Authorization required or customer id not found</div> <div><strong>404</strong>: location id does not exist.</div> <div><strong>500</strong>: internal server error</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Customer id | 
  **locationId** | **string**|  | 

### Return type

[**CapabilitiesResponse**](CapabilitiesResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomerPrototypeGetLocationConfigAuditEvents**
> ConfigAuditEventsResponse CustomerPrototypeGetLocationConfigAuditEvents(ctx, id, locationId, includes, startTime, optional)
Get a Config Audit Trail Events for a Location ID.

<div><strong>200</strong>: Success.</div> <div><strong>401</strong>: Authorization required or customer id not found</div> <div><strong>404</strong>: Location id or WifiNetwork does not exist and is not known to Plume</div> <div><strong>500</strong>: Internal server error</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Customer id | 
  **locationId** | **string**|  | 
  **includes** | **string**|  | 
  **startTime** | **time.Time**|  | 
 **optional** | ***CustomerApiCustomerPrototypeGetLocationConfigAuditEventsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CustomerApiCustomerPrototypeGetLocationConfigAuditEventsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **limit** | **optional.Float64**|  | 
 **direction** | **optional.String**|  | 

### Return type

[**ConfigAuditEventsResponse**](ConfigAuditEventsResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomerPrototypeGetLocationFreezeAutoExpire**
> XAny CustomerPrototypeGetLocationFreezeAutoExpire(ctx, id, locationId)
Get all devices/persons except some to be frozen for a Location ID.

<div><strong>200</strong>: Success, updated.</div> <div><strong>401</strong>: Authorization required or customer id not found.</div> <div><strong>404</strong>: Location id, does not exist.</div> <div><strong>500</strong>: Internal server error</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Customer id | 
  **locationId** | **string**|  | 

### Return type

[**XAny**](x-any.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomerPrototypeGetLocationGuardEventStats**
> SecurityEventsResponse CustomerPrototypeGetLocationGuardEventStats(ctx, id, locationId, timePeriod, optional)
Get the Guard Event Stats for a Location ID.

<div><strong>200</strong>: Success.</div> <div><strong>401</strong>: Authorization required or customer id not found</div> <div><strong>404</strong>: Location id or WifiNetwork does not exist and is not known to Plume</div> <div><strong>500</strong>: Internal server error</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Customer id | 
  **locationId** | **string**|  | 
  **timePeriod** | **string**| Any of \&quot;last24Hours\&quot;,\&quot;last7Days\&quot;,\&quot;last30Days\&quot; | 
 **optional** | ***CustomerApiCustomerPrototypeGetLocationGuardEventStatsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CustomerApiCustomerPrototypeGetLocationGuardEventStatsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **macs** | **optional.String**| array of macs[] | 
 **persons** | **optional.String**| array of personIds[] | 
 **protectionType** | **optional.String**| filter by protectionType: ihp | ohp. Returns all types by default. | 
 **eventTypes** | **optional.String**| filter by event type, any combo of: &#39;adBlocking&#39;,&#39;teenagers&#39;,&#39;kids&#39;,&#39;adultAndSensitive&#39;,&#39;secureAndProtect&#39;,&#39;ipThreatOutbound&#39;,&#39;ipThreatInbound&#39;, &#39;iotProtect&#39;. Returns all types by default. | 
 **groupOfUnassignedDevices** | **optional.Bool**| to include the group of unassigned devices | [default to false]

### Return type

[**SecurityEventsResponse**](SecurityEventsResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomerPrototypeGetLocationGuardEventsTldOrIp**
> SecurityEventsResponse CustomerPrototypeGetLocationGuardEventsTldOrIp(ctx, id, locationId, optional)
Get the Guard Event Domain Groups for a Location ID.

<div><strong>200</strong>: Success.</div> <div><strong>401</strong>: Authorization required or customer id not found</div> <div><strong>404</strong>: Location id or WifiNetwork does not exist and is not known to Plume</div> <div><strong>500</strong>: Internal server error</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Customer id | 
  **locationId** | **string**|  | 
 **optional** | ***CustomerApiCustomerPrototypeGetLocationGuardEventsTldOrIpOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CustomerApiCustomerPrototypeGetLocationGuardEventsTldOrIpOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **macs** | **optional.String**| array of macs[] | 
 **persons** | **optional.String**| array of personIds[] | 
 **tldOrIp** | **optional.String**| top level domain or IP address | 
 **protectionType** | **optional.String**| filter by protectionType: ihp | ohp. Returns all types by default. | 
 **eventTypes** | **optional.String**| filter by event type, any combo of: &#39;adBlocking&#39;,&#39;teenagers&#39;,&#39;kids&#39;,&#39;adultAndSensitive&#39;,&#39;secureAndProtect&#39;,&#39;ipThreatOutbound&#39;,&#39;ipThreatInbound&#39;, &#39;iotProtect&#39;. Returns all types by default. | 
 **timePeriod** | **optional.String**| Any of \&quot;last24Hours\&quot;, \&quot;last7Days\&quot;, \&quot;last30Days\&quot; | [default to last30Days]
 **groupOfUnassignedDevices** | **optional.Bool**| to include the group of unassigned devices | [default to false]

### Return type

[**SecurityEventsResponse**](SecurityEventsResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomerPrototypeGetLocationGuardPersonEventsSummary**
> PersonEventStats CustomerPrototypeGetLocationGuardPersonEventsSummary(ctx, id, locationId, timePeriod)
Get the Guard Event Stats for all persons in a Location ID.

<div><strong>200</strong>: Success.</div> <div><strong>401</strong>: Authorization required or customer id not found</div> <div><strong>404</strong>: Location id or WifiNetwork does not exist and is not known to Plume</div> <div><strong>500</strong>: Internal server error</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Customer id | 
  **locationId** | **string**|  | 
  **timePeriod** | **string**| Any of \&quot;last24Hours\&quot;,\&quot;last7Days\&quot;,\&quot;last30Days\&quot; | 

### Return type

[**PersonEventStats**](personEventStats.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomerPrototypeGetLocationKvStates**
> interface{} CustomerPrototypeGetLocationKvStates(ctx, id, locationId)
Retrieve all kvStates on a particular Node for a Location ID.

<div><strong>200</strong>: Success, your new info looks good.</div> <div><strong>401</strong>: Authorization required or customer id not found</div> <div><strong>404</strong>: location id does not exist.</div> <div><strong>422</strong>: nodeId must be defined.</div> <div><strong>425</strong>: nodeId must belong to the location.</div> <div><strong>500</strong>: internal server error</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Customer id | 
  **locationId** | **string**|  | 

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomerPrototypeGetLocationOverlordConfigs**
> interface{} CustomerPrototypeGetLocationOverlordConfigs(ctx, id, locationId)
Gets all the configs from Overlord for a specified location.

<div><strong>200</strong>: Success, got the data.</div> <div><strong>401</strong>: Authorization required or customer id not found.</div> <div><strong>404</strong>: Location id, does not exist.</div> <div><strong>500</strong>: Internal server error</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Customer id | 
  **locationId** | **string**|  | 

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomerPrototypeGetLocationQoe**
> interface{} CustomerPrototypeGetLocationQoe(ctx, id, locationId)
Get QoE recent 1 minute data for a whole location.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Customer id | 
  **locationId** | **string**|  | 

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomerPrototypeGetLocationRooms**
> interface{} CustomerPrototypeGetLocationRooms(ctx, id, locationId)
Get a Location's Rooms config by location ID.

<div><strong>200</strong>: Success.</div> <div><strong>400</strong>: Required fields missing or field type is incorrect.</div> <div><strong>401</strong>: Authorization required or customer id not found</div> <div><strong>404</strong>: Location id not exist and is not known to Plume</div> <div><strong>500</strong>: Internal server error</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Customer id | 
  **locationId** | **string**|  | 

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomerPrototypeGetLocationSecurityPolicy**
> Person CustomerPrototypeGetLocationSecurityPolicy(ctx, id, locationId)
Get a Security Policy for a Location ID.

<div><strong>200</strong>: Success.</div> <div><strong>401</strong>: Authorization required or customer id not found</div> <div><strong>404</strong>: Location id or WifiNetwork does not exist and is not known to Plume</div> <div><strong>500</strong>: Internal server error</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Customer id | 
  **locationId** | **string**|  | 

### Return type

[**Person**](Person.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomerPrototypeGetLocationSecurityPolicyEvents**
> SecurityEventsResponse CustomerPrototypeGetLocationSecurityPolicyEvents(ctx, id, locationId, startTime, optional)
Get a Security Policy Events for a Location ID.

<div><strong>200</strong>: Success.</div> <div><strong>401</strong>: Authorization required or customer id not found</div> <div><strong>404</strong>: Location id or WifiNetwork does not exist and is not known to Plume</div> <div><strong>500</strong>: Internal server error</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Customer id | 
  **locationId** | **string**|  | 
  **startTime** | **time.Time**|  | 
 **optional** | ***CustomerApiCustomerPrototypeGetLocationSecurityPolicyEventsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CustomerApiCustomerPrototypeGetLocationSecurityPolicyEventsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **includes** | **optional.String**|  | 
 **limit** | **optional.Float64**|  | 
 **direction** | **optional.String**|  | 
 **protectionType** | **optional.String**|  | [default to ihp]

### Return type

[**SecurityEventsResponse**](SecurityEventsResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomerPrototypeGetLocationSecurityPolicyHourlyCounts**
> Counts CustomerPrototypeGetLocationSecurityPolicyHourlyCounts(ctx, id, locationId)
Get a Security Policy Hourly Blocked Counts for a Location ID.

<div><strong>200</strong>: Success.</div> <div><strong>401</strong>: Authorization required or customer id not found</div> <div><strong>404</strong>: Location id or WifiNetwork does not exist and is not known to Plume</div> <div><strong>500</strong>: Internal server error</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Customer id | 
  **locationId** | **string**|  | 

### Return type

[**Counts**](Counts.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomerPrototypeGetLocationWanConfiguration**
> interface{} CustomerPrototypeGetLocationWanConfiguration(ctx, id, locationId)
Get WAN Configuration for a Location ID.

<div><strong>200</strong>: Success, WAN Settings returned.</div> <div><strong>401</strong>: Authorization required or customer id not found.</div> <div><strong>404</strong>: Location id, does not exist.</div> <div><strong>500</strong>: Internal server error</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Customer id | 
  **locationId** | **string**|  | 

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomerPrototypeGetLocationWanSettings**
> LocationWanSettings CustomerPrototypeGetLocationWanSettings(ctx, id, locationId)
DEPRECATED: Get the WAN Settings for a Location ID.

<div><strong>200</strong>: Success, WAN Settings returned.</div> <div><strong>401</strong>: Authorization required or customer id not found.</div> <div><strong>404</strong>: Location id, does not exist.</div> <div><strong>500</strong>: Internal server error</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Customer id | 
  **locationId** | **string**|  | 

### Return type

[**LocationWanSettings**](LocationWanSettings.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomerPrototypeGetLocations**
> Location CustomerPrototypeGetLocations(ctx, id, optional)
Queries locations of Customer.

<div><strong>200</strong>: Success, full object returned.</div> <div><strong>401</strong>: Authorization required.</div> <div><strong>404</strong>: LocationId not found.</div> <div><strong>500</strong>: Internal server error.</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Customer id | 
 **optional** | ***CustomerApiCustomerPrototypeGetLocationsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CustomerApiCustomerPrototypeGetLocationsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **include** | **optional.String**| CSV value of objects to add to the response: summary (is the only option for now) | 

### Return type

[**Location**](Location.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomerPrototypeGetManagersListForLocation**
> interface{} CustomerPrototypeGetManagersListForLocation(ctx, id, locationId)
Get a list of all managers the are assigned to manage your location.

<div><strong>200</strong>: Success.</div> <div><strong>404</strong>: Location does not exist.</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Customer id | 
  **locationId** | **string**|  | 

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomerPrototypeGetMigration**
> Migration CustomerPrototypeGetMigration(ctx, id, optional)
Fetches hasOne relation _migration.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Customer id | 
 **optional** | ***CustomerApiCustomerPrototypeGetMigrationOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CustomerApiCustomerPrototypeGetMigrationOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **refresh** | **optional.Bool**|  | 

### Return type

[**Migration**](Migration.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomerPrototypeGetModulesFromController**
> interface{} CustomerPrototypeGetModulesFromController(ctx, id, locationId)
Retrieve all firmaware modules for a Location ID.

<div><strong>200</strong>: Success, your new info looks good.</div> <div><strong>401</strong>: Authorization required or customer id not found</div> <div><strong>404</strong>: location id does not exist.</div> <div><strong>500</strong>: internal server error</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Customer id | 
  **locationId** | **string**|  | 

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomerPrototypeGetMotionHistory**
> []interface{} CustomerPrototypeGetMotionHistory(ctx, id, locationId, optional)
Fetch the motion density history for this location

<div><strong>200</strong>: Success, motion density array returned.</div> <div><strong>401</strong>: Authorization required or customer id not found</div> <div><strong>404</strong>: customer id or location id does not exist and is not known to Plume</div> <div><strong>500</strong>: internal server error</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Customer id | 
  **locationId** | **string**|  | 
 **optional** | ***CustomerApiCustomerPrototypeGetMotionHistoryOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CustomerApiCustomerPrototypeGetMotionHistoryOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **from** | **optional.Float64**| UTC unix ts | 
 **to** | **optional.Float64**| UTC unix ts, defaults to now | 
 **bucket** | **optional.Float64**| number of seconds in density calculation window; returned data points represent % of non-zero intensity values in the window | [default to 3600]

### Return type

**[]interface{}**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomerPrototypeGetMotionStateHistory**
> []interface{} CustomerPrototypeGetMotionStateHistory(ctx, id, locationId, optional)
Fetch the motion state history for this location

<div><strong>200</strong>: Success, motion state array returned (Each element of the array is in the form [\"val\", \"unix_ts\"], where \"val\" is one of:  <div>0 - Not armed, not tripped</div> <div>1 - Not armed, tripped</div> <div>2 - Armed, not tripped</div> <div>3 - Armed, tripped</div></div> <div><strong>401</strong>: Authorization required or customer id not found</div> <div><strong>404</strong>: customer id or location id does not exist and is not known to Plume</div> <div><strong>500</strong>: internal server error</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Customer id | 
  **locationId** | **string**|  | 
 **optional** | ***CustomerApiCustomerPrototypeGetMotionStateHistoryOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CustomerApiCustomerPrototypeGetMotionStateHistoryOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **from** | **optional.Float64**| UTC unix ts | 
 **to** | **optional.Float64**| UTC unix ts, defaults to now | 
 **bucket** | **optional.Float64**| number of seconds in density calculation window; returned data points represent % of non-zero intensity values in the window | [default to 3600]

### Return type

**[]interface{}**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomerPrototypeGetNetworkAccessNetworks**
> []NetworkAccessNetwork CustomerPrototypeGetNetworkAccessNetworks(ctx, id, locationId)
Get information about networkAccess networks

<div><strong>204</strong>: Success.</div> <div><strong>404</strong>: Location does not exist.</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Customer id | 
  **locationId** | **string**|  | 

### Return type

[**[]NetworkAccessNetwork**](NetworkAccessNetwork.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomerPrototypeGetNetworkConfigurationHome**
> NetworkConfigurationHome CustomerPrototypeGetNetworkConfigurationHome(ctx, id, locationId)
Get the current overall settings and status of the Advanced Networking settings for a Location ID.

<div><strong>200</strong>: Success, current networkConfiguration settings returned.</div> <div><strong>401</strong>: Authorization required or customer id not found</div> <div><strong>500</strong>: internal server error</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Customer id | 
  **locationId** | **string**|  | 

### Return type

[**NetworkConfigurationHome**](NetworkConfigurationHome.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomerPrototypeGetNetworkMode**
> InlineResponse2005 CustomerPrototypeGetNetworkMode(ctx, id, locationId)
Get the current Network Mode for a Location ID.

<div><strong>200</strong>: Success, current NetworkMode returned.</div> <div><strong>401</strong>: Authorization required or customer id not found</div> <div><strong>404</strong>: location id or NetworkMode does not exist.</div> <div><strong>500</strong>: internal server error</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Customer id | 
  **locationId** | **string**|  | 

### Return type

[**InlineResponse2005**](inline_response_200_5.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomerPrototypeGetNodeBlePairingPin**
> interface{} CustomerPrototypeGetNodeBlePairingPin(ctx, id, locationId, nodeId, token)
Get BLE pairing pin for a node that is claimed by the selected location

<div><strong>200</strong>: Success, pin generated.</div> <div><strong>400</strong>: Required fields missing or field type is incorrect.</div> <div><strong>404</strong>: Location or node does not exist.</div> <div><strong>422</strong>: Invalid token. </div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Customer id | 
  **locationId** | **string**|  | 
  **nodeId** | **string**|  | 
  **token** | **string**|  | 

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomerPrototypeGetNodeBySerialNumber**
> NodeResponse CustomerPrototypeGetNodeBySerialNumber(ctx, id, locationId, nodeId)
Returns a single Node for a Location ID with its list of connected devices.

<div><strong>200</strong>: Success, node returned.</div> <div><strong>404</strong>: customer id or location id does not exist. Or, nodeId not claimed to this account.</div> <div><strong>500</strong>: internal server error</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Customer id | 
  **locationId** | **string**|  | 
  **nodeId** | **string**| id of node | 

### Return type

[**NodeResponse**](NodeResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomerPrototypeGetNodeKvStates**
> interface{} CustomerPrototypeGetNodeKvStates(ctx, id, locationId, nodeId)
Retrieve all kvStates on a particular Node for a Location ID.

<div><strong>200</strong>: Success, your new info looks good.</div> <div><strong>401</strong>: Authorization required or customer id not found</div> <div><strong>404</strong>: location id does not exist.</div> <div><strong>422</strong>: nodeId must be defined.</div> <div><strong>425</strong>: nodeId must belong to the location.</div> <div><strong>500</strong>: internal server error</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Customer id | 
  **locationId** | **string**|  | 
  **nodeId** | **string**|  | 

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomerPrototypeGetNodes**
> InlineResponse20010 CustomerPrototypeGetNodes(ctx, id, locationId)
Retrieve the Node settings and status for a Location ID.

<div><strong>200</strong>: Success, array of Nodes returned.</div> <div><strong>404</strong>: customer id or location id does not exist and is not known to Plume</div> <div><strong>500</strong>: internal server error</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Customer id | 
  **locationId** | **string**|  | 

### Return type

[**InlineResponse20010**](inline_response_200_10.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomerPrototypeGetOhpLocationIdentifier**
> interface{} CustomerPrototypeGetOhpLocationIdentifier(ctx, id, locationId)
Get the current OHP identifier for a Location ID.

<div><strong>200</strong>: Success, current DPP configurator returned.</div> <div><strong>401</strong>: Authorization required or customer id not found</div> <div><strong>500</strong>: internal server error</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Customer id | 
  **locationId** | **string**|  | 

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomerPrototypeGetOnboardingLocationIdentifier**
> interface{} CustomerPrototypeGetOnboardingLocationIdentifier(ctx, id, locationId)
Get the onboarding identifier for a Location ID.

<div><strong>200</strong>: Success, current DPP configurator returned.</div> <div><strong>401</strong>: Authorization required or customer id not found</div> <div><strong>500</strong>: internal server error</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Customer id | 
  **locationId** | **string**|  | 

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomerPrototypeGetPersonAppTimeAppsDataUsage**
> []interface{} CustomerPrototypeGetPersonAppTimeAppsDataUsage(ctx, id, locationId, personId, timePeriod, optional)
Fetch the AppTime Apps Data Usage Stats for a Person.

<div><strong>200</strong>: Success, AppTime Stats returned.</div> <div><strong>401</strong>: Authorization required or customer id not found</div> <div><strong>404</strong>: location id or person does not exist.</div> <div><strong>500</strong>: internal server error</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Customer id | 
  **locationId** | **string**|  | 
  **personId** | **string**|  | 
  **timePeriod** | **string**| Any of \&quot;weekly\&quot;,\&quot;dailyToday\&quot;,\&quot;dailyYesterday\&quot;,\&quot;daily2DaysAgo\&quot;,\&quot;daily3DaysAgo\&quot;,\&quot;daily4DaysAgo\&quot;,\&quot;daily5DaysAgo\&quot;,\&quot;daily6DaysAgo\&quot;,\&quot;last30Days\&quot;,\&quot;last12Months\&quot; | 
 **optional** | ***CustomerApiCustomerPrototypeGetPersonAppTimeAppsDataUsageOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CustomerApiCustomerPrototypeGetPersonAppTimeAppsDataUsageOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **limit** | **optional.Float64**| Maximum number of apps to return. Defaults to 20 | 
 **grouping** | **optional.String**| typing of Grouping for the purposes of applying the limit. Can be: &#39;perTimeSlot&#39; ONLY | [default to perTimeSlot]

### Return type

**[]interface{}**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomerPrototypeGetPersonAppTimeAppsOnlineTime**
> []interface{} CustomerPrototypeGetPersonAppTimeAppsOnlineTime(ctx, id, locationId, personId, timePeriod, optional)
Fetch the AppTime Apps Online Time Stats for a Person.

<div><strong>200</strong>: Success, AppTime Stats returned.</div> <div><strong>401</strong>: Authorization required or customer id not found</div> <div><strong>404</strong>: location id or person does not exist.</div> <div><strong>500</strong>: internal server error</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Customer id | 
  **locationId** | **string**|  | 
  **personId** | **string**|  | 
  **timePeriod** | **string**| Any of \&quot;weekly\&quot;,\&quot;dailyToday\&quot;,\&quot;dailyYesterday\&quot;,\&quot;daily2DaysAgo\&quot;,\&quot;daily3DaysAgo\&quot;,\&quot;daily4DaysAgo\&quot;,\&quot;daily5DaysAgo\&quot;,\&quot;daily6DaysAgo\&quot;,\&quot;last30Days\&quot;,\&quot;last12Months\&quot; | 
 **optional** | ***CustomerApiCustomerPrototypeGetPersonAppTimeAppsOnlineTimeOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CustomerApiCustomerPrototypeGetPersonAppTimeAppsOnlineTimeOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **limit** | **optional.Float64**| Maximum number of apps to return. Defaults to 20 | 
 **grouping** | **optional.String**| typing of Grouping for the purposes of applying the limit. Can be: &#39;perTimeSlot&#39; ONLY | [default to perTimeSlot]

### Return type

**[]interface{}**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomerPrototypeGetPersonAppTimeCategoriesDataUsage**
> []interface{} CustomerPrototypeGetPersonAppTimeCategoriesDataUsage(ctx, id, locationId, personId, timePeriod, optional)
Fetch the AppTime Categories Data Usage Stats for a Person.

<div><strong>200</strong>: Success, AppTime Stats returned.</div> <div><strong>401</strong>: Authorization required or customer id not found</div> <div><strong>404</strong>: location id or person does not exist.</div> <div><strong>500</strong>: internal server error</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Customer id | 
  **locationId** | **string**|  | 
  **personId** | **string**|  | 
  **timePeriod** | **string**| Any of \&quot;weekly\&quot;,\&quot;dailyToday\&quot;,\&quot;dailyYesterday\&quot;,\&quot;daily2DaysAgo\&quot;,\&quot;daily3DaysAgo\&quot;,\&quot;daily4DaysAgo\&quot;,\&quot;daily5DaysAgo\&quot;,\&quot;daily6DaysAgo\&quot;,\&quot;last30Days\&quot;,\&quot;last12Months\&quot; | 
 **optional** | ***CustomerApiCustomerPrototypeGetPersonAppTimeCategoriesDataUsageOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CustomerApiCustomerPrototypeGetPersonAppTimeCategoriesDataUsageOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **limit** | **optional.Float64**| Maximum number of categories to return. Defaults to 20 | 
 **grouping** | **optional.String**| typing of Grouping for the purposes of applying the limit. Can be: &#39;overall&#39;|&#39;perTimeSlot&#39; | [default to perTimeSlot]

### Return type

**[]interface{}**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomerPrototypeGetPersonAppTimeCategoriesOnlineTime**
> []interface{} CustomerPrototypeGetPersonAppTimeCategoriesOnlineTime(ctx, id, locationId, personId, timePeriod, optional)
Fetch the AppTime Categories Online Time Stats for a Person.

<div><strong>200</strong>: Success, AppTime Stats returned.</div> <div><strong>401</strong>: Authorization required or customer id not found</div> <div><strong>404</strong>: location id or person does not exist.</div> <div><strong>500</strong>: internal server error</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Customer id | 
  **locationId** | **string**|  | 
  **personId** | **string**|  | 
  **timePeriod** | **string**| Any of \&quot;weekly\&quot;,\&quot;dailyToday\&quot;,\&quot;dailyYesterday\&quot;,\&quot;daily2DaysAgo\&quot;,\&quot;daily3DaysAgo\&quot;,\&quot;daily4DaysAgo\&quot;,\&quot;daily5DaysAgo\&quot;,\&quot;daily6DaysAgo\&quot;,\&quot;last30Days\&quot;,\&quot;last12Months\&quot; | 
 **optional** | ***CustomerApiCustomerPrototypeGetPersonAppTimeCategoriesOnlineTimeOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CustomerApiCustomerPrototypeGetPersonAppTimeCategoriesOnlineTimeOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **limit** | **optional.Float64**| Maximum number of categories to return. Defaults to 20 | 
 **grouping** | **optional.String**| typing of Grouping for the purposes of applying the limit. Can be: &#39;overall&#39;|&#39;perTimeSlot&#39; | [default to perTimeSlot]

### Return type

**[]interface{}**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomerPrototypeGetPersonById**
> Person CustomerPrototypeGetPersonById(ctx, id, locationId, personId)
Get a Person by ID for a Location ID.

<div><strong>200</strong>: Success.</div> <div><strong>401</strong>: Authorization required or customer id not found</div> <div><strong>404</strong>: Location id does not exist and is not known to Plume</div> <div><strong>500</strong>: Internal server error</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Customer id | 
  **locationId** | **string**|  | 
  **personId** | **string**|  | 

### Return type

[**Person**](Person.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomerPrototypeGetPersonSecurityPolicyEvents**
> SecurityEventsResponse CustomerPrototypeGetPersonSecurityPolicyEvents(ctx, id, locationId, personId, startTime, optional)
Get a Security Policy Events for Person for a Location ID.

<div><strong>200</strong>: Success.</div> <div><strong>401</strong>: Authorization required or customer id not found</div> <div><strong>404</strong>: Location id or WifiNetwork does not exist and is not known to Plume</div> <div><strong>500</strong>: Internal server error</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Customer id | 
  **locationId** | **string**|  | 
  **personId** | **string**| person | 
  **startTime** | **time.Time**|  | 
 **optional** | ***CustomerApiCustomerPrototypeGetPersonSecurityPolicyEventsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CustomerApiCustomerPrototypeGetPersonSecurityPolicyEventsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **includes** | **optional.String**|  | 
 **limit** | **optional.Float64**|  | 
 **direction** | **optional.String**|  | 
 **protectionType** | **optional.String**|  | [default to ihp]

### Return type

[**SecurityEventsResponse**](SecurityEventsResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomerPrototypeGetPersonSecurityPolicyHourlyCounts**
> Counts CustomerPrototypeGetPersonSecurityPolicyHourlyCounts(ctx, id, locationId, personId)
Get a Security Policy Hourly Blocked Counts for a Person for a Location ID.

<div><strong>200</strong>: Success.</div> <div><strong>401</strong>: Authorization required or customer id not found</div> <div><strong>404</strong>: Location id or WifiNetwork does not exist and is not known to Plume</div> <div><strong>500</strong>: Internal server error</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Customer id | 
  **locationId** | **string**|  | 
  **personId** | **string**| person | 

### Return type

[**Counts**](Counts.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomerPrototypeGetPersons**
> []Person CustomerPrototypeGetPersons(ctx, id, locationId)
Get all Persons for a Location ID.

<div><strong>200</strong>: Success.</div> <div><strong>401</strong>: Authorization required or customer id not found</div> <div><strong>404</strong>: Location id does not exist and is not known to Plume</div> <div><strong>500</strong>: Internal server error</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Customer id | 
  **locationId** | **string**|  | 

### Return type

[**[]Person**](Person.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomerPrototypeGetPortForwards**
> []PortForward CustomerPrototypeGetPortForwards(ctx, id, locationId, mac)
Get all existing Port Forwarding entries for an existing DHCP IP reservation tied to a MAC address at a Location ID.

<div><strong>200</strong>: Success, current Port Forwarding entries  returned.</div> <div><strong>401</strong>: Authorization required or customer id not found</div> <div><strong>404</strong>: NetworkConfiguration or dhcpReservations value is empty.</div> <div><strong>422</strong>: mac is empty or invalid.</div> <div><strong>500</strong>: internal server error</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Customer id | 
  **locationId** | **string**|  | 
  **mac** | **string**|  | 

### Return type

[**[]PortForward**](PortForward.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomerPrototypeGetPrimarySecondaryNetworks**
> XAny CustomerPrototypeGetPrimarySecondaryNetworks(ctx, id, locationId)
Get networks for wpa3 transition flow

<div><strong>200</strong>: Success, returns the data</div> <div><strong>401</strong>: Authorization required or customer id not found</div> <div><strong>404</strong>: location id or wifiNetwork does not exist</div> <div><strong>500</strong>: Internal server error</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Customer id | 
  **locationId** | **string**|  | 

### Return type

[**XAny**](x-any.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomerPrototypeGetQoe1Minute**
> interface{} CustomerPrototypeGetQoe1Minute(ctx, id, locationId, mac, optional)
Device or pod QoE live mode data.

<div><strong>200</strong>: Success.</div> <div><strong>401</strong>: Authorization required.</div> <div><strong>404</strong>: Location id does not exist.</div> <div><strong>500</strong>: Internal server error</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Customer id | 
  **locationId** | **string**|  | 
  **mac** | **string**| mac address or pod id | 
 **optional** | ***CustomerApiCustomerPrototypeGetQoe1MinuteOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CustomerApiCustomerPrototypeGetQoe1MinuteOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **startTime** | **optional.Float64**| start timestamp | 
 **timestampISOFormat** | **optional.Bool**| either timestamp utc number or ISO string | [default to false]

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomerPrototypeGetQoeSeconds**
> interface{} CustomerPrototypeGetQoeSeconds(ctx, id, locationId, mac, optional)
Device or pod QoE super live mode data.

<div><strong>200</strong>: Success.</div> <div><strong>401</strong>: Authorization required.</div> <div><strong>404</strong>: Location id does not exist.</div> <div><strong>500</strong>: Internal server error</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Customer id | 
  **locationId** | **string**|  | 
  **mac** | **string**| mac address or pod id | 
 **optional** | ***CustomerApiCustomerPrototypeGetQoeSecondsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CustomerApiCustomerPrototypeGetQoeSecondsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **startTime** | **optional.Float64**| start timestamp | 
 **timestampISOFormat** | **optional.Bool**| either timestamp utc number or ISO string | [default to false]

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomerPrototypeGetRemoteConnectionsConfig**
> interface{} CustomerPrototypeGetRemoteConnectionsConfig(ctx, id, locationId)
Get the Unauthorized Remote Connections config for a Location ID.

<div><strong>200</strong>: Success.</div> <div><strong>401</strong>: Authorization required or customer id not found</div> <div><strong>404</strong>: Location id does not exist and is not known to Plume</div> <div><strong>500</strong>: Internal server error</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Customer id | 
  **locationId** | **string**|  | 

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomerPrototypeGetRoles**
> []Role CustomerPrototypeGetRoles(ctx, id, optional)
Queries roles of Customer.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Customer id | 
 **optional** | ***CustomerApiCustomerPrototypeGetRolesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CustomerApiCustomerPrototypeGetRolesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **filter** | **optional.String**|  | 

### Return type

[**[]Role**](Role.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomerPrototypeGetSecondaryNetworkInvitation**
> Invitation CustomerPrototypeGetSecondaryNetworkInvitation(ctx, id, locationId)
Update home devices visible to guests.

<div><strong>200</strong>: Success, Invitation returned.</div> <div><strong>401</strong>: Authorization required or customer id not found.</div> <div><strong>404</strong>: Customer id, location id</div> <div><strong>500</strong>: Internal server error</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Customer id | 
  **locationId** | **string**|  | 

### Return type

[**Invitation**](Invitation.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomerPrototypeGetSecurityRealizedStates**
> interface{} CustomerPrototypeGetSecurityRealizedStates(ctx, id, locationId)
Retrieve all securityStates for a Location ID.

<div><strong>200</strong>: Success, your new info looks good.</div> <div><strong>401</strong>: Authorization required or customer id not found</div> <div><strong>404</strong>: location id does not exist.</div> <div><strong>500</strong>: internal server error</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Customer id | 
  **locationId** | **string**|  | 

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomerPrototypeGetServiceLevel**
> interface{} CustomerPrototypeGetServiceLevel(ctx, id, locationId)
Get the service level for this location

<div><strong>200</strong>: Success, return service Level object.</div> <div><strong>401</strong>: Authorization required or customer id not found</div> <div><strong>404</strong>: customer id or location id does not exist and is not known to Plume</div> <div><strong>500</strong>: internal server error</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Customer id | 
  **locationId** | **string**|  | 

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomerPrototypeGetSniffing**
> interface{} CustomerPrototypeGetSniffing(ctx, id, locationId)
Get DNS, HTTP, UPnP and mDNS sniffing toggles for a Location ID.

<div><strong>200</strong>: Success, current sniffing toggles returned.</div> <div><strong>401</strong>: Authorization required or customer id not found</div> <div><strong>500</strong>: internal server error</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Customer id | 
  **locationId** | **string**|  | 

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomerPrototypeGetSpeedTestResults**
> interface{} CustomerPrototypeGetSpeedTestResults(ctx, id, locationId, nodeId, optional)
retrieve the speed test result for a node.

<div><strong>200</strong>: Success, run.</div> <div><strong>422</strong>: locationId or nodeId isn't defined.</div> <div><strong>500</strong>: Internal server error</div> <div><strong>503</strong>: Service Unavailable.</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Customer id | 
  **locationId** | **string**|  | 
  **nodeId** | **string**|  | 
 **optional** | ***CustomerApiCustomerPrototypeGetSpeedTestResultsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CustomerApiCustomerPrototypeGetSpeedTestResultsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **granularity** | **optional.String**| days/hours/minutes | [default to days]
 **limit** | **optional.Float64**| X # of days/hours/minutes | [default to 7]

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomerPrototypeGetSpeedTestResultsByRequestId**
> interface{} CustomerPrototypeGetSpeedTestResultsByRequestId(ctx, id, locationId, nodeId, requestId)
retrieve single speed test result by request id for a node.

<div><strong>200</strong>: Success.</div> <div><strong>422</strong>: locationId or nodeId isn't defined.</div> <div><strong>404</strong>: Speed test not found.</div> <div><strong>500</strong>: Internal server error</div> <div><strong>503</strong>: Service Unavailable.</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Customer id | 
  **locationId** | **string**|  | 
  **nodeId** | **string**|  | 
  **requestId** | **string**|  | 

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomerPrototypeGetSpeedTestResultsForApp**
> XAny CustomerPrototypeGetSpeedTestResultsForApp(ctx, id, locationId, optional)
Get the current speed test aggregation result for a Location ID.

<div><strong>200</strong>: Success, current speedTest result and most active devices returned.</div> <div><strong>401</strong>: Authorization required or customer id not found</div> <div><strong>404</strong>: location id or NetworkMode does not exist.</div> <div><strong>500</strong>: internal server error</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Customer id | 
  **locationId** | **string**|  | 
 **optional** | ***CustomerApiCustomerPrototypeGetSpeedTestResultsForAppOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CustomerApiCustomerPrototypeGetSpeedTestResultsForAppOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **excludeDevices** | **optional.Bool**|  | 

### Return type

[**XAny**](x-any.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomerPrototypeGetSsid**
> InlineResponse2007 CustomerPrototypeGetSsid(ctx, id, locationId)
Get the current WiFi SSID for a Location ID.

<div><strong>200</strong>: Success, current Wifi Network returned.</div> <div><strong>401</strong>: Authorization required or customer id not found</div> <div><strong>404</strong>: location id or WifiNetwork does not exist.</div> <div><strong>500</strong>: internal server error</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Customer id | 
  **locationId** | **string**|  | 

### Return type

[**InlineResponse2007**](inline_response_200_7.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomerPrototypeGetSubscription**
> interface{} CustomerPrototypeGetSubscription(ctx, id, locationId)
Get Subscription details for this location

<div><strong>200</strong>: Success, subscription details returned</div> <div><strong>401</strong>: Authorization required or customer id not found</div> <div><strong>404</strong>: customer id or location id does not exist and is not known to Plume</div> <div><strong>500</strong>: internal server error</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Customer id | 
  **locationId** | **string**|  | 

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomerPrototypeGetSummary**
> SummaryResponse CustomerPrototypeGetSummary(ctx, id, locationId)
DEPRECATED: The system summary for a location including topology, optimizations, and firmware upgrades.

<div><strong>200</strong>: Success, system info plus topology array returned.</div> <div><strong>404</strong>: customer id or location id does not exist and is not known to Plume</div> <div><strong>500</strong>: internal server error</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Customer id | 
  **locationId** | **string**|  | 

### Return type

[**SummaryResponse**](SummaryResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomerPrototypeGetTaskStatuses**
> interface{} CustomerPrototypeGetTaskStatuses(ctx, id, locationId)
Retrieve all task statuses of nodes from controller

<div><strong>204</strong>: Success.</div> <div><strong>404</strong>: customer id or location id does not exist.</div> <div><strong>500</strong>: internal server error</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Customer id | 
  **locationId** | **string**|  | 

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomerPrototypeGetTermsAndPrivacyAccepted**
> TermsAndPrivacy CustomerPrototypeGetTermsAndPrivacyAccepted(ctx, id, optional)
Fetches hasOne relation termsAndPrivacyAccepted.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Customer id | 
 **optional** | ***CustomerApiCustomerPrototypeGetTermsAndPrivacyAcceptedOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CustomerApiCustomerPrototypeGetTermsAndPrivacyAcceptedOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **refresh** | **optional.Bool**|  | 

### Return type

[**TermsAndPrivacy**](TermsAndPrivacy.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomerPrototypeGetTopology**
> TopologyResponse CustomerPrototypeGetTopology(ctx, id, locationId)
DEPRECATED: The topology for a location including channels and devices.

Please use the GET /Customers/{id}/locations/{locationId}/forceGraph API as a replacement. <div><strong>200</strong>: Success, array of Nodes returned.</div> <div><strong>404</strong>: customer id, location id, or topology does not exist and is not known to Plume</div> <div><strong>500</strong>: internal server error</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Customer id | 
  **locationId** | **string**|  | 

### Return type

[**TopologyResponse**](TopologyResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomerPrototypeGetTos**
> XAny CustomerPrototypeGetTos(ctx, id, locationId, mac)
Describes the current state of TOS for the given client.

<div><strong>200</strong>: Ok.</div> <div><strong>401</strong>: Authorization required.</div> <div><strong>404</strong>: Location id, does not exist.</div> <div><strong>404</strong>: No device found with provided mac address</div> <div><strong>422</strong>: Invalid MAC.</div> <div><strong>500</strong>: Internal server error.</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Customer id | 
  **locationId** | **string**|  | 
  **mac** | **string**|  | 

### Return type

[**XAny**](x-any.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomerPrototypeGetUpnp**
> Upnp CustomerPrototypeGetUpnp(ctx, id, locationId)
Get the current UPnP setting for a Location ID.

<div><strong>200</strong>: Success, current Upnp returned.</div> <div><strong>401</strong>: Authorization required or customer id not found</div> <div><strong>500</strong>: internal server error</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Customer id | 
  **locationId** | **string**|  | 

### Return type

[**Upnp**](Upnp.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomerPrototypeGetVapStates**
> interface{} CustomerPrototypeGetVapStates(ctx, id, locationId)
Retrieve all Vap State on a particular Node for a Location ID.

<div><strong>200</strong>: Success, your new info looks good.</div> <div><strong>401</strong>: Authorization required or customer id not found</div> <div><strong>404</strong>: location id does not exist.</div> <div><strong>500</strong>: internal server error</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Customer id | 
  **locationId** | **string**|  | 

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomerPrototypeGetVapsAndStaStatesFromBackhaul**
> interface{} CustomerPrototypeGetVapsAndStaStatesFromBackhaul(ctx, id, locationId)
Retrieve all Vap State on a particular Node for a Location ID.

<div><strong>200</strong>: Success, your new info looks good.</div> <div><strong>401</strong>: Authorization required or customer id not found</div> <div><strong>404</strong>: location id does not exist.</div> <div><strong>500</strong>: internal server error</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Customer id | 
  **locationId** | **string**|  | 

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomerPrototypeGetWhitelistApprovalRequests**
> interface{} CustomerPrototypeGetWhitelistApprovalRequests(ctx, id, locationId)
Get a list of pending approval requests for this location.

<div><strong>200</strong>: Success.</div> <div><strong>400</strong>: Required fields missing or field type is incorrect.</div> <div><strong>401</strong>: Authorization required or customer id not found</div> <div><strong>404</strong>: Location id, CustomerId or requst id does not exist and is not known to Plume</div> <div><strong>500</strong>: Internal server error</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Customer id | 
  **locationId** | **string**|  | 

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomerPrototypeGetWifiDashboard**
> WifiDashboardResponse CustomerPrototypeGetWifiDashboard(ctx, id, locationId)
WiFi Dashboard

<div><strong>200</strong>: Success, response object returned.</div> <div><strong>401</strong>: Authorization required or customer id not found</div> <div><strong>404</strong>: customer id or location id does not exist and is not known to Plume</div> <div><strong>500</strong>: internal server error</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Customer id | 
  **locationId** | **string**|  | 

### Return type

[**WifiDashboardResponse**](WifiDashboardResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomerPrototypeGetWifiInvitationById**
> Invitation CustomerPrototypeGetWifiInvitationById(ctx, id, locationId, zoneId, keyId)
Update home devices visible to guests.

<div><strong>200</strong>: Success, Invitation returned.</div> <div><strong>400</strong>: Required fields missing.</div> <div><strong>401</strong>: Authorization required or customer id not found.</div> <div><strong>404</strong>: Customer id, location id, or WifiNetwork accessZone zoneId/keyId does not exist.</div> <div><strong>500</strong>: Internal server error</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Customer id | 
  **locationId** | **string**|  | 
  **zoneId** | **string**|  | 
  **keyId** | **float64**| keys id be added | 

### Return type

[**Invitation**](Invitation.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomerPrototypeGetWifiMotion**
> WifiMotion CustomerPrototypeGetWifiMotion(ctx, id, locationId)
Get WifiMotion config for this location

<div><strong>200</strong>: Success, wifiMotion object returned.</div> <div><strong>401</strong>: Authorization required or customer id not found</div> <div><strong>404</strong>: customer id or location id does not exist and is not known to Plume</div> <div><strong>500</strong>: internal server error</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Customer id | 
  **locationId** | **string**|  | 

### Return type

[**WifiMotion**](WifiMotion.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomerPrototypeGetWifiNetwork**
> InlineResponse2002 CustomerPrototypeGetWifiNetwork(ctx, id, locationId)
Get the current WiFi SSID and PSK for a Location ID.

<div><strong>200</strong>: Success, current Wifi Network returned.</div> <div><strong>401</strong>: Authorization required or customer id not found</div> <div><strong>404</strong>: location id or WifiNetwork does not exist.</div> <div><strong>500</strong>: internal server error</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Customer id | 
  **locationId** | **string**|  | 

### Return type

[**InlineResponse2002**](inline_response_200_2.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomerPrototypeGetWifiNetworkDpp**
> interface{} CustomerPrototypeGetWifiNetworkDpp(ctx, id, locationId)
Get the current DPP configurator for a Location ID.

<div><strong>200</strong>: Success, current DPP configurator returned.</div> <div><strong>401</strong>: Authorization required or customer id not found</div> <div><strong>500</strong>: internal server error</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Customer id | 
  **locationId** | **string**|  | 

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomerPrototypeGetWifiNetworks**
> []interface{} CustomerPrototypeGetWifiNetworks(ctx, id, locationId)
WiFi Networks

<div><strong>200</strong>: Success, response object returned.</div> <div><strong>401</strong>: Authorization required or customer id not found</div> <div><strong>404</strong>: customer id or location id does not exist and is not known to Plume</div> <div><strong>500</strong>: internal server error</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Customer id | 
  **locationId** | **string**|  | 

### Return type

**[]interface{}**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomerPrototypeHasLocationById**
> Location CustomerPrototypeHasLocationById(ctx, id, locationId)
Verify that a Customer Id has a Location Id.

<div><strong>200</strong>: Success, no data returned.</div> <div><strong>401</strong>: Authorization required.</div> <div><strong>404</strong>: LocationId not found.</div> <div><strong>500</strong>: Internal server error.</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Customer id | 
  **locationId** | **string**|  | 

### Return type

[**Location**](Location.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomerPrototypeIosDeviceTokenExists**
> interface{} CustomerPrototypeIosDeviceTokenExists(ctx, id, deviceToken)
Provides feedback as to whether an iOS deviceToken was previously registered for push notifications.

<div><strong>200</strong>: Success, exists:true|false returned.</div> <div><strong>404</strong>: customer id does not exist.</div> <div><strong>500</strong>: internal server error</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Customer id | 
  **deviceToken** | **string**|  | 

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomerPrototypeLinkAccount**
> interface{} CustomerPrototypeLinkAccount(ctx, id, provider, userId, sessionToken, optional)
link the outside account, such as Samsung user.

<div><strong>200</strong>: Success, the outside account inserted into the customer info/object.</div> <div><strong>401</strong>: Authorization required or customer id not found</div> <div><strong>500</strong>: internal server error</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Customer id | 
  **provider** | **string**|  | 
  **userId** | **string**|  | 
  **sessionToken** | **string**|  | 
 **optional** | ***CustomerApiCustomerPrototypeLinkAccountOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CustomerApiCustomerPrototypeLinkAccountOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **userName** | **optional.String**|  | 
 **userDisplayName** | **optional.String**|  | 

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomerPrototypeListCustomSharedSchedules**
> []LocationCustomSchedule CustomerPrototypeListCustomSharedSchedules(ctx, id, locationId)
Get custom shared schedules for a given Location ID.

<div><strong>200</strong>: Success, custom schedules list returned.</div> <div><strong>401</strong>: Authorization required or customer id not found</div> <div><strong>404</strong>: location id does not exist.</div> <div><strong>500</strong>: internal server error</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Customer id | 
  **locationId** | **string**|  | 

### Return type

[**[]LocationCustomSchedule**](LocationCustomSchedule.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomerPrototypeMarketingExport**
> interface{} CustomerPrototypeMarketingExport(ctx, id, locationId, optional)
Get detailed information of a location for CRM campaigns.

<div><strong>200</strong>: Success, location data in response.</div> <div><strong>500</strong>: internal server error</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Customer id | 
  **locationId** | **string**|  | 
 **optional** | ***CustomerApiCustomerPrototypeMarketingExportOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CustomerApiCustomerPrototypeMarketingExportOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **wifiMotionCapable** | **optional.Bool**|  | 
 **wifiMotionEnable** | **optional.Bool**|  | 
 **onlineProtectionEnabled** | **optional.Bool**|  | 
 **personsWithoutAssignedDevices** | **optional.Bool**|  | 
 **peopleProfileEverCreated** | **optional.Bool**|  | 
 **blockedSecurityEventsCountThirtyDay** | **optional.Bool**|  | 
 **devicesOnlineThirtyDays** | **optional.Bool**|  | 
 **mostActiveDevicesThirtyDays** | **optional.Bool**|  | 
 **appTimeCapable** | **optional.Bool**|  | 
 **subscription** | **optional.Bool**|  | 
 **lastThirtyDaysSpeedTestAverages** | **optional.Bool**|  | 

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomerPrototypeMigrationStatus**
> interface{} CustomerPrototypeMigrationStatus(ctx, id)
Returns cloud migration status for customer

<div><strong>200</strong>: Success, return the search result.</div> <div><strong>401</strong>: Authorization required or customer id not found</div> <div><strong>404</strong>: Customer does not exist.</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Customer id | 

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomerPrototypeOptimize**
> string CustomerPrototypeOptimize(ctx, id, locationId, optional)
Manually initiate an Optimize request for a Location ID.

<div><strong>200</strong>: Success, optimize request sent.</div> <div><strong>401</strong>: Authorization required or customer id not found.</div> <div><strong>404</strong>: location id, does not exist.</div> <div><strong>500</strong>: internal server error</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Customer id | 
  **locationId** | **string**|  | 
 **optional** | ***CustomerApiCustomerPrototypeOptimizeOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CustomerApiCustomerPrototypeOptimizeOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **forcePcs** | **optional.Bool**|  | 

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomerPrototypeOverlordDeleteAppQoeConfig**
> CustomerPrototypeOverlordDeleteAppQoeConfig(ctx, id, locationId)
Resets a appQoe config. AppQoe is to monitor the Quality of Experience of these Apps in the house, which is what this PRD covers. This QoE monitoring will allow CSPs understand likely issues with applications.

<div><strong>202</strong>: Success, reset.</div> <div><strong>401</strong>: Authorization required or customer id not found.</div> <div><strong>404</strong>: Location id, does not exist.</div> <div><strong>500</strong>: Internal server error</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Customer id | 
  **locationId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomerPrototypeOverlordDeleteFlowCacheConfig**
> CustomerPrototypeOverlordDeleteFlowCacheConfig(ctx, id, locationId)
Resets a flowCache config. Enable/disable Flow Cache to help support devQA to check influence on the first stage of the investigation.

<div><strong>202</strong>: Success, reset.</div> <div><strong>401</strong>: Authorization required or customer id not found.</div> <div><strong>404</strong>: Location id, does not exist.</div> <div><strong>500</strong>: Internal server error</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Customer id | 
  **locationId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomerPrototypeOverlordDeleteSamKnowsConfig**
> CustomerPrototypeOverlordDeleteSamKnowsConfig(ctx, id, locationId)
Resets a samKnows config. SamKnows is a provider of internet performance measurement services. They offer the SamKnows Router Agent, which supports a range of QoS and QoE performance measurements. These measurements can be executed both on an ad-hoc and scheduled basis.

<div><strong>202</strong>: Success, reset.</div> <div><strong>401</strong>: Authorization required or customer id not found.</div> <div><strong>404</strong>: Location id, does not exist.</div> <div><strong>500</strong>: Internal server error</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Customer id | 
  **locationId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomerPrototypeOverlordDeleteSipAlgConfig**
> CustomerPrototypeOverlordDeleteSipAlgConfig(ctx, id, locationId)
Resets a sipAlg config. sipAlg is an application within many routers. It inspects any VoIP traffic to prevent problems caused by firewalls and if necessary modifies the VoIP packets.

<div><strong>202</strong>: Success, reset.</div> <div><strong>401</strong>: Authorization required or customer id not found.</div> <div><strong>404</strong>: Location id, does not exist.</div> <div><strong>500</strong>: Internal server error</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Customer id | 
  **locationId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomerPrototypeOverlordDeleteStatsConfig**
> CustomerPrototypeOverlordDeleteStatsConfig(ctx, id, locationId)
Resets a stats config. Location Stats configuration, used to toggle which stats should be collected.

<div><strong>202</strong>: Success, reset.</div> <div><strong>401</strong>: Authorization required or customer id not found.</div> <div><strong>404</strong>: Location id, does not exist.</div> <div><strong>500</strong>: Internal server error</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Customer id | 
  **locationId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomerPrototypeOverlordUpdateAppQoeConfig**
> CustomerPrototypeOverlordUpdateAppQoeConfig(ctx, id, locationId, optional)
Updates a appQoe config. AppQoe is to monitor the Quality of Experience of these Apps in the house, which is what this PRD covers. This QoE monitoring will allow CSPs understand likely issues with applications.

<div><strong>202</strong>: Success, accepted and forwarded the data.</div> <div><strong>400</strong>: Required fields missing.</div> <div><strong>401</strong>: Authorization required or customer id not found.</div> <div><strong>404</strong>: Location id, does not exist.</div> <div><strong>422</strong>: Invalid data.</div> <div><strong>500</strong>: Internal server error</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Customer id | 
  **locationId** | **string**|  | 
 **optional** | ***CustomerApiCustomerPrototypeOverlordUpdateAppQoeConfigOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CustomerApiCustomerPrototypeOverlordUpdateAppQoeConfigOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **mode** | **optional.String**|  string enum: [ AUTO, ENABLE, DISABLE ] | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomerPrototypeOverlordUpdateFlowCacheConfig**
> CustomerPrototypeOverlordUpdateFlowCacheConfig(ctx, id, locationId, optional)
Updates a flowCache config. Enable/disable Flow Cache to help support devQA to check influence on the first stage of the investigation.

<div><strong>202</strong>: Success, accepted and forwarded the data.</div> <div><strong>400</strong>: Required fields missing.</div> <div><strong>401</strong>: Authorization required or customer id not found.</div> <div><strong>404</strong>: Location id, does not exist.</div> <div><strong>422</strong>: Invalid data.</div> <div><strong>500</strong>: Internal server error</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Customer id | 
  **locationId** | **string**|  | 
 **optional** | ***CustomerApiCustomerPrototypeOverlordUpdateFlowCacheConfigOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CustomerApiCustomerPrototypeOverlordUpdateFlowCacheConfigOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **enable** | **optional.Bool**|  boolean | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomerPrototypeOverlordUpdateSamKnowsConfig**
> CustomerPrototypeOverlordUpdateSamKnowsConfig(ctx, id, locationId, optional)
Updates a samKnows config. SamKnows is a provider of internet performance measurement services. They offer the SamKnows Router Agent, which supports a range of QoS and QoE performance measurements. These measurements can be executed both on an ad-hoc and scheduled basis.

<div><strong>202</strong>: Success, accepted and forwarded the data.</div> <div><strong>400</strong>: Required fields missing.</div> <div><strong>401</strong>: Authorization required or customer id not found.</div> <div><strong>404</strong>: Location id, does not exist.</div> <div><strong>422</strong>: Invalid data.</div> <div><strong>500</strong>: Internal server error</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Customer id | 
  **locationId** | **string**|  | 
 **optional** | ***CustomerApiCustomerPrototypeOverlordUpdateSamKnowsConfigOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CustomerApiCustomerPrototypeOverlordUpdateSamKnowsConfigOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **mode** | **optional.String**|  string enum: [ AUTO, ENABLE, DISABLE ] | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomerPrototypeOverlordUpdateSipAlgConfig**
> CustomerPrototypeOverlordUpdateSipAlgConfig(ctx, id, locationId, optional)
Updates a sipAlg config. sipAlg is an application within many routers. It inspects any VoIP traffic to prevent problems caused by firewalls and if necessary modifies the VoIP packets.

<div><strong>202</strong>: Success, accepted and forwarded the data.</div> <div><strong>400</strong>: Required fields missing.</div> <div><strong>401</strong>: Authorization required or customer id not found.</div> <div><strong>404</strong>: Location id, does not exist.</div> <div><strong>422</strong>: Invalid data.</div> <div><strong>500</strong>: Internal server error</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Customer id | 
  **locationId** | **string**|  | 
 **optional** | ***CustomerApiCustomerPrototypeOverlordUpdateSipAlgConfigOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CustomerApiCustomerPrototypeOverlordUpdateSipAlgConfigOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **mode** | **optional.String**|  string enum: [ AUTO, ENABLE, DISABLE ] | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomerPrototypeOverlordUpdateStatsConfig**
> CustomerPrototypeOverlordUpdateStatsConfig(ctx, id, locationId, optional)
Updates a stats config. Location Stats configuration, used to toggle which stats should be collected.

<div><strong>202</strong>: Success, accepted and forwarded the data.</div> <div><strong>400</strong>: Required fields missing.</div> <div><strong>401</strong>: Authorization required or customer id not found.</div> <div><strong>404</strong>: Location id, does not exist.</div> <div><strong>422</strong>: Invalid data.</div> <div><strong>500</strong>: Internal server error</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Customer id | 
  **locationId** | **string**|  | 
 **optional** | ***CustomerApiCustomerPrototypeOverlordUpdateStatsConfigOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CustomerApiCustomerPrototypeOverlordUpdateStatsConfigOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **offChannelScan24** | **optional.String**|  string enum: [ AUTO, ENABLE, DISABLE ] | 
 **offChannelScan50** | **optional.String**|  string enum: [ AUTO, ENABLE, DISABLE ] | 
 **offChannelScan60** | **optional.String**|  string enum: [ AUTO, ENABLE, DISABLE ] | 
 **clientAuthFails** | **optional.String**|  string enum: [ AUTO, ENABLE, DISABLE ] | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomerPrototypePatchAccessZone**
> WifiNetwork CustomerPrototypePatchAccessZone(ctx, id, locationId, zoneId, optional)
Update an access zone

<div><strong>200</strong>: Success, wifiNetwork returned</div> <div><strong>400</strong>: Required fields missing</div> <div><strong>401</strong>: Authorization required or customer id not found</div> <div><strong>404</strong>: Customer id, location id, or WifiNetwork does not exist</div> <div><strong>422</strong>: Validation failed</div> <div><strong>500</strong>: Internal server error</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Customer id | 
  **locationId** | **string**|  | 
  **zoneId** | **float64**| id of access zone | 
 **optional** | ***CustomerApiCustomerPrototypePatchAccessZoneOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CustomerApiCustomerPrototypePatchAccessZoneOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **description** | **optional.String**|  | 
 **accessibleDevices** | **optional.String**| array of home macs[] visible to this access zone | 

### Return type

[**WifiNetwork**](WifiNetwork.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomerPrototypePatchAppPrioritizationLocationConfig**
> XAny CustomerPrototypePatchAppPrioritizationLocationConfig(ctx, id, locationId, optional)
Update app prioritization config.

<div><strong>200</strong>: Success.</div> <div><strong>401</strong>: Authorization required or customer id not found</div> <div><strong>404</strong>: Location id or WifiNetwork does not exist and is not known to Plume</div> <div><strong>500</strong>: Internal server error</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Customer id | 
  **locationId** | **string**|  | 
 **optional** | ***CustomerApiCustomerPrototypePatchAppPrioritizationLocationConfigOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CustomerApiCustomerPrototypePatchAppPrioritizationLocationConfigOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **enabled** | **optional.Bool**| true if app prioritization is enabled | 
 **mode** | **optional.String**| App Prioritization mode - any of auto | enable | disable | 
 **isFirstTimeUserExperience** | **optional.Bool**| true if it is first time user experience | 
 **template** | **optional.String**| Template for app prioritization | 
 **customSettingEnabled** | **optional.Bool**| true if custom setting is enabled | 
 **customSetting** | **optional.String**| Settings for app prioritization | 

### Return type

[**XAny**](x-any.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomerPrototypePatchAppTimeIpFlow**
> interface{} CustomerPrototypePatchAppTimeIpFlow(ctx, id, locationId, enable, optional)
Patch IP flows config

<div><strong>200</strong>: Success.</div> <div><strong>400</strong>: Required fields missing or field type is incorrect.</div> <div><strong>401</strong>: Authorization required or customer id not found</div> <div><strong>404</strong>: Location id or device does not exist and is not known to Plume</div> <div><strong>500</strong>: Internal server error</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Customer id | 
  **locationId** | **string**|  | 
  **enable** | **bool**|  | 
 **optional** | ***CustomerApiCustomerPrototypePatchAppTimeIpFlowOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CustomerApiCustomerPrototypePatchAppTimeIpFlowOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **expiresAt** | **optional.String**|  | 

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomerPrototypePatchAttributesPatchCustomersid**
> Customer CustomerPrototypePatchAttributesPatchCustomersid(ctx, id, optional)
Patch attributes for a model instance and persist it into the data source.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Customer id | 
 **optional** | ***CustomerApiCustomerPrototypePatchAttributesPatchCustomersidOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CustomerApiCustomerPrototypePatchAttributesPatchCustomersidOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **data** | [**optional.Interface of Customer**](Customer.md)| An object of model property name/value pairs | 

### Return type

[**Customer**](Customer.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomerPrototypePatchAttributesPutCustomersid**
> Customer CustomerPrototypePatchAttributesPutCustomersid(ctx, id, optional)
Patch attributes for a model instance and persist it into the data source.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Customer id | 
 **optional** | ***CustomerApiCustomerPrototypePatchAttributesPutCustomersidOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CustomerApiCustomerPrototypePatchAttributesPutCustomersidOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **data** | [**optional.Interface of Customer**](Customer.md)| An object of model property name/value pairs | 

### Return type

[**Customer**](Customer.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomerPrototypePatchCampaignCaptivePortalBranding**
> NetworkConfig CustomerPrototypePatchCampaignCaptivePortalBranding(ctx, id, locationId, networkId, optional)
Patch the Captive Portal campaign branding properties for a given Location ID/NetworkId.

<div><strong>200</strong>: Success, campaign branding patched.</div> <div><strong>401</strong>: Authorization required or customer id not found</div> <div><strong>404</strong>: location id or CaptivePortal Network does not exist.</div> <div><strong>500</strong>: internal server error</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Customer id | 
  **locationId** | **string**|  | 
  **networkId** | **string**|  | 
 **optional** | ***CustomerApiCustomerPrototypePatchCampaignCaptivePortalBrandingOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CustomerApiCustomerPrototypePatchCampaignCaptivePortalBrandingOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **payload** | **optional.String**|  | 

### Return type

[**NetworkConfig**](NetworkConfig.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomerPrototypePatchCampaignCaptivePortalNetwork**
> NetworkConfig CustomerPrototypePatchCampaignCaptivePortalNetwork(ctx, id, locationId, networkId, optional)
Patch the Captive Portal campaign for a given Location ID/NetworkId.

<div><strong>200</strong>: Success, campaign patched.</div> <div><strong>401</strong>: Authorization required or customer id not found</div> <div><strong>404</strong>: location id or CaptivePortal Network does not exist.</div> <div><strong>500</strong>: internal server error</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Customer id | 
  **locationId** | **string**|  | 
  **networkId** | **string**|  | 
 **optional** | ***CustomerApiCustomerPrototypePatchCampaignCaptivePortalNetworkOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CustomerApiCustomerPrototypePatchCampaignCaptivePortalNetworkOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **campaignPayload** | **optional.String**|  | 

### Return type

[**NetworkConfig**](NetworkConfig.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomerPrototypePatchCaptivePortal**
> NetworkConfig CustomerPrototypePatchCaptivePortal(ctx, id, locationId, networkId, optional)
Update a Captive Portal for a given Location ID/NetworkId.

<div><strong>200</strong>: Success.</div> <div><strong>401</strong>: Authorization required or customer id not found</div> <div><strong>404</strong>: Location id or NetworkId does not exist and is not known to Plume</div> <div><strong>422</strong>: NetworkId/SSIDs must be the unique and valid values.</div> <div><strong>500</strong>: Internal server error</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Customer id | 
  **locationId** | **string**|  | 
  **networkId** | **string**|  | 
 **optional** | ***CustomerApiCustomerPrototypePatchCaptivePortalOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CustomerApiCustomerPrototypePatchCaptivePortalOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **ssid** | **optional.String**|  | 
 **enable** | **optional.Bool**|  | 
 **encryptionKey** | **optional.String**|  | 
 **bandwidthLimit** | **optional.String**| attributes: \&quot;enabled\&quot; boolean, \&quot;type\&quot;: \&quot;absolute\&quot;|\&quot;percentage\&quot;, \&quot;upload\&quot;/\&quot;download\&quot; - either as percentage or absolute (Mbps) | 
 **wpaMode** | **optional.String**|  | 
 **sessionTimeLimitSec** | **optional.Float64**|  | 

### Return type

[**NetworkConfig**](NetworkConfig.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomerPrototypePatchCaptivePortalAuthorizedClients**
> interface{} CustomerPrototypePatchCaptivePortalAuthorizedClients(ctx, id, locationId, networkId, mac, expireAt)
Post CaptivePortal authorized clients

<div><strong>204</strong>: Success.</div> <div><strong>401</strong>: Authorization required or customer id not found</div> <div><strong>404</strong>: Location id/NetworkId does not exist and is not known to Plume</div> <div><strong>500</strong>: Internal server error</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Customer id | 
  **locationId** | **string**|  | 
  **networkId** | **string**|  | 
  **mac** | **string**|  | 
  **expireAt** | **string**|  | 

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomerPrototypePatchCommand**
> CustomerPrototypePatchCommand(ctx, id, locationId, providerId, providerUserId)
Returns cloud migration status for customer

<div><strong>204</strong>: Success, no content.</div> <div><strong>400</strong>: Missing providerId or providerUserId body parameter</div> <div><strong>401</strong>: Authorization required or customer id not found</div> <div><strong>409</strong>: Accounts are already linked for providerUserId</div> <div><strong>422</strong>: Unfinished document for providerId does not exist</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Customer id | 
  **locationId** | **string**|  | 
  **providerId** | **string**| enum to identify provider ex. commandAlexa | 
  **providerUserId** | **string**| id of the user in provider&#39;s system | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomerPrototypePatchCustomDeviceType**
> CustomDeviceType CustomerPrototypePatchCustomDeviceType(ctx, id, locationId, mac, optional)
Update a Customer's device type configuration (user feedback).

<div><strong>200</strong>: Success, device type has been updated<br/>but not validated as a device that <br/>has ever connected.</div> <div><strong>400</strong>: nickname value must be defined.</div> <div><strong>404</strong>: customer id and/or mac does not exist.</div> <div><strong>422</strong>: nickname value must be less than 33 characters.</div> <div><strong>500</strong>: internal server error</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Customer id | 
  **locationId** | **string**|  | 
  **mac** | **string**|  | 
 **optional** | ***CustomerApiCustomerPrototypePatchCustomDeviceTypeOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CustomerApiCustomerPrototypePatchCustomDeviceTypeOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **category** | **optional.String**|  | 
 **brand** | **optional.String**|  | 
 **model** | **optional.String**|  | 
 **osName** | **optional.String**|  | 

### Return type

[**CustomDeviceType**](CustomDeviceType.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomerPrototypePatchCustomSharedSchedule**
> LocationCustomSchedule CustomerPrototypePatchCustomSharedSchedule(ctx, id, locationId, templateId, schedules, optional)
Patch a custom shared schedule freeze template for a Location ID.

<div><strong>200</strong>: Success, your new info looks good.</div> <div><strong>401</strong>: Authorization required or customer id not found</div> <div><strong>404</strong>: location id does not exist.</div> <div><strong>422</strong>: templateId must be defined.</div> <div><strong>422</strong>: schedules value is invalid.</div> <div><strong>425</strong>: templateId must belong to the location.</div> <div><strong>500</strong>: internal server error</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Customer id | 
  **locationId** | **string**|  | 
  **templateId** | **string**|  | 
  **schedules** | **string**|  | 
 **optional** | ***CustomerApiCustomerPrototypePatchCustomSharedScheduleOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CustomerApiCustomerPrototypePatchCustomSharedScheduleOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **name** | **optional.String**|  | 
 **type_** | **optional.String**|  | 

### Return type

[**LocationCustomSchedule**](LocationCustomSchedule.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomerPrototypePatchDevice**
> DeviceResponse CustomerPrototypePatchDevice(ctx, id, locationId, mac, optional)
Patches a single Device to mark it favorite for a Location ID.

<div><strong>200</strong>: Success, everything looks good.</div> <div><strong>404</strong>: customer id or location id does not exist. Or, device not found in this network 's history.</div> <div><strong>500</strong>: internal server error</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Customer id | 
  **locationId** | **string**|  | 
  **mac** | **string**| mac of device | 
 **optional** | ***CustomerApiCustomerPrototypePatchDeviceOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CustomerApiCustomerPrototypePatchDeviceOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **favorite** | **optional.Bool**|  | 
 **nickname** | **optional.String**|  | 
 **mobileAppDeviceUuid** | **optional.String**|  | 

### Return type

[**DeviceResponse**](DeviceResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomerPrototypePatchDeviceAppTime**
> interface{} CustomerPrototypePatchDeviceAppTime(ctx, id, locationId, mac, optional)
Update a Device's AppTime config by location ID.

<div><strong>200</strong>: Success.</div> <div><strong>400</strong>: Required fields missing or field type is incorrect.</div> <div><strong>401</strong>: Authorization required or customer id not found</div> <div><strong>404</strong>: Location id or device does not exist and is not known to Plume</div> <div><strong>500</strong>: Internal server error</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Customer id | 
  **locationId** | **string**|  | 
  **mac** | **string**|  | 
 **optional** | ***CustomerApiCustomerPrototypePatchDeviceAppTimeOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CustomerApiCustomerPrototypePatchDeviceAppTimeOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **enable** | **optional.Bool**|  | 

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomerPrototypePatchDeviceClientSteering**
> ClientSteeringConfiguration CustomerPrototypePatchDeviceClientSteering(ctx, id, locationId, mac, optional)
Toggle auto:on/off client steering for a device.

<div><strong>200</strong>: Success, updated.</div> <div><strong>400</strong>: Required fields missing.</div> <div><strong>401</strong>: Authorization required or customer id not found.</div> <div><strong>404</strong>: Location id, does not exist.</div> <div><strong>422</strong>: Invalid mac address.</div> <div><strong>500</strong>: Internal server error</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Customer id | 
  **locationId** | **string**|  | 
  **mac** | **string**| device mac address | 
 **optional** | ***CustomerApiCustomerPrototypePatchDeviceClientSteeringOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CustomerApiCustomerPrototypePatchDeviceClientSteeringOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **auto** | **optional.Bool**|  | 
 **steeringClass** | **optional.String**| override deviceTypeId for testing purposes | 

### Return type

[**ClientSteeringConfiguration**](ClientSteeringConfiguration.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomerPrototypePatchDeviceGroup**
> NetworkAccessDeviceGroup CustomerPrototypePatchDeviceGroup(ctx, id, locationId, networkId, groupId, optional)
Change a device group name or device members.

<div><strong>200</strong>: Success.</div> <div><strong>422</strong>: Schema validation failed.</div> <div><strong>404</strong>: Location does not exist.</div> <div><strong>404</strong>: Network does not exist.</div> <div><strong>403</strong>: Not allowed to modify standalone groups or groups in unsupported networks.</div> <div><strong>401</strong>: Unauthorized.</div> <div><strong>400</strong>: Invalid JSON or missing arguments.</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Customer id | 
  **locationId** | **string**|  | 
  **networkId** | **string**|  | 
  **groupId** | **string**|  | 
 **optional** | ***CustomerApiCustomerPrototypePatchDeviceGroupOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CustomerApiCustomerPrototypePatchDeviceGroupOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **name** | **optional.String**|  | 
 **devices** | **optional.String**|  | 

### Return type

[**NetworkAccessDeviceGroup**](NetworkAccessDeviceGroup.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomerPrototypePatchDeviceOHPConfiguration**
> interface{} CustomerPrototypePatchDeviceOHPConfiguration(ctx, id, locationId, mac, optional)
Update the Device UUID Mapping for Out of Home Protection.

<div><strong>204</strong>: Success.</div> <div><strong>400</strong>: Required fields missing or field type is incorrect.</div> <div><strong>401</strong>: Authorization required or customer id not found</div> <div><strong>404</strong>: Location id or Device does not exist and is not known to Plume</div> <div><strong>500</strong>: Internal server error</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Customer id | 
  **locationId** | **string**|  | 
  **mac** | **string**|  | 
 **optional** | ***CustomerApiCustomerPrototypePatchDeviceOHPConfigurationOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CustomerApiCustomerPrototypePatchDeviceOHPConfigurationOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **oHPNotificationsFlags** | **optional.String**| OHP feature flags | 
 **disableMobilizeSdk** | **optional.Bool**| enable or disable OHP SDK on the device | 

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomerPrototypePatchDeviceQos**
> []XAny CustomerPrototypePatchDeviceQos(ctx, id, locationId, mac, prioritization)
Update QoS of a single device

<div><strong>202</strong>: Success.</div> <div><strong>404</strong>: Location does not exist.</div> <div><strong>422</strong>: Prioritization is not a valid value.</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Customer id | 
  **locationId** | **string**|  | 
  **mac** | **string**|  | 
  **prioritization** | **string**|  | 

### Return type

[**[]XAny**](x-any.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomerPrototypePatchDeviceSecurityPolicy**
> Person CustomerPrototypePatchDeviceSecurityPolicy(ctx, id, locationId, mac, optional)
Update a Device's Security Policy for a location ID.

<div><strong>200</strong>: Success.</div> <div><strong>401</strong>: Required fields missing or field type is incorrect.</div> <div><strong>401</strong>: Authorization required or customer id not found</div> <div><strong>404</strong>: Location id, WifiNetwork, or Person id does not exist and is not known to Plume</div> <div><strong>409</strong>: Device is assigned to a person so its security policy must be configured on the Person</div> <div><strong>422</strong>: Mac addresses must be valid.</div> <div><strong>500</strong>: Internal server error</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Customer id | 
  **locationId** | **string**|  | 
  **mac** | **string**|  | 
 **optional** | ***CustomerApiCustomerPrototypePatchDeviceSecurityPolicyOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CustomerApiCustomerPrototypePatchDeviceSecurityPolicyOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **secureAndProtect** | **optional.Bool**|  | 
 **iotProtect** | **optional.Bool**|  | 
 **iotProtectReason** | **optional.String**|  | 
 **content** | **optional.String**| Valid values: &#39;kids || teenagers || adBlocking || adultAndSensitive || workAppropriate&#39; | 

### Return type

[**Person**](Person.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomerPrototypePatchDeviceSoundingState**
> interface{} CustomerPrototypePatchDeviceSoundingState(ctx, id, locationId, soundingStates)
Patch the sounding states for the given devices

<div><strong>200</strong>: Success, device sounding states returned.</div> <div><strong>401</strong>: Authorization required or customer id not found</div> <div><strong>404</strong>: customer id or location id does not exist and is not known to Plume</div> <div><strong>500</strong>: internal server error</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Customer id | 
  **locationId** | **string**|  | 
  **soundingStates** | **interface{}**|  | 

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomerPrototypePatchDpp**
> CustomerPrototypePatchDpp(ctx, id, locationId, mode)
Patch the DPP configuration mode for a Location ID.

<div><strong>202</strong>: Success, DPP updated.</div> <div><strong>400</strong>: Required fields missing or field type is incorrect.</div> <div><strong>401</strong>: Authorization required or customer id not found</div> <div><strong>422</strong>: DPP value is invalid.</div> <div><strong>500</strong>: internal server error</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Customer id | 
  **locationId** | **string**|  | 
  **mode** | **string**| auto || enable || disable | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomerPrototypePatchEthernetLan**
> EthernetLan CustomerPrototypePatchEthernetLan(ctx, id, locationId, ethernetLan)
Update the ethernetLan setting for a Location ID.

Supported modes are: * enable/disable/auto  <div><strong>200</strong>: Success, new ethernetLan settings saved.</div> <div><strong>401</strong>: Authorization required or customer id not found</div> <div><strong>422</strong>: Input validation error, see output for details.</div> <div><strong>500</strong>: Internal server error</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Customer id | 
  **locationId** | **string**|  | 
  **ethernetLan** | [**EthernetLan**](EthernetLan.md)| ethernetLan object | 

### Return type

[**EthernetLan**](EthernetLan.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomerPrototypePatchFlowStats**
> CustomerPrototypePatchFlowStats(ctx, id, locationId, optional)
Patches the flow stats configuration

<div><strong>202</strong>: Success, your new info looks good.</div> <div><strong>401</strong>: Authorization required or customer id not found.</div> <div><strong>404</strong>: location id, does not exist.</div> <div><strong>422</strong>: Input value is invalid.</div> <div><strong>500</strong>: internal server error</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Customer id | 
  **locationId** | **string**|  | 
 **optional** | ***CustomerApiCustomerPrototypePatchFlowStatsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CustomerApiCustomerPrototypePatchFlowStatsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **iotDeviceConfig** | **optional.String**| auto || enable || disable | 
 **screenDeviceConfig** | **optional.String**| auto || enable || disable | 
 **lanIotDeviceConfig** | **optional.String**| auto || enable || disable | 
 **interfaceStatsConfig** | **optional.String**| auto || enable || disable | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomerPrototypePatchFrontHaul**
> NetworkConfig CustomerPrototypePatchFrontHaul(ctx, id, locationId, networkId, optional)
Update a Front Haul for a given Location ID/NetworkId.

<div><strong>200</strong>: Success.</div> <div><strong>401</strong>: Authorization required or customer id not found</div> <div><strong>404</strong>: Location id or NetworkId does not exist and is not known to Plume</div> <div><strong>422</strong>: NetworkId/SSIDs must be the unique and valid values.</div> <div><strong>500</strong>: Internal server error</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Customer id | 
  **locationId** | **string**|  | 
  **networkId** | **string**|  | 
 **optional** | ***CustomerApiCustomerPrototypePatchFrontHaulOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CustomerApiCustomerPrototypePatchFrontHaulOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **ssid** | **optional.String**|  | 
 **enable** | **optional.Bool**|  | 
 **encryptionKey** | **optional.String**|  | 
 **accessZone** | **optional.String**|  | 
 **wpaMode** | **optional.String**|  | 
 **ssidBroadcast** | **optional.Bool**|  | 

### Return type

[**NetworkConfig**](NetworkConfig.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomerPrototypePatchGroupOfUnassignedDevicesSecurityPolicy**
> interface{} CustomerPrototypePatchGroupOfUnassignedDevicesSecurityPolicy(ctx, id, locationId, optional)
Update a Location's Default Device Group Security Policy by location ID.

<div><strong>200</strong>: Success.</div> <div><strong>400</strong>: Required fields missing or field type is incorrect.</div> <div><strong>401</strong>: Authorization required or customer id not found</div> <div><strong>404</strong>: Location id or WifiNetwork does not exist and is not known to Plume</div> <div><strong>500</strong>: Internal server error</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Customer id | 
  **locationId** | **string**|  | 
 **optional** | ***CustomerApiCustomerPrototypePatchGroupOfUnassignedDevicesSecurityPolicyOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CustomerApiCustomerPrototypePatchGroupOfUnassignedDevicesSecurityPolicyOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **secureAndProtect** | **optional.Bool**|  | 
 **iotProtect** | **optional.Bool**|  | 
 **content** | **optional.String**| Valid values: &#39;kids || teenagers || adBlocking || adultAndSensitive || workAppropriate&#39; | 

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomerPrototypePatchHomeAwayActive**
> HomeSecurity CustomerPrototypePatchHomeAwayActive(ctx, id, locationId, homeAwayActive)
Enable/disable homeAway wifiMotionEvents activation for this location

<div><strong>200</strong>: Success, updated HomeSecurity object returned.</div> <div><strong>401</strong>: Authorization required or customer id not found</div> <div><strong>404</strong>: customer id or location id does not exist and is not known to Plume</div> <div><strong>500</strong>: internal server error</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Customer id | 
  **locationId** | **string**|  | 
  **homeAwayActive** | **bool**| Enable/disable motion events based on location Homeaway state | 

### Return type

[**HomeSecurity**](HomeSecurity.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomerPrototypePatchHomeSecurity**
> HomeSecurity CustomerPrototypePatchHomeSecurity(ctx, id, locationId, optional)
Enable/disable live motion streaming and/or motion events for this location

<div><strong>200</strong>: Success, updated HomeSecurity object returned.</div> <div><strong>401</strong>: Authorization required or customer id not found</div> <div><strong>404</strong>: customer id or location id does not exist and is not known to Plume</div> <div><strong>500</strong>: internal server error</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Customer id | 
  **locationId** | **string**|  | 
 **optional** | ***CustomerApiCustomerPrototypePatchHomeSecurityOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CustomerApiCustomerPrototypePatchHomeSecurityOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **source** | **optional.String**| Source of patch request; must be one of \&quot;user\&quot; or \&quot;geofence\&quot; | 
 **liveMotionEnabled** | **optional.Bool**|  | 
 **motionEventsEnabled** | **optional.Bool**|  | 
 **homeAwayActive** | **optional.Bool**| Enable/disable motion events based on location Homeaway state | 

### Return type

[**HomeSecurity**](HomeSecurity.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomerPrototypePatchHomeSecuritySensitivity**
> HomeSecurity CustomerPrototypePatchHomeSecuritySensitivity(ctx, id, locationId, optional)
Configure motion event configuration for this location

<div><strong>200</strong>: Success, updated HomeSecurity object returned.</div> <div><strong>401</strong>: Authorization required or customer id not found</div> <div><strong>404</strong>: customer id or location id does not exist and is not known to Plume</div> <div><strong>500</strong>: internal server error</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Customer id | 
  **locationId** | **string**|  | 
 **optional** | ***CustomerApiCustomerPrototypePatchHomeSecuritySensitivityOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CustomerApiCustomerPrototypePatchHomeSecuritySensitivityOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **cooldown** | **optional.Float64**| sets required rest period for motion detected events to end, in seconds | 
 **petMode** | **optional.String**| adjusts sensitivity of motion detected events for pets; must be one of \&quot;none\&quot;, \&quot;under10\&quot;, \&quot;10to30\&quot;, \&quot;over30\&quot; and can only be set if sensitivity &#x3D; high | 
 **sensitivity** | **optional.String**| adjusts sensitivity of motion detected events; must be one of \&quot;low\&quot;, \&quot;medium\&quot;, \&quot;high\&quot; | 

### Return type

[**HomeSecurity**](HomeSecurity.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomerPrototypePatchIPv6**
> interface{} CustomerPrototypePatchIPv6(ctx, id, locationId, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Customer id | 
  **locationId** | **string**|  | 
 **optional** | ***CustomerApiCustomerPrototypePatchIPv6Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CustomerApiCustomerPrototypePatchIPv6Opts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **mode** | **optional.String**|  | 
 **dns** | **optional.String**|  | 
 **addressingConfig** | **optional.String**|  | 

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomerPrototypePatchKvConfigs**
> []KvConfig CustomerPrototypePatchKvConfigs(ctx, id, locationId, nodeId, kvConfigs)
Retrieve all kvConfigs on a particular Node for a Location ID.

<div><strong>200</strong>: Success, your new info looks good.</div> <div><strong>401</strong>: Authorization required or customer id not found</div> <div><strong>404</strong>: location id does not exist.</div> <div><strong>422</strong>: nodeId must be defined.</div> <div><strong>425</strong>: nodeId must belong to the location.</div> <div><strong>500</strong>: internal server error</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Customer id | 
  **locationId** | **string**|  | 
  **nodeId** | **string**|  | 
  **kvConfigs** | **string**|  | 

### Return type

[**[]KvConfig**](KvConfig.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomerPrototypePatchLocation**
> Location CustomerPrototypePatchLocation(ctx, id, locationId, optional)
Update a Location's serviceId.

<div><strong>200</strong>: Success.</div> <div><strong>401</strong>: Authorization required or customer id not found</div> <div><strong>404</strong>: Location id does not exist and is not known to Plume</div> <div><strong>422</strong>: You must specify at least one parameter to patch.</div> <div><strong>422</strong>: Only integration role can set profile to property.</div> <div><strong>500</strong>: Internal server error</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Customer id | 
  **locationId** | **string**|  | 
 **optional** | ***CustomerApiCustomerPrototypePatchLocationOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CustomerApiCustomerPrototypePatchLocationOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **serviceId** | **optional.String**|  | 
 **profile** | **optional.String**|  | 

### Return type

[**Location**](Location.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomerPrototypePatchLocationAppTime**
> interface{} CustomerPrototypePatchLocationAppTime(ctx, id, locationId, optional)
Update a Location's AppTime config by location ID.

<div><strong>200</strong>: Success.</div> <div><strong>400</strong>: Required fields missing or field type is incorrect.</div> <div><strong>401</strong>: Authorization required or customer id not found</div> <div><strong>404</strong>: Location id not exist and is not known to Plume</div> <div><strong>500</strong>: Internal server error</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Customer id | 
  **locationId** | **string**|  | 
 **optional** | ***CustomerApiCustomerPrototypePatchLocationAppTimeOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CustomerApiCustomerPrototypePatchLocationAppTimeOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **enable** | **optional.Bool**|  | 
 **appliesToAllDevices** | **optional.Bool**|  | [default to false]
 **sandboxSizeMb** | **optional.String**|  | 

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomerPrototypePatchLocationBandSteering**
> LocationBandSteering CustomerPrototypePatchLocationBandSteering(ctx, id, locationId, optional)
Set mode for band steering

<div><strong>200</strong>: Success, updated.</div> <div><strong>400</strong>: Required fields missing.</div> <div><strong>401</strong>: Authorization required or customer id not found.</div> <div><strong>404</strong>: Location id, does not exist.</div> <div><strong>422</strong>: Invalid mode.</div> <div><strong>500</strong>: Internal server error</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Customer id | 
  **locationId** | **string**|  | 
 **optional** | ***CustomerApiCustomerPrototypePatchLocationBandSteeringOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CustomerApiCustomerPrototypePatchLocationBandSteeringOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **mode** | **optional.String**| auto || enable || disable | 

### Return type

[**LocationBandSteering**](LocationBandSteering.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomerPrototypePatchLocationFreezeAutoExpire**
> CustomerPrototypePatchLocationFreezeAutoExpire(ctx, id, locationId, optional)
Put all devices except some to be frozen for a Location ID.

<div><strong>200</strong>: Success, updated.</div> <div><strong>401</strong>: Authorization required or customer id not found.</div> <div><strong>404</strong>: Location id, does not exist.</div> <div><strong>500</strong>: Internal server error</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Customer id | 
  **locationId** | **string**|  | 
 **optional** | ***CustomerApiCustomerPrototypePatchLocationFreezeAutoExpireOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CustomerApiCustomerPrototypePatchLocationFreezeAutoExpireOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **includedDeviceMacs** | **optional.String**|  | 
 **includedPersonIds** | **optional.String**|  | 
 **expiresAt** | **optional.String**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomerPrototypePatchLocationManager**
> Location CustomerPrototypePatchLocationManager(ctx, id, locationId, managerId, optional)
Update type of access of manager on location.

<div><strong>200</strong>: Success.</div> <div><strong>401</strong>: Authorization required or customer id not found</div> <div><strong>404</strong>: Location id does not exist and is not known to Plume</div> <div><strong>500</strong>: Internal server error</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Customer id | 
  **locationId** | **string**|  | 
  **managerId** | **string**|  | 
 **optional** | ***CustomerApiCustomerPrototypePatchLocationManagerOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CustomerApiCustomerPrototypePatchLocationManagerOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **accessType** | **optional.String**|  | 
 **name** | **optional.String**|  | 
 **notificationOptions** | **optional.String**|  | 

### Return type

[**Location**](Location.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomerPrototypePatchLocationQoeLiveMode**
> interface{} CustomerPrototypePatchLocationQoeLiveMode(ctx, id, locationId, enable, optional)
Update the location qoe liveMode by api call and Kafka message

<div><strong>200</strong>: Success, the new info looks good.</div> <div><strong>401</strong>: Authorization required or customer id not found</div> <div><strong>400</strong>: enalbe and expiresAt, reportingInterval validation error.</div> <div><strong>422</strong>: expiresAt and reportingInterval validation error.</div> <div><strong>500</strong>: internal server error</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Customer id | 
  **locationId** | **string**|  | 
  **enable** | **bool**|  | 
 **optional** | ***CustomerApiCustomerPrototypePatchLocationQoeLiveModeOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CustomerApiCustomerPrototypePatchLocationQoeLiveModeOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **expiresAt** | **optional.String**|  | 
 **reportingInterval** | **optional.Float64**|  | 

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomerPrototypePatchLocationSecurityPolicy**
> interface{} CustomerPrototypePatchLocationSecurityPolicy(ctx, id, locationId, optional)
Update a Location's Security Policy by location ID.

<div><strong>200</strong>: Success.</div> <div><strong>400</strong>: Required fields missing or field type is incorrect.</div> <div><strong>401</strong>: Authorization required or customer id not found</div> <div><strong>404</strong>: Location id or WifiNetwork does not exist and is not known to Plume</div> <div><strong>500</strong>: Internal server error</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Customer id | 
  **locationId** | **string**|  | 
 **optional** | ***CustomerApiCustomerPrototypePatchLocationSecurityPolicyOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CustomerApiCustomerPrototypePatchLocationSecurityPolicyOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **secureAndProtect** | **optional.Bool**|  | 
 **iotProtect** | **optional.Bool**|  | 
 **content** | **optional.String**| Valid values: &#39;kids || teenagers || adBlocking || adultAndSensitive || workAppropriate&#39; | 
 **appliesToAllDevices** | **optional.String**| hash map of security policy IDs that should be applied to all devices | 

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomerPrototypePatchMulticast**
> Multicast CustomerPrototypePatchMulticast(ctx, id, locationId, multicast)
Update the multicast settings for a Location ID.

Supported modes for individual settings are: * igmpSnooping: enable/disable/auto * igmpProxy: igmpv1/igmpv2/igmpv3/disable/auto * mldProxy: mldv1/mldv2/disable/disable/auto * multicastToUnicast: enable/disable/auto  <div><strong>200</strong>: Success, new multicast settings saved.</div> <div><strong>401</strong>: Authorization required or customer id not found</div> <div><strong>422</strong>: Input validation error, see output for details.</div> <div><strong>500</strong>: Internal server error</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Customer id | 
  **locationId** | **string**|  | 
  **multicast** | [**Multicast**](Multicast.md)| multicast object | 

### Return type

[**Multicast**](Multicast.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomerPrototypePatchNetworkAccessNetwork**
> interface{} CustomerPrototypePatchNetworkAccessNetwork(ctx, id, locationId, networkId, purgatory)
Enable or disable purgatory in the network

<div><strong>204</strong>: Success.</div> <div><strong>404</strong>: Location does not exist.</div> <div><strong>404</strong>: Network does not exist.</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Customer id | 
  **locationId** | **string**|  | 
  **networkId** | **string**|  | 
  **purgatory** | **bool**|  | 

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomerPrototypePatchOptimizations**
> Optimizations CustomerPrototypePatchOptimizations(ctx, id, locationId, optional)
Enable/disable optimizations for a Location ID.

<div><strong>200</strong>: Success, updated.</div> <div><strong>401</strong>: Authorization required or customer id not found.</div> <div><strong>404</strong>: Location id, does not exist.</div> <div><strong>422</strong>: Invalid dfsMode, prefer160MhzMode, zeroWaitDfsMode, hopPenalty or preCACScheduler provided.</div> <div><strong>500</strong>: Internal server error</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Customer id | 
  **locationId** | **string**|  | 
 **optional** | ***CustomerApiCustomerPrototypePatchOptimizationsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CustomerApiCustomerPrototypePatchOptimizationsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **auto** | **optional.String**| defaults to true | 
 **dfsMode** | **optional.String**| enum of values include: auto, enable, disable, demo, HomeNonDFSChannels, usDfs, deviceAware | 
 **prefer160MhzMode** | **optional.String**| enum of values include: auto, enable, disable | 
 **zeroWaitDfsMode** | **optional.String**| enum of values include: auto, enable, disable | 
 **hopPenalty** | **optional.String**| enum of values include: auto, low, medium, high | 
 **preCACScheduler** | **optional.String**| enum of values include: auto, enable, disable | [default to auto]

### Return type

[**Optimizations**](Optimizations.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomerPrototypePatchPerson**
> Person CustomerPrototypePatchPerson(ctx, id, locationId, personId, optional)
Update a Person for a location ID.

<div><strong>200</strong>: Success.</div> <div><strong>401</strong>: Authorization required or customer id not found</div> <div><strong>404</strong>: Location id or Person id does not exist and is not known to Plume</div> <div><strong>409</strong>: primaryDevice is not included in the list of assignedDevices[]</div> <div><strong>422</strong>: Mac addresses must be valid.</div> <div><strong>500</strong>: Internal server error</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Customer id | 
  **locationId** | **string**|  | 
  **personId** | **string**|  | 
 **optional** | ***CustomerApiCustomerPrototypePatchPersonOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CustomerApiCustomerPrototypePatchPersonOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **nickname** | **optional.String**|  | 
 **imageId** | **optional.String**| unique identifier for referencing a Person&#39;s hosted profile image | 
 **primaryDevice** | **optional.String**| mac addresses of Person&#39;s primary device | 
 **assignedDevices** | **optional.String**| mac addresses assigned to this Person | 
 **homeAwayNotification** | **optional.Bool**| track person homeAway state | 
 **permission** | **optional.String**| permission object for creating or deleting the manager | 
 **email** | **optional.String**| email for sending the manager invite | 
 **serviceLinking** | **optional.String**| serviceLinking object that links this Person object to a 3rd party&#39;s Person | 

### Return type

[**Person**](Person.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomerPrototypePatchPersonAppTime**
> interface{} CustomerPrototypePatchPersonAppTime(ctx, id, locationId, personId, optional)
Update a Person's AppTime config by location ID.

<div><strong>200</strong>: Success.</div> <div><strong>400</strong>: Required fields missing or field type is incorrect.</div> <div><strong>401</strong>: Authorization required or customer id not found</div> <div><strong>404</strong>: Location id or person does not exist and is not known to Plume</div> <div><strong>500</strong>: Internal server error</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Customer id | 
  **locationId** | **string**|  | 
  **personId** | **string**|  | 
 **optional** | ***CustomerApiCustomerPrototypePatchPersonAppTimeOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CustomerApiCustomerPrototypePatchPersonAppTimeOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **enable** | **optional.Bool**|  | 

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomerPrototypePatchPersonProfile**
> interface{} CustomerPrototypePatchPersonProfile(ctx, id, locationId, personId, optional)
Update a Person's Profile for a location ID.

<div><strong>200</strong>: Success.</div> <div><strong>400</strong>: Required fields missing or field type is incorrect.</div> <div><strong>401</strong>: Authorization required or customer id not found</div> <div><strong>404</strong>: Location id, WifiNetwork, or Person id does not exist and is not known to Plume</div> <div><strong>500</strong>: Internal server error</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Customer id | 
  **locationId** | **string**|  | 
  **personId** | **string**|  | 
 **optional** | ***CustomerApiCustomerPrototypePatchPersonProfileOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CustomerApiCustomerPrototypePatchPersonProfileOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **type_** | **optional.String**| Valid values: &#39;employee&#39; | 

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomerPrototypePatchPersonSecurityPolicy**
> interface{} CustomerPrototypePatchPersonSecurityPolicy(ctx, id, locationId, personId, optional)
Update a Person's Security Policy for a location ID.

<div><strong>200</strong>: Success.</div> <div><strong>400</strong>: Required fields missing or field type is incorrect.</div> <div><strong>401</strong>: Authorization required or customer id not found</div> <div><strong>404</strong>: Location id, WifiNetwork, or Person id does not exist and is not known to Plume</div> <div><strong>500</strong>: Internal server error</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Customer id | 
  **locationId** | **string**|  | 
  **personId** | **string**|  | 
 **optional** | ***CustomerApiCustomerPrototypePatchPersonSecurityPolicyOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CustomerApiCustomerPrototypePatchPersonSecurityPolicyOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **secureAndProtect** | **optional.Bool**|  | 
 **iotProtect** | **optional.Bool**|  | 
 **content** | **optional.String**| Valid values: &#39;kids || teenagers || adBlocking || adultAndSensitive || workAppropriate&#39; | 

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomerPrototypePatchRemoteConnectionsConfig**
> interface{} CustomerPrototypePatchRemoteConnectionsConfig(ctx, id, locationId, mode)
Patch a Remote Connections Config for the given Location ID.

<div><strong>200</strong>: Success.</div> <div><strong>401</strong>: Authorization required or customer id not found</div> <div><strong>404</strong>: Location id does not exist and is not known to Plume</div> <div><strong>500</strong>: Internal server error</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Customer id | 
  **locationId** | **string**|  | 
  **mode** | **string**| Any of \&quot;auto\&quot;, \&quot;enabled\&quot;, \&quot;disabled\&quot;, \&quot;highRiskOnly\&quot; | 

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomerPrototypePatchRoom**
> Room CustomerPrototypePatchRoom(ctx, id, locationId, roomId, optional)
Patch a Room for a Location ID/Room ID.

<div><strong>200</strong>: Success.</div> <div><strong>401</strong>: Authorization required or customer id not found</div> <div><strong>404</strong>: Location id does not exist and is not known to Plume</div> <div><strong>422</strong>: Devices and Nodes must be defined and mac addresses must be valid.</div> <div><strong>500</strong>: Internal server error</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Customer id | 
  **locationId** | **string**|  | 
  **roomId** | **string**|  | 
 **optional** | ***CustomerApiCustomerPrototypePatchRoomOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CustomerApiCustomerPrototypePatchRoomOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **name** | **optional.String**| name of this Room | 
 **devices** | **optional.String**| mac addresses of devices assigned to this Room | 
 **nodes** | **optional.String**| nodeIds assigned to this Room | 

### Return type

[**Room**](Room.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomerPrototypePatchSecurityConfiguration**
> SecurityConfiguration CustomerPrototypePatchSecurityConfiguration(ctx, id, locationId, securityConfig)
Patch Security Configurations for location (preferredIntelligence, etc)

<div><strong>200</strong>: Success, updated.</div> <div><strong>400</strong>: Required fields missing.</div> <div><strong>401</strong>: Authorization required or customer id not found.</div> <div><strong>404</strong>: Location id, does not exist.</div> <div><strong>422</strong>: Invalid securityConfig.</div> <div><strong>500</strong>: Internal server error</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Customer id | 
  **locationId** | **string**|  | 
  **securityConfig** | **string**|  | 

### Return type

[**SecurityConfiguration**](SecurityConfiguration.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomerPrototypePatchServiceLevel**
> interface{} CustomerPrototypePatchServiceLevel(ctx, id, locationId, status)
Set the service level for this location

<div><strong>200</strong>: Success, updated service Level object returned.</div> <div><strong>400</strong>: Required fields missing.</div> <div><strong>401</strong>: Authorization required or customer id not found</div> <div><strong>404</strong>: customer id or location id does not exist and is not known to Plume</div> <div><strong>422</strong>: Invalid 'status' value.</div> <div><strong>500</strong>: internal server error</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Customer id | 
  **locationId** | **string**|  | 
  **status** | **string**|  | 

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomerPrototypePatchSubscription**
> CustomerPrototypePatchSubscription(ctx, id, locationId, status)
Patch Subscription details for this location

<div><strong>202</strong>: Success, status patched</div> <div><strong>401</strong>: Authorization required or customer id not found</div> <div><strong>404</strong>: customer id or location id does not exist and is not known to Plume</div> <div><strong>422</strong>: Status is invalid.</div> <div><strong>500</strong>: internal server error</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Customer id | 
  **locationId** | **string**|  | 
  **status** | **string**| enum of values include: Active, Suspended | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomerPrototypePatchWifiMotion**
> WifiMotion CustomerPrototypePatchWifiMotion(ctx, id, locationId, auto)
Enable/disable WifiMotion feature for this location

<div><strong>200</strong>: Success, updated object returned.</div> <div><strong>401</strong>: Authorization required or customer id not found</div> <div><strong>404</strong>: customer id or location id does not exist and is not known to Plume</div> <div><strong>500</strong>: internal server error</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Customer id | 
  **locationId** | **string**|  | 
  **auto** | **bool**|  | 

### Return type

[**WifiMotion**](WifiMotion.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomerPrototypePatchWifiNetwork**
> WifiNetwork CustomerPrototypePatchWifiNetwork(ctx, id, locationId, optional)
Update the SSID of the WifiNetwork

<div><strong>200</strong>: Success, access zone returned</div> <div><strong>400</strong>: Required fields missing</div> <div><strong>401</strong>: Authorization required or customer id not found</div> <div><strong>404</strong>: Customer id, location id, or WifiNetwork does not exist</div> <div><strong>422</strong>: Validation failed</div> <div><strong>500</strong>: Internal server error</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Customer id | 
  **locationId** | **string**|  | 
 **optional** | ***CustomerApiCustomerPrototypePatchWifiNetworkOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CustomerApiCustomerPrototypePatchWifiNetworkOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **ssid** | **optional.String**|  | 
 **uapsd** | **optional.Bool**|  | 
 **groupRekey** | **optional.String**| auto || enable || disable | 
 **fastTransition** | **optional.String**| auto || enable || disable | 
 **minWifiMode24** | **optional.String**| auto || 11b || 11g || 11n | 
 **privateMode** | **optional.Bool**| Stop collecting user info like DNS-Queries, UserAgent etc | 
 **enabled** | **optional.Bool**| enabled:true for active WiFi radios, enabled:false to turn &#x60;off&#x60; all WiFi radios | 

### Return type

[**WifiNetwork**](WifiNetwork.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomerPrototypePostCampaignPreviewCaptivePortalNetwork**
> NetworkConfig CustomerPrototypePostCampaignPreviewCaptivePortalNetwork(ctx, id, locationId, networkId, optional)
POST Captive Portal campaign preview for a given Location ID/NetworkId.

<div><strong>200</strong>: Success, campaign posted.</div> <div><strong>401</strong>: Authorization required or customer id not found</div> <div><strong>404</strong>: location id or CaptivePortal Network does not exist.</div> <div><strong>500</strong>: internal server error</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Customer id | 
  **locationId** | **string**|  | 
  **networkId** | **string**|  | 
 **optional** | ***CustomerApiCustomerPrototypePostCampaignPreviewCaptivePortalNetworkOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CustomerApiCustomerPrototypePostCampaignPreviewCaptivePortalNetworkOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **campaignPayload** | **optional.String**|  | 

### Return type

[**NetworkConfig**](NetworkConfig.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomerPrototypePostCaptivePortal**
> NetworkConfig CustomerPrototypePostCaptivePortal(ctx, id, locationId, ssid, optional)
Create a Captive Portal Network for a Location ID.

<div><strong>200</strong>: Success.</div> <div><strong>401</strong>: Authorization required or customer id not found</div> <div><strong>404</strong>: Location id or NetworkId does not exist and is not known to Plume</div> <div><strong>422</strong>: NetworkId/SSIDs must be the unique and valid values.</div> <div><strong>500</strong>: Internal server error</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Customer id | 
  **locationId** | **string**|  | 
  **ssid** | **string**|  | 
 **optional** | ***CustomerApiCustomerPrototypePostCaptivePortalOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CustomerApiCustomerPrototypePostCaptivePortalOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **networkId** | **optional.String**|  | 
 **enable** | **optional.Bool**|  | [default to true]
 **encryptionKey** | **optional.String**|  | 
 **bandwidthLimit** | **optional.String**| attributes: \&quot;enabled\&quot; boolean, \&quot;type\&quot;: \&quot;absolute\&quot;|\&quot;percentage\&quot;, \&quot;upload\&quot;/\&quot;download\&quot; - either as percentage or absolute (Mbps) | 
 **sessionTimeLimitSec** | **optional.Float64**|  | 
 **wpaMode** | **optional.String**|  | 
 **language** | **optional.String**|  | 

### Return type

[**NetworkConfig**](NetworkConfig.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomerPrototypePostCaptivePortalAuthorizedClients**
> interface{} CustomerPrototypePostCaptivePortalAuthorizedClients(ctx, id, locationId, networkId, mac, expireAt)
Post CaptivePortal authorized clients

<div><strong>204</strong>: Success.</div> <div><strong>401</strong>: Authorization required or customer id not found</div> <div><strong>404</strong>: Location id/NetworkId does not exist and is not known to Plume</div> <div><strong>500</strong>: Internal server error</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Customer id | 
  **locationId** | **string**|  | 
  **networkId** | **string**|  | 
  **mac** | **string**|  | 
  **expireAt** | **string**|  | 

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomerPrototypePostCaptivePortalCampaignAsset**
> NetworkConfig CustomerPrototypePostCaptivePortalCampaignAsset(ctx, id, locationId)
Upload campaign asset for given location.

<div><strong>200</strong>: Success, appId info returned.</div> <div><strong>401</strong>: Authorization required or customer id not found</div> <div><strong>404</strong>: location id or url does not exist.</div> <div><strong>500</strong>: internal server error</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Customer id | 
  **locationId** | **string**|  | 

### Return type

[**NetworkConfig**](NetworkConfig.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomerPrototypePostCaptivePortalEnableGuestEmailCollection**
> interface{} CustomerPrototypePostCaptivePortalEnableGuestEmailCollection(ctx, id, locationId, networkId)
Patch the Captive Portal Network to be compliant for guest email collection.

<div><strong>200</strong>: Success, CaptivePortal Networks has been patched.</div> <div><strong>401</strong>: Authorization required or customer id not found</div> <div><strong>404</strong>: location id or secondary networks does not exist.</div> <div><strong>500</strong>: internal server error</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Customer id | 
  **locationId** | **string**|  | 
  **networkId** | **string**|  | 

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomerPrototypePostCaptivePortalNetworkUsageStats**
> []interface{} CustomerPrototypePostCaptivePortalNetworkUsageStats(ctx, id, locationId, networkId, optional)
Fetch the Captive Portal Network Usage stats for the given network.

<div><strong>200</strong>: Success, CaptivePortal Networks returned.</div> <div><strong>401</strong>: Authorization required or customer id not found</div> <div><strong>404</strong>: location id or secondary networks does not exist.</div> <div><strong>500</strong>: internal server error</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Customer id | 
  **locationId** | **string**|  | 
  **networkId** | **string**|  | 
 **optional** | ***CustomerApiCustomerPrototypePostCaptivePortalNetworkUsageStatsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CustomerApiCustomerPrototypePostCaptivePortalNetworkUsageStatsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **inclusions** | **optional.String**| Fields to include in response | 

### Return type

**[]interface{}**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomerPrototypePostCustomSharedSchedule**
> LocationCustomSchedule CustomerPrototypePostCustomSharedSchedule(ctx, id, locationId, name, type_, schedules)
Create \"custom shared\" schedules that shared by all persons and devices in a location.

<div><strong>200</strong>: Success, custom shared schedules applied.</div> <div><strong>400</strong>: Required fields missing or field type is incorrect.</div> <div><strong>401</strong>: Authorization required or customer id not found</div> <div><strong>404</strong>: Location id does not exist or is not known to Plume</div> <div><strong>422</strong>: schedules value is invalid.</div> <div><strong>500</strong>: Internal server error</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Customer id | 
  **locationId** | **string**|  | 
  **name** | **string**|  | 
  **type_** | **string**|  | 
  **schedules** | **string**|  | 

### Return type

[**LocationCustomSchedule**](LocationCustomSchedule.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomerPrototypePostDeviceFreeze**
> XAny CustomerPrototypePostDeviceFreeze(ctx, id, locationId, mac, freezeTemplateId)
Post a shared schedule uuid freeze for a device for a Location ID.

<div><strong>200</strong>: Success, updated.</div> <div><strong>401</strong>: Authorization required or customer id not found.</div> <div><strong>404</strong>: Location id, does not exist.</div> <div><strong>404</strong>: FreezeTemplateId not found</div> <div><strong>404</strong>: Device not found</div> <div><strong>422</strong>: GroupOfUnassignedDevices has active freeze schedule</div> <div><strong>422</strong>: Person has active freeze schedule</div> <div><strong>500</strong>: Internal server error</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Customer id | 
  **locationId** | **string**|  | 
  **mac** | **string**|  | 
  **freezeTemplateId** | **string**| Valid templates are uuids | 

### Return type

[**XAny**](x-any.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomerPrototypePostDeviceGroup**
> NetworkAccessDeviceGroup CustomerPrototypePostDeviceGroup(ctx, id, locationId, networkId, name, optional)
Create a named device group within a network and optionally specify member devices.

<div><strong>200</strong>: Success.</div> <div><strong>422</strong>: Schema validation failed.</div> <div><strong>404</strong>: Location does not exist.</div> <div><strong>404</strong>: Network does not exist.</div> <div><strong>403</strong>: Not allowed to create groups in unsupported networks.</div> <div><strong>401</strong>: Unauthorized.</div> <div><strong>400</strong>: Invalid JSON or missing arguments.</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Customer id | 
  **locationId** | **string**|  | 
  **networkId** | **string**|  | 
  **name** | **string**|  | 
 **optional** | ***CustomerApiCustomerPrototypePostDeviceGroupOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CustomerApiCustomerPrototypePostDeviceGroupOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **devices** | **optional.String**|  | 

### Return type

[**NetworkAccessDeviceGroup**](NetworkAccessDeviceGroup.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomerPrototypePostDeviceQos**
> []XAny CustomerPrototypePostDeviceQos(ctx, id, locationId, mac, prioritization)
Set QoS of a single device

<div><strong>202</strong>: Success.</div> <div><strong>404</strong>: Location does not exist.</div> <div><strong>422</strong>: Prioritization is not a valid value.</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Customer id | 
  **locationId** | **string**|  | 
  **mac** | **string**|  | 
  **prioritization** | **string**|  | 

### Return type

[**[]XAny**](x-any.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomerPrototypePostDeviceSecurityPolicyAnomalyExperience**
> interface{} CustomerPrototypePostDeviceSecurityPolicyAnomalyExperience(ctx, id, locationId, mac, fqdn)
Initiate an Anomaly Experience (demo) for a Device on a location.

<div><strong>200</strong>: Success.</div> <div><strong>400</strong>: Required fields missing or field type is incorrect.</div> <div><strong>401</strong>: Authorization required or customer id not found</div> <div><strong>404</strong>: Location id or WifiNetwork does not exist and is not known to Plume</div> <div><strong>500</strong>: Internal server error</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Customer id | 
  **locationId** | **string**|  | 
  **mac** | **string**|  | 
  **fqdn** | **string**|  | 

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomerPrototypePostDeviceSecurityPolicyAnomalyWhitelist**
> interface{} CustomerPrototypePostDeviceSecurityPolicyAnomalyWhitelist(ctx, id, locationId, mac, fqdn, optional)
Approve a previously blacklisted anomalous dns for a Device on a location.

<div><strong>200</strong>: Success.</div> <div><strong>400</strong>: Required fields missing or field type is incorrect.</div> <div><strong>401</strong>: Authorization required or customer id not found</div> <div><strong>404</strong>: Location id or WifiNetwork does not exist and is not known to Plume</div> <div><strong>422</strong>: DNS value is invalid.</div> <div><strong>500</strong>: Internal server error</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Customer id | 
  **locationId** | **string**|  | 
  **mac** | **string**|  | 
  **fqdn** | **string**|  | 
 **optional** | ***CustomerApiCustomerPrototypePostDeviceSecurityPolicyAnomalyWhitelistOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CustomerApiCustomerPrototypePostDeviceSecurityPolicyAnomalyWhitelistOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **reason** | **optional.String**|  | [default to trust]
 **ttl** | **optional.Float64**|  | 

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomerPrototypePostDeviceSecurityPolicyWebsitesBlacklist**
> interface{} CustomerPrototypePostDeviceSecurityPolicyWebsitesBlacklist(ctx, id, locationId, mac, optional)
Update a Device's Security Policy for a location ID to include a blacklisted DNS entry.

<div><strong>200</strong>: Success.</div> <div><strong>400</strong>: Required fields missing or field type is incorrect.</div> <div><strong>401</strong>: Authorization required or customer id not found</div> <div><strong>404</strong>: Location id, WifiNetwork, or Device does not exist and is not known to Plume</div> <div><strong>422</strong>: DNS value is invalid.</div> <div><strong>500</strong>: Internal server error</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Customer id | 
  **locationId** | **string**|  | 
  **mac** | **string**|  | 
 **optional** | ***CustomerApiCustomerPrototypePostDeviceSecurityPolicyWebsitesBlacklistOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CustomerApiCustomerPrototypePostDeviceSecurityPolicyWebsitesBlacklistOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **dns** | **optional.String**|  | 
 **type_** | **optional.String**|  | 
 **value** | **optional.String**|  | 
 **direction** | **optional.String**|  | 
 **geoLocation** | **optional.String**|  | 
 **endTimestamp** | **optional.Float64**| the end time stamp,  UTC unix epoch timestamp in ms | 
 **akamaiCategoryId** | **optional.Float64**| the akamai category id, number | 

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomerPrototypePostDeviceSecurityPolicyWebsitesWhitelist**
> interface{} CustomerPrototypePostDeviceSecurityPolicyWebsitesWhitelist(ctx, id, locationId, mac, optional)
Update a Device's Security Policy for a location ID to include a whitelisted DNS entry.

<div><strong>200</strong>: Success.</div> <div><strong>400</strong>: Required fields missing or field type is incorrect.</div> <div><strong>401</strong>: Authorization required or customer id not found</div> <div><strong>404</strong>: Location id, WifiNetwork, or Device does not exist and is not known to Plume</div> <div><strong>422</strong>: DNS value is invalid.</div> <div><strong>500</strong>: Internal server error</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Customer id | 
  **locationId** | **string**|  | 
  **mac** | **string**|  | 
 **optional** | ***CustomerApiCustomerPrototypePostDeviceSecurityPolicyWebsitesWhitelistOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CustomerApiCustomerPrototypePostDeviceSecurityPolicyWebsitesWhitelistOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **dns** | **optional.String**|  | 
 **type_** | **optional.String**|  | 
 **value** | **optional.String**|  | 
 **direction** | **optional.String**|  | 
 **geoLocation** | **optional.String**|  | 
 **eventType** | **optional.String**| EventType field from event response - can be &#39;kids&#39;, &#39;teenagers&#39;, &#39;secureAndProtect, etc&#39; | 
 **source** | **optional.String**| Source field from events response - can be &#39;brightcloud&#39;, &#39;webpulse&#39;, &#39;gatekeeper&#39;, &#39;gatekeeper-ohp&#39; | 
 **endTimestamp** | **optional.Float64**| the end time stamp,  UTC unix epoch timestamp in ms | 
 **akamaiCategoryId** | **optional.Float64**| the akamai category id, number | 

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomerPrototypePostDeviceToAccessZone**
> XAny CustomerPrototypePostDeviceToAccessZone(ctx, id, locationId, zoneId, mac)
Add a device mac to a WiFi Access Zone

<div><strong>200</strong>: Success, all access zones returned</div> <div><strong>400</strong>: Required fields missing</div> <div><strong>401</strong>: Authorization required or customer id not found</div> <div><strong>404</strong>: Customer id, location id, or WifiNetwork does not exist</div> <div><strong>422</strong>: Validation failed</div> <div><strong>500</strong>: Internal server error</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Customer id | 
  **locationId** | **string**|  | 
  **zoneId** | **float64**| id of access zone | 
  **mac** | **string**| the device mac to be added to the access zone | 

### Return type

[**XAny**](x-any.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomerPrototypePostFrontHaul**
> NetworkConfig CustomerPrototypePostFrontHaul(ctx, id, locationId, ssid, encryptionKey, optional)
Create a Front Haul Network for a Location ID.

<div><strong>200</strong>: Success.</div> <div><strong>401</strong>: Authorization required or customer id not found</div> <div><strong>404</strong>: Location id or NetworkId does not exist and is not known to Plume</div> <div><strong>422</strong>: NetworkId/SSIDs must be the unique and valid values.</div> <div><strong>500</strong>: Internal server error</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Customer id | 
  **locationId** | **string**|  | 
  **ssid** | **string**|  | 
  **encryptionKey** | **string**|  | 
 **optional** | ***CustomerApiCustomerPrototypePostFrontHaulOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CustomerApiCustomerPrototypePostFrontHaulOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **networkId** | **optional.String**|  | 
 **enable** | **optional.Bool**|  | [default to true]
 **accessZone** | **optional.String**|  | 
 **wpaMode** | **optional.String**|  | 
 **ssidBroadcast** | **optional.Bool**|  | 

### Return type

[**NetworkConfig**](NetworkConfig.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomerPrototypePostFrontHaulsDpp**
> interface{} CustomerPrototypePostFrontHaulsDpp(ctx, id, locationId, networkId, optional)
Create the DPP setting for a Fronthaul Network.

<div><strong>202</strong>: Success, new DPP configurator generated.</div> <div><strong>400</strong>: Required fields missing or field type is incorrect.</div> <div><strong>401</strong>: Authorization required or customer id not found</div> <div><strong>404</strong>: Location id or Fronthaul Network does not exist.</div> <div><strong>422</strong>: Invalid keys.</div> <div><strong>500</strong>: internal server error</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Customer id | 
  **locationId** | **string**|  | 
  **networkId** | **string**|  | 
 **optional** | ***CustomerApiCustomerPrototypePostFrontHaulsDppOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CustomerApiCustomerPrototypePostFrontHaulsDppOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **enabled** | **optional.Bool**| should we configure dpp for this network - defaults to true | 
 **curve** | **optional.String**| one of predefined elliptic curves, - optional,  if missing in request default to prime256v1 | 
 **privateKey** | **optional.String**| privateKey, must also provide public part if present, optional | 
 **publicKey** | **optional.String**| publicKey | 

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomerPrototypePostFrontHaulsDppBootstrap**
> interface{} CustomerPrototypePostFrontHaulsDppBootstrap(ctx, id, locationId, networkId, optional)
Create a bootstrap for DPP setting for a Fronthaul Network.

<div><strong>200</strong>: Success, new DPP configurator generated.</div> <div><strong>400</strong>: Required fields missing or field type is incorrect.</div> <div><strong>401</strong>: Authorization required or customer id not found</div> <div><strong>404</strong>: Location id or Fronthaul Network does not exist.</div> <div><strong>422</strong>: Invalid curve.</div> <div><strong>500</strong>: internal server error</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Customer id | 
  **locationId** | **string**|  | 
  **networkId** | **string**|  | 
 **optional** | ***CustomerApiCustomerPrototypePostFrontHaulsDppBootstrapOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CustomerApiCustomerPrototypePostFrontHaulsDppBootstrapOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **curve** | **optional.String**| one of predefined elliptic curves, - optional,  if missing in requset default to prime256v1 | 

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomerPrototypePostFrontHaulsDppEnrollment**
> interface{} CustomerPrototypePostFrontHaulsDppEnrollment(ctx, id, locationId, networkId, bootstrapUri)
Create an enrollment for DPP setting for a fronthaul secondary network.

<div><strong>202</strong>: Success, new DPP configurator generated.</div> <div><strong>400</strong>: Required fields missing or field type is incorrect.</div> <div><strong>401</strong>: Authorization required or customer id not found</div> <div><strong>404</strong>: Location id or wifi network does not exist.</div> <div><strong>404</strong>: Configurator keys for network not found.</div> <div><strong>500</strong>: internal server error</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Customer id | 
  **locationId** | **string**|  | 
  **networkId** | **string**|  | 
  **bootstrapUri** | **string**|  | 

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomerPrototypePostGroupOfUnassignedDevicesFreezeTemplateId**
> XAny CustomerPrototypePostGroupOfUnassignedDevicesFreezeTemplateId(ctx, id, locationId, freezeTemplateId)
POST GroupOfUnassignedDevices to be frozen for a Location ID.

<div><strong>200</strong>: Success, updated.</div> <div><strong>401</strong>: Authorization required or customer id not found.</div> <div><strong>404</strong>: Location id, does not exist.</div> <div><strong>404</strong>: Freeze Template Id not found.</div> <div><strong>409</strong>: Freeze Template Id already applied.</div> <div><strong>500</strong>: Internal server error</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Customer id | 
  **locationId** | **string**|  | 
  **freezeTemplateId** | **string**| Valid templates are uuids | 

### Return type

[**XAny**](x-any.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomerPrototypePostGroupOfUnassignedDevicesSecurityPolicyWebsitesBlacklist**
> interface{} CustomerPrototypePostGroupOfUnassignedDevicesSecurityPolicyWebsitesBlacklist(ctx, id, locationId, optional)
Update a Location's Default Device Group Security Policy for a location ID to include a blacklisted DNS entry.

<div><strong>200</strong>: Success.</div> <div><strong>400</strong>: Required fields missing or field type is incorrect.</div> <div><strong>401</strong>: Authorization required or customer id not found</div> <div><strong>404</strong>: Location id, WifiNetwork, or Device does not exist and is not known to Plume</div> <div><strong>422</strong>: DNS value is invalid.</div> <div><strong>500</strong>: Internal server error</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Customer id | 
  **locationId** | **string**|  | 
 **optional** | ***CustomerApiCustomerPrototypePostGroupOfUnassignedDevicesSecurityPolicyWebsitesBlacklistOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CustomerApiCustomerPrototypePostGroupOfUnassignedDevicesSecurityPolicyWebsitesBlacklistOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **dns** | **optional.String**|  | 
 **type_** | **optional.String**|  | 
 **value** | **optional.String**|  | 
 **direction** | **optional.String**|  | 
 **geoLocation** | **optional.String**|  | 
 **endTimestamp** | **optional.Float64**| the end time stamp,  UTC unix epoch timestamp in ms | 
 **akamaiCategoryId** | **optional.Float64**| the akamai category id, number | 

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomerPrototypePostGroupOfUnassignedDevicesSecurityPolicyWebsitesWhitelist**
> interface{} CustomerPrototypePostGroupOfUnassignedDevicesSecurityPolicyWebsitesWhitelist(ctx, id, locationId, optional)
Update a Location's Default Device Group Security Policy for a location ID to include a whitelisted DNS entry.

<div><strong>200</strong>: Success.</div> <div><strong>400</strong>: Required fields missing or field type is incorrect.</div> <div><strong>401</strong>: Authorization required or customer id not found</div> <div><strong>404</strong>: Location id, WifiNetwork, or Device does not exist and is not known to Plume</div> <div><strong>422</strong>: DNS value is invalid.</div> <div><strong>500</strong>: Internal server error</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Customer id | 
  **locationId** | **string**|  | 
 **optional** | ***CustomerApiCustomerPrototypePostGroupOfUnassignedDevicesSecurityPolicyWebsitesWhitelistOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CustomerApiCustomerPrototypePostGroupOfUnassignedDevicesSecurityPolicyWebsitesWhitelistOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **dns** | **optional.String**|  | 
 **type_** | **optional.String**|  | 
 **value** | **optional.String**|  | 
 **direction** | **optional.String**|  | 
 **geoLocation** | **optional.String**|  | 
 **endTimestamp** | **optional.Float64**| the end time stamp,  UTC unix epoch timestamp in ms | 
 **akamaiCategoryId** | **optional.Float64**| the akamai category id, number | 

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomerPrototypePostKvConfigs**
> KvConfig CustomerPrototypePostKvConfigs(ctx, id, locationId, nodeId, module, key, value, optional)
Retrieve all kvConfigs on a particular Node for a Location ID.

<div><strong>200</strong>: Success, your new info looks good.</div> <div><strong>401</strong>: Authorization required or customer id not found</div> <div><strong>404</strong>: location id does not exist.</div> <div><strong>422</strong>: nodeId must be defined.</div> <div><strong>425</strong>: nodeId must belong to the location.</div> <div><strong>500</strong>: internal server error</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Customer id | 
  **locationId** | **string**|  | 
  **nodeId** | **string**|  | 
  **module** | **string**|  | 
  **key** | **string**|  | 
  **value** | **string**|  | 
 **optional** | ***CustomerApiCustomerPrototypePostKvConfigsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CustomerApiCustomerPrototypePostKvConfigsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------






 **persist** | **optional.Bool**|  | [default to false]

### Return type

[**KvConfig**](KvConfig.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomerPrototypePostLocationSecurityPolicyOHPDeviceSetup**
> interface{} CustomerPrototypePostLocationSecurityPolicyOHPDeviceSetup(ctx, id, locationId, optional)
Setup a Mobile Device for Security Out of Home Protection (returns a Deeplink for use with Mobolize).

<div><strong>200</strong>: Success.</div> <div><strong>400</strong>: Required fields missing or field type is incorrect.</div> <div><strong>401</strong>: Authorization required or customer id not found</div> <div><strong>404</strong>: Location id or WifiNetwork does not exist and is not known to Plume</div> <div><strong>500</strong>: Internal server error</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Customer id | 
  **locationId** | **string**|  | 
 **optional** | ***CustomerApiCustomerPrototypePostLocationSecurityPolicyOHPDeviceSetupOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CustomerApiCustomerPrototypePostLocationSecurityPolicyOHPDeviceSetupOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **lanIpv4** | **optional.String**| Mobile device lanIpv4 address, if any | 
 **lanIpv6** | **optional.String**| Mobile device lanIpv6 address, if any | 

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomerPrototypePostLocationSecurityPolicyWebsitesBlacklist**
> interface{} CustomerPrototypePostLocationSecurityPolicyWebsitesBlacklist(ctx, id, locationId, optional)
Update a Location's Security Policy for a location ID to include a blacklisted DNS entry.

<div><strong>200</strong>: Success.</div> <div><strong>400</strong>: Required fields missing or field type is incorrect.</div> <div><strong>401</strong>: Authorization required or customer id not found</div> <div><strong>404</strong>: Location id, WifiNetwork, or Device does not exist and is not known to Plume</div> <div><strong>422</strong>: DNS value is invalid.</div> <div><strong>500</strong>: Internal server error</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Customer id | 
  **locationId** | **string**|  | 
 **optional** | ***CustomerApiCustomerPrototypePostLocationSecurityPolicyWebsitesBlacklistOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CustomerApiCustomerPrototypePostLocationSecurityPolicyWebsitesBlacklistOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **dns** | **optional.String**|  | 
 **type_** | **optional.String**|  | 
 **value** | **optional.String**|  | 
 **direction** | **optional.String**|  | 
 **geoLocation** | **optional.String**|  | 
 **endTimestamp** | **optional.Float64**| the end time stamp,  UTC unix epoch timestamp in ms | 
 **akamaiCategoryId** | **optional.Float64**| the akamai category id, number | 

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomerPrototypePostLocationSecurityPolicyWebsitesWhitelist**
> interface{} CustomerPrototypePostLocationSecurityPolicyWebsitesWhitelist(ctx, id, locationId, optional)
Update a Location's Security Policy for a location ID to include a whitelisted DNS entry.

<div><strong>200</strong>: Success.</div> <div><strong>400</strong>: Required fields missing or field type is incorrect.</div> <div><strong>401</strong>: Authorization required or customer id not found</div> <div><strong>404</strong>: Location id, WifiNetwork, or Device does not exist and is not known to Plume</div> <div><strong>422</strong>: DNS value is invalid.</div> <div><strong>500</strong>: Internal server error</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Customer id | 
  **locationId** | **string**|  | 
 **optional** | ***CustomerApiCustomerPrototypePostLocationSecurityPolicyWebsitesWhitelistOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CustomerApiCustomerPrototypePostLocationSecurityPolicyWebsitesWhitelistOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **dns** | **optional.String**|  | 
 **type_** | **optional.String**|  | 
 **value** | **optional.String**|  | 
 **direction** | **optional.String**|  | 
 **geoLocation** | **optional.String**|  | 
 **eventType** | **optional.String**| EventType field from events response - can be &#39;kids&#39;, &#39;teenagers&#39;, &#39;secureAndProtect&#39;, etc | 
 **source** | **optional.String**| Source field from events response - can be &#39;brightcloud&#39;, &#39;webpulse&#39;, &#39;gatekeeper&#39;, &#39;gatekeeper-ohp&#39; | 
 **endTimestamp** | **optional.Float64**| the end time stamp,  UTC unix epoch timestamp in ms | 
 **akamaiCategoryId** | **optional.Float64**| the akamai category id, number | 

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomerPrototypePostManager**
> LocationAccess CustomerPrototypePostManager(ctx, id, email, name, locationId, optional)
Assign a manager to your location 

<div><strong>200</strong>: Success.</div> <div><strong>400</strong>: Required fields missing or field type is incorrect.</div> <div><strong>404</strong>: Location does not exist.</div> <div><strong>422</strong>: Invalid email, name, access type or manager is already assigned to this location </div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Customer id | 
  **email** | **string**|  | 
  **name** | **string**|  | 
  **locationId** | **string**|  | 
 **optional** | ***CustomerApiCustomerPrototypePostManagerOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CustomerApiCustomerPrototypePostManagerOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **accessType** | **optional.String**|  | 
 **notificationOptions** | **optional.String**|  | 

### Return type

[**LocationAccess**](LocationAccess.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomerPrototypePostOnboardingCheckpoint**
> InlineResponse2009 CustomerPrototypePostOnboardingCheckpoint(ctx, id, locationId, optional)
Record the new Onboarding Checkpoint for the Location ID.

<div><strong>200</strong>: Success, most recent checkpoint saved.</div> <div><strong>401</strong>: Authorization required or customer id not found</div> <div><strong>404</strong>: location id does not exist and is not known to Plume</div> <div><strong>422</strong>: checkpoint value must be defined.</div> <div><strong>500</strong>: internal server error</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Customer id | 
  **locationId** | **string**|  | 
 **optional** | ***CustomerApiCustomerPrototypePostOnboardingCheckpointOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CustomerApiCustomerPrototypePostOnboardingCheckpointOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **checkpoint** | **optional.String**| is the last passed onboarding step by the customer: &#39;PodsAdded&#39; or &#39;OnboardingComplete&#39;; | 
 **podsSeenByBle** | **optional.String**| is the number of Nodes the app discovered by BLE when the onboarding was completed by the customer, submit with PodsAdded | 
 **appOs** | **optional.String**| is the version of the app used during the onboarding, submit with PodsAdded | 
 **osVersion** | **optional.String**| is the phone OS version used during the onboarding, submit with PodsAdded | 

### Return type

[**InlineResponse2009**](inline_response_200_9.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomerPrototypePostPersonFreeze**
> XAny CustomerPrototypePostPersonFreeze(ctx, id, locationId, personId, freezeTemplateId)
Post a shared schedule uuid freeze for a person for a Location ID.

<div><strong>200</strong>: Success, updated.</div> <div><strong>401</strong>: Authorization required or customer id not found.</div> <div><strong>404</strong>: Location id, does not exist.</div> <div><strong>404</strong>: FreezeTemplateId not found.</div> <div><strong>404</strong>: Person not found.</div> <div><strong>500</strong>: Internal server error</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Customer id | 
  **locationId** | **string**|  | 
  **personId** | **string**|  | 
  **freezeTemplateId** | **string**| Valid templates are uuids | 

### Return type

[**XAny**](x-any.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomerPrototypePostPersonSecurityPolicyWebsitesBlacklist**
> interface{} CustomerPrototypePostPersonSecurityPolicyWebsitesBlacklist(ctx, id, locationId, personId, optional)
Update a Person's Security Policy for a location ID to include a blacklisted DNS entry.

<div><strong>200</strong>: Success.</div> <div><strong>400</strong>: Required fields missing or field type is incorrect.</div> <div><strong>401</strong>: Authorization required or customer id not found</div> <div><strong>404</strong>: Location id, WifiNetwork, or Person id does not exist and is not known to Plume</div> <div><strong>422</strong>: DNS value is invalid.</div> <div><strong>500</strong>: Internal server error</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Customer id | 
  **locationId** | **string**|  | 
  **personId** | **string**|  | 
 **optional** | ***CustomerApiCustomerPrototypePostPersonSecurityPolicyWebsitesBlacklistOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CustomerApiCustomerPrototypePostPersonSecurityPolicyWebsitesBlacklistOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **dns** | **optional.String**|  | 
 **type_** | **optional.String**|  | 
 **value** | **optional.String**|  | 
 **direction** | **optional.String**|  | 
 **geoLocation** | **optional.String**|  | 
 **endTimestamp** | **optional.Float64**| the end time stamp,  UTC unix epoch timestamp in ms | 
 **akamaiCategoryId** | **optional.Float64**| the akamai category id, number | 

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomerPrototypePostPersonSecurityPolicyWebsitesWhitelist**
> interface{} CustomerPrototypePostPersonSecurityPolicyWebsitesWhitelist(ctx, id, locationId, personId, optional)
Update a Person's Security Policy for a location ID to include a whitelisted DNS entry.

<div><strong>200</strong>: Success.</div> <div><strong>400</strong>: Required fields missing or field type is incorrect.</div> <div><strong>401</strong>: Authorization required or customer id not found</div> <div><strong>404</strong>: Location id, WifiNetwork, or Person id does not exist and is not known to Plume</div> <div><strong>422</strong>: DNS value is invalid.</div> <div><strong>500</strong>: Internal server error</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Customer id | 
  **locationId** | **string**|  | 
  **personId** | **string**|  | 
 **optional** | ***CustomerApiCustomerPrototypePostPersonSecurityPolicyWebsitesWhitelistOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CustomerApiCustomerPrototypePostPersonSecurityPolicyWebsitesWhitelistOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **dns** | **optional.String**|  | 
 **type_** | **optional.String**|  | 
 **value** | **optional.String**|  | 
 **direction** | **optional.String**|  | 
 **geoLocation** | **optional.String**|  | 
 **eventType** | **optional.String**| EventType field from events response - can be &#39;kids&#39;, &#39;teenagers&#39;, &#39;secureAndProtect&#39;, etc | 
 **source** | **optional.String**| Source field from events response - can be &#39;brightcloud&#39;, &#39;webpulse&#39;, &#39;gatekeeper&#39;, &#39;gatekeeper-ohp&#39; | 
 **endTimestamp** | **optional.Float64**| the end time stamp,  UTC unix epoch timestamp in ms | 
 **akamaiCategoryId** | **optional.Float64**| the akamai category id, number | 

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomerPrototypePostPersons**
> Person CustomerPrototypePostPersons(ctx, id, locationId, nickname, imageId, optional)
Create  a Person for a Location ID.

<div><strong>200</strong>: Success.</div> <div><strong>401</strong>: Authorization required or customer id not found</div> <div><strong>404</strong>: Location id does not exist and is not known to Plume</div> <div><strong>422</strong>: Nickname must be defined and mac addresses must be valid and email needs to be provided when permission is provided.</div> <div><strong>500</strong>: Internal server error</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Customer id | 
  **locationId** | **string**|  | 
  **nickname** | **string**|  | 
  **imageId** | **string**| unique identifier for referencing a Person&#39;s hosted profile image, defaults are PROFILE_MAN and PROFILE_WOMAN | 
 **optional** | ***CustomerApiCustomerPrototypePostPersonsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CustomerApiCustomerPrototypePostPersonsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **assignedDevices** | **optional.String**| mac addresses of devices assigned to this Person | 
 **profile** | **optional.String**| Profile object contains &#39;type&#39; field - valid values: &#39;employee&#39; | 
 **email** | **optional.String**| email | 
 **permission** | **optional.String**| Permission object for creating a manager for the location | 
 **serviceLinking** | **optional.String**| serviceLinking object that links this Person object to a 3rd party&#39;s Person | 

### Return type

[**Person**](Person.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomerPrototypePostPortForward**
> PortForward CustomerPrototypePostPortForward(ctx, id, locationId, mac, optional)
Record a new Port Forwarding entry for an existing DHCP IP reservation tied to a MAC address at a Location ID.

<div><strong>200</strong>: Success, all PortForwards are returned.</div> <div><strong>401</strong>: Authorization required or customer id not found</div> <div><strong>422</strong>: networkConfiguration, dhcpReservation, PortForward is empty.</div> <div><strong>422</strong>: mac is empty, or invalid, externalPort/internalPort is out of range, or protocol is invalid, or duplicate externalPort.</div> <div><strong>500</strong>: internal server error</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Customer id | 
  **locationId** | **string**|  | 
  **mac** | **string**|  | 
 **optional** | ***CustomerApiCustomerPrototypePostPortForwardOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CustomerApiCustomerPrototypePostPortForwardOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **externalPort** | **optional.String**|  | 
 **internalPort** | **optional.String**|  | 
 **protocol** | **optional.String**|  | 
 **name** | **optional.String**|  | 

### Return type

[**PortForward**](PortForward.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomerPrototypePostRemoteConnectionsAllow**
> interface{} CustomerPrototypePostRemoteConnectionsAllow(ctx, id, locationId, mac, type_, value, expiresAt)
Post a Remote Connection Allow IpAddress/ttl for the given device and Location ID.

<div><strong>200</strong>: Success.</div> <div><strong>400</strong>: Required fields missing.</div> <div><strong>401</strong>: Authorization required or customer id not found</div> <div><strong>404</strong>: Location id or Device mac does not exist and is not known to Plume</div> <div><strong>422</strong>: Fields have an invalid type or value.</div> <div><strong>500</strong>: Internal server error</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Customer id | 
  **locationId** | **string**|  | 
  **mac** | **string**|  | 
  **type_** | **string**| either ipv4 or ipv6 | 
  **value** | **string**| ipaddress | 
  **expiresAt** | **string**| UTC timestamp in ISO 8601 format | 

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomerPrototypePostRemoteConnectionsAllowAll**
> interface{} CustomerPrototypePostRemoteConnectionsAllowAll(ctx, id, locationId, mac, expiresAt)
Post a Remote Connection Allow All/ttl for the given device and Location ID.

<div><strong>200</strong>: Success.</div> <div><strong>400</strong>: Required fields missing.</div> <div><strong>401</strong>: Authorization required or customer id not found</div> <div><strong>404</strong>: Location id or Device mac does not exist and is not known to Plume</div> <div><strong>422</strong>: Fields have an invalid type or value.</div> <div><strong>500</strong>: Internal server error</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Customer id | 
  **locationId** | **string**|  | 
  **mac** | **string**|  | 
  **expiresAt** | **string**| UTC timestamp in ISO 8601 format | 

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomerPrototypePostRooms**
> Room CustomerPrototypePostRooms(ctx, id, locationId, name, optional)
Create a Room for a Location ID.

<div><strong>200</strong>: Success.</div> <div><strong>401</strong>: Authorization required or customer id not found</div> <div><strong>404</strong>: Location id does not exist and is not known to Plume</div> <div><strong>422</strong>: Devices and Nodes must be defined and mac addresses must be valid.</div> <div><strong>500</strong>: Internal server error</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Customer id | 
  **locationId** | **string**|  | 
  **name** | **string**| name of this Room | 
 **optional** | ***CustomerApiCustomerPrototypePostRoomsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CustomerApiCustomerPrototypePostRoomsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **devices** | **optional.String**| mac addresses of devices assigned to this Room | 
 **nodes** | **optional.String**| nodeIds assigned to this Room | 

### Return type

[**Room**](Room.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomerPrototypePostRunMobileIspSpeedTest**
> interface{} CustomerPrototypePostRunMobileIspSpeedTest(ctx, id, locationId, optional)
Run ISP speed test for GW node on mobile request.

<div><strong>200</strong>: Success, run.</div> <div><strong>401</strong>: Authorization required.</div> <div><strong>404</strong>: Customer or location does not exists.</div> <div><strong>500</strong>: Internal server error</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Customer id | 
  **locationId** | **string**|  | 
 **optional** | ***CustomerApiCustomerPrototypePostRunMobileIspSpeedTestOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CustomerApiCustomerPrototypePostRunMobileIspSpeedTestOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **requestId** | **optional.String**|  | 
 **serverId** | **optional.Float64**|  | 
 **uplinkType** | **optional.String**|  | [default to wire]

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomerPrototypePostSpeedTest**
> interface{} CustomerPrototypePostSpeedTest(ctx, id, locationId, nodeId, testType, optional)
Run speed test for a node.

<div><strong>200</strong>: Success, run.</div> <div><strong>400</strong>: Required fields missing.</div> <div><strong>401</strong>: Authorization required.</div> <div><strong>404</strong>: Customer, location or node does not exists.</div> <div><strong>422</strong>: Invalid test type.</div> <div><strong>500</strong>: Internal server error</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Customer id | 
  **locationId** | **string**|  | 
  **nodeId** | **string**|  | 
  **testType** | **string**|  | [default to OOKLA]
 **optional** | ***CustomerApiCustomerPrototypePostSpeedTestOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CustomerApiCustomerPrototypePostSpeedTestOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **serverId** | **optional.Float64**|  | 
 **uplinkType** | **optional.String**|  | [default to wire]

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomerPrototypePostWhitelistApprovalRequest**
> interface{} CustomerPrototypePostWhitelistApprovalRequest(ctx, id, locationId, value, type_)
Post a request for a whitelist exception to be added to your person profile.

<div><strong>200</strong>: Success.</div> <div><strong>400</strong>: Required fields missing or field type is incorrect.</div> <div><strong>401</strong>: Authorization required or customer id not found</div> <div><strong>404</strong>: Location id, CustomerId or requst id does not exist and is not known to Plume</div> <div><strong>500</strong>: Internal server error</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Customer id | 
  **locationId** | **string**|  | 
  **value** | **string**|  | 
  **type_** | **string**|  | 

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomerPrototypePostWifiAccessZone**
> []WifiAccessZone CustomerPrototypePostWifiAccessZone(ctx, id, locationId, description, type_, optional)
Create a new WiFi Access Zone

<div><strong>200</strong>: Success, all access zones returned</div> <div><strong>400</strong>: Required fields missing</div> <div><strong>401</strong>: Authorization required or customer id not found</div> <div><strong>404</strong>: Customer id, location id, or WifiNetwork does not exist</div> <div><strong>422</strong>: Validation failed</div> <div><strong>500</strong>: Internal server error</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Customer id | 
  **locationId** | **string**|  | 
  **description** | **string**| name of access zone | 
  **type_** | **string**| for now, must be &#39;guests&#39; | 
 **optional** | ***CustomerApiCustomerPrototypePostWifiAccessZoneOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CustomerApiCustomerPrototypePostWifiAccessZoneOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **accessibleDevices** | **optional.String**| macs of home devices visible to this guest access zone | 

### Return type

[**[]WifiAccessZone**](WifiAccessZone.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomerPrototypePostWifiKey**
> []WifiNetworkKey CustomerPrototypePostWifiKey(ctx, id, locationId, accessZone, encryptionKey, enable, format, optional)
Create a new WiFi Password

<div><strong>200</strong>: Success, all passwords returned</div> <div><strong>400</strong>: Required fields missing</div> <div><strong>401</strong>: Authorization required or customer id not found</div> <div><strong>404</strong>: Customer id, location id, or WifiNetwork does not exist</div> <div><strong>422</strong>: Password validation failed</div> <div><strong>500</strong>: Internal server error</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Customer id | 
  **locationId** | **string**|  | 
  **accessZone** | **string**| home | guests | internetAccessOnly | 
  **encryptionKey** | **string**|  | 
  **enable** | **bool**| devices can connect using this encryptionKey | 
  **format** | **string**| encryptionKey | phoneNumber | 
 **optional** | ***CustomerApiCustomerPrototypePostWifiKeyOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CustomerApiCustomerPrototypePostWifiKeyOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------






 **expiresAt** | **optional.Time**| UTC in ISO 8601 String format | 
 **content** | **optional.String**| Valid values: &#39;adultAndSensitive&#39; | 

### Return type

[**[]WifiNetworkKey**](WifiNetworkKey.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomerPrototypePostWifiNetwork**
> InlineResponse2002 CustomerPrototypePostWifiNetwork(ctx, id, locationId, optional)
Update a WiFi SSID and/or PSK for a Location ID.

<div><strong>200</strong>: Success, in your future 100 mbps for all devices I see. -Yoda.</div> <div><strong>401</strong>: Authorization required or customer id not found</div> <div><strong>404</strong>: location id does not exist and is not known to Plume</div> <div><strong>409</strong>: A WifiNetwork already exists for this location.</div> <div><strong>422</strong>: encryptionKey or ssid must be defined, or key length < 8.</div> <div><strong>500</strong>: internal server error</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Customer id | 
  **locationId** | **string**|  | 
 **optional** | ***CustomerApiCustomerPrototypePostWifiNetworkOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CustomerApiCustomerPrototypePostWifiNetworkOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **encryptionKey** | **optional.String**| Needs to be a minimum of 8 characters | 
 **ssid** | **optional.String**|  | 
 **wpaMode** | **optional.String**| psk-mixed (WPA+WPA2) || sae-mixed (WPA2+WPA3) || psk2 (WPA2 only) || sae (WPA3 only) | 

### Return type

[**InlineResponse2002**](inline_response_200_2.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomerPrototypePostWifiNetworkDpp**
> interface{} CustomerPrototypePostWifiNetworkDpp(ctx, id, locationId, optional)
Create the DPP setting for a Location ID.

<div><strong>202</strong>: Success, new DPP configurator generated.</div> <div><strong>400</strong>: Required fields missing or field type is incorrect.</div> <div><strong>401</strong>: Authorization required or customer id not found</div> <div><strong>500</strong>: internal server error</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Customer id | 
  **locationId** | **string**|  | 
 **optional** | ***CustomerApiCustomerPrototypePostWifiNetworkDppOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CustomerApiCustomerPrototypePostWifiNetworkDppOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **enabled** | **optional.Bool**| should we configure dpp for this network - defaults to true | 
 **curve** | **optional.String**| one of predefined elliptic curves, - optional,  if missing in request default to prime256v1 | 
 **privateKey** | **optional.String**| privateKey, must also provide public part if present, optional | 
 **publicKey** | **optional.String**| publicKey | 

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomerPrototypePostWifiNetworkDppBootstrap**
> interface{} CustomerPrototypePostWifiNetworkDppBootstrap(ctx, id, locationId, optional)
Create a bootstrap for DPP setting for a wifi network.

<div><strong>200</strong>: Success, new DPP configurator generated.</div> <div><strong>400</strong>: Required fields missing or field type is incorrect.</div> <div><strong>401</strong>: Authorization required or customer id not found</div> <div><strong>404</strong>: Location id or wifi network does not exist.</div> <div><strong>422</strong>: Invalid curve.</div> <div><strong>500</strong>: internal server error</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Customer id | 
  **locationId** | **string**|  | 
 **optional** | ***CustomerApiCustomerPrototypePostWifiNetworkDppBootstrapOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CustomerApiCustomerPrototypePostWifiNetworkDppBootstrapOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **curve** | **optional.String**| one of predefined elliptic curves, - optional,  if missing in requset default to prime256v1 | 

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomerPrototypePostWifiNetworkDppEnrollment**
> interface{} CustomerPrototypePostWifiNetworkDppEnrollment(ctx, id, locationId, bootstrapUri)
Create an enrollment for DPP setting for a wifi network.

<div><strong>202</strong>: Success, new DPP configurator generated.</div> <div><strong>400</strong>: Required fields missing or field type is incorrect.</div> <div><strong>401</strong>: Authorization required or customer id not found</div> <div><strong>404</strong>: Location id or wifi network does not exist.</div> <div><strong>404</strong>: Configurator keys for network not found.</div> <div><strong>500</strong>: internal server error</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Customer id | 
  **locationId** | **string**|  | 
  **bootstrapUri** | **string**|  | 

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomerPrototypePublishSlowChangingDimensionConfigs**
> interface{} CustomerPrototypePublishSlowChangingDimensionConfigs(ctx, id)
Publish all slow changing dimension Kafka messages

<div><strong>204</strong>: Success.</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Customer id | 

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomerPrototypePutAuthorizationsPutCustomersidLocationslocationIdAuthorizations**
> Authorizations CustomerPrototypePutAuthorizationsPutCustomersidLocationslocationIdAuthorizations(ctx, id, locationId, optional)
Configure number of authorized leaf pods for a Location ID.

<div><strong>200</strong>: Success, updated.</div> <div><strong>400</strong>: Required fields are missing.</div> <div><strong>401</strong>: Authorization required or customer id not found.</div> <div><strong>404</strong>: Location id, does not exist.</div> <div><strong>500</strong>: Internal server error</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Customer id | 
  **locationId** | **string**|  | 
 **optional** | ***CustomerApiCustomerPrototypePutAuthorizationsPutCustomersidLocationslocationIdAuthorizationsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CustomerApiCustomerPrototypePutAuthorizationsPutCustomersidLocationslocationIdAuthorizationsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **numPodsAuthorized** | **optional.Float64**| number of leaf pods that are authorized to be claimed and be a part of the Plume network | 

### Return type

[**Authorizations**](Authorizations.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomerPrototypePutAuthorizationsPutCustomersidLocationslocationIdNodeAuthorizations**
> Authorizations CustomerPrototypePutAuthorizationsPutCustomersidLocationslocationIdNodeAuthorizations(ctx, id, locationId, optional)
Configure number of authorized leaf pods grouped by model id for a Location ID.

<div><strong>200</strong>: Success, updated.</div> <div><strong>400</strong>: Required fields are missing.</div> <div><strong>401</strong>: Authorization required or customer id not found.</div> <div><strong>404</strong>: Location id, does not exist.</div> <div><strong>500</strong>: Internal server error</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Customer id | 
  **locationId** | **string**|  | 
 **optional** | ***CustomerApiCustomerPrototypePutAuthorizationsPutCustomersidLocationslocationIdNodeAuthorizationsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CustomerApiCustomerPrototypePutAuthorizationsPutCustomersidLocationslocationIdNodeAuthorizationsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **numNodesAuthorized** | **optional.String**| number of leaf pods grouped by model id that are authorized to be claimed and be a part of the Plume network | 

### Return type

[**Authorizations**](Authorizations.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomerPrototypePutBackhaul**
> LocationBackhaul CustomerPrototypePutBackhaul(ctx, id, locationId, optional)
Toggle secure backhaul for a Location ID.

<div><strong>200</strong>: Success, updated.</div> <div><strong>400</strong>: Required fields missing.</div> <div><strong>401</strong>: Authorization required or customer id not found.</div> <div><strong>404</strong>: Location id, does not exist.</div> <div><strong>500</strong>: Internal server error</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Customer id | 
  **locationId** | **string**|  | 
 **optional** | ***CustomerApiCustomerPrototypePutBackhaulOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CustomerApiCustomerPrototypePutBackhaulOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **mode** | **optional.String**| auto || enable || disable | 
 **dynamicBeacon** | **optional.String**| A valid state for the dynamic beaconing setting. Either auto, enable, or disable | 
 **wds** | **optional.String**| auto || enable || disable | 
 **wpaMode** | **optional.String**| auto || psk2 || sae-mixed | 

### Return type

[**LocationBackhaul**](LocationBackhaul.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomerPrototypePutBandSteering**
> LocationBandSteering CustomerPrototypePutBandSteering(ctx, id, locationId, auto)
Enable/disable band steering for a Location ID (deprecated)

<div><strong>200</strong>: Success, updated.</div> <div><strong>400</strong>: Required fields missing.</div> <div><strong>401</strong>: Authorization required or customer id not found.</div> <div><strong>404</strong>: Location id, does not exist.</div> <div><strong>500</strong>: Internal server error</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Customer id | 
  **locationId** | **string**|  | 
  **auto** | **bool**|  | [default to true]

### Return type

[**LocationBandSteering**](LocationBandSteering.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomerPrototypePutBleMode**
> InlineResponse2003 CustomerPrototypePutBleMode(ctx, id, locationId, optional)
Enable or Disable BLE beaconing for all Pods at a location for Pod location services (e.g. for Pods Naming).

<div>With the mode of \"on\", all connected pods at this location will have their bluetooth beacon turned on for locating purposes. Each BLE beacon contains the serial number of the transmitting Pod. A setting of \"off\", turns off the BLE beaconing for all Pods. With mode set to \"wps\", all connected pods at this location will have their bluetooth beacon turned on for WPS related proximity measurements.</div> <div><strong>200</strong>: Success, your new info looks good.</div> <div><strong>401</strong>: Authorization required or customer id not found</div> <div><strong>404</strong>: location id does not exist.</div> <div><strong>422</strong>: bleMode must be defined.</div> <div><strong>500</strong>: internal server error</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Customer id | 
  **locationId** | **string**|  | 
 **optional** | ***CustomerApiCustomerPrototypePutBleModeOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CustomerApiCustomerPrototypePutBleModeOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **mode** | **optional.String**| on/off/wps/connectable | 

### Return type

[**InlineResponse2003**](inline_response_200_3.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomerPrototypePutBleModeForNode**
> InlineResponse2003 CustomerPrototypePutBleModeForNode(ctx, id, locationId, nodeId, optional)
Enable or Disable BLE beaconing for the specific Pod at a location.

<div>With the mode of \"on\", all connected pods at this location will have their bluetooth beacon turned on for locating purposes. Each BLE beacon contains the serial number of the transmitting Pod. A setting of \"off\", turns off the BLE beaconing for all Pods. With mode set to \"wps\", all connected pods at this location will have their bluetooth beacon turned on for WPS related proximity measurements.</div> <div><strong>200</strong>: Success, your new info looks good.</div> <div><strong>401</strong>: Authorization required or customer id not found</div> <div><strong>404</strong>: location id does not exist.</div> <div><strong>422</strong>: bleMode must be defined.</div> <div><strong>500</strong>: internal server error</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Customer id | 
  **locationId** | **string**|  | 
  **nodeId** | **string**|  | 
 **optional** | ***CustomerApiCustomerPrototypePutBleModeForNodeOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CustomerApiCustomerPrototypePutBleModeForNodeOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **mode** | **optional.String**| on/off/wps/connectable | 

### Return type

[**InlineResponse2003**](inline_response_200_3.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomerPrototypePutCaptivePortalSendDetails**
> CustomerPrototypePutCaptivePortalSendDetails(ctx, id, locationId, networkId, optional)
Send Captive Portal Guest details to requesters email for a given Location ID/NetworkId.

<div><strong>204</strong>: Success, updated.</div> <div><strong>401</strong>: Authorization required or customer id not found</div> <div><strong>404</strong>: location id or CaptivePortal Network does not exist.</div> <div><strong>500</strong>: internal server error</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Customer id | 
  **locationId** | **string**|  | 
  **networkId** | **string**|  | 
 **optional** | ***CustomerApiCustomerPrototypePutCaptivePortalSendDetailsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CustomerApiCustomerPrototypePutCaptivePortalSendDetailsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **duration** | **optional.Float64**| number of days for how far back in history for data | 
 **limit** | **optional.Float64**| limit how many emails we wish to return | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomerPrototypePutCommand**
> CustomerPrototypePutCommand(ctx, id, locationId, providerId)
Returns cloud migration status for customer

<div><strong>204</strong>: Success, no content.</div> <div><strong>400</strong>: Missing providerId body parameter</div> <div><strong>401</strong>: Authorization required or customer id not found</div> <div><strong>409</strong>: Accounts are already linked for providerId</div> <div><strong>422</strong>: Invalid providerId</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Customer id | 
  **locationId** | **string**|  | 
  **providerId** | **string**| enum to identify provider ex. commandAlexa | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomerPrototypePutControlMode**
> LocationControlMode CustomerPrototypePutControlMode(ctx, id, locationId, mode)
Set control mode for a Location ID.

<div><strong>200</strong>: Success, updated.</div> <div><strong>400</strong>: Required fields missing.</div> <div><strong>401</strong>: Authorization required or customer id not found.</div> <div><strong>404</strong>: Location id, does not exist.</div> <div><strong>500</strong>: Internal server error</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Customer id | 
  **locationId** | **string**|  | 
  **mode** | **string**|  | 

### Return type

[**LocationControlMode**](LocationControlMode.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomerPrototypePutCouncilmanResync**
> CustomerPrototypePutCouncilmanResync(ctx, id, locationId)
Push Security Configurations to Councilman.

<div><strong>204</strong>: Success.</div> <div><strong>401</strong>: Authorization required or customer id not found.</div> <div><strong>404</strong>: Location id, does not exist.</div> <div><strong>500</strong>: Internal server error</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Customer id | 
  **locationId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomerPrototypePutDeviceFreeze**
> XAny CustomerPrototypePutDeviceFreeze(ctx, id, locationId, mac, freezeTemplateId, optional)
Put a device to be frozen for a Location ID.

<div><strong>200</strong>: Success, updated.</div> <div><strong>401</strong>: Authorization required or customer id not found.</div> <div><strong>404</strong>: Location id, does not exist.</div> <div><strong>500</strong>: Internal server error</div> <div><strong>501</strong>: Not Implemented, if location is utilizing shared location freeze schedules</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Customer id | 
  **locationId** | **string**|  | 
  **mac** | **string**|  | 
  **freezeTemplateId** | **string**| Valid templates are &#39;untilMidinight&#39;, &#39;schoolNights&#39;, etc. | 
 **optional** | ***CustomerApiCustomerPrototypePutDeviceFreezeOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CustomerApiCustomerPrototypePutDeviceFreezeOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **deleteAllExceptSuspend** | **optional.Bool**|  | 
 **schedules** | **optional.String**|  | 
 **enable** | **optional.Bool**|  | 
 **name** | **optional.String**|  | 

### Return type

[**XAny**](x-any.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomerPrototypePutDeviceFreezeAutoExpire**
> XAny CustomerPrototypePutDeviceFreezeAutoExpire(ctx, id, locationId, mac, expiresAt)
Put a device to be frozen for a Location ID.

<div><strong>200</strong>: Success, updated.</div> <div><strong>401</strong>: Authorization required or customer id not found.</div> <div><strong>404</strong>: Location id, does not exist.</div> <div><strong>500</strong>: Internal server error</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Customer id | 
  **locationId** | **string**|  | 
  **mac** | **string**|  | 
  **expiresAt** | **string**|  | 

### Return type

[**XAny**](x-any.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomerPrototypePutDeviceFreezeForever**
> XAny CustomerPrototypePutDeviceFreezeForever(ctx, id, locationId, mac, optional)
Put a device forever freeze for a Location ID.

<div><strong>200</strong>: Success, updated.</div> <div><strong>401</strong>: Authorization required or customer id not found.</div> <div><strong>404</strong>: Location id, does not exist.</div> <div><strong>500</strong>: Internal server error</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Customer id | 
  **locationId** | **string**|  | 
  **mac** | **string**|  | 
 **optional** | ***CustomerApiCustomerPrototypePutDeviceFreezeForeverOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CustomerApiCustomerPrototypePutDeviceFreezeForeverOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **deleteAllExceptSuspend** | **optional.Bool**|  | 
 **enable** | **optional.Bool**|  | 

### Return type

[**XAny**](x-any.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomerPrototypePutDeviceFreezeResidentialGwManaged**
> XAny CustomerPrototypePutDeviceFreezeResidentialGwManaged(ctx, id, locationId, mac, optional)
Put a device residentialGwManaged freeze for a Location ID.

<div><strong>200</strong>: Success, updated.</div> <div><strong>401</strong>: Authorization required or customer id not found.</div> <div><strong>404</strong>: Location id, does not exist.</div> <div><strong>500</strong>: Internal server error</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Customer id | 
  **locationId** | **string**|  | 
  **mac** | **string**|  | 
 **optional** | ***CustomerApiCustomerPrototypePutDeviceFreezeResidentialGwManagedOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CustomerApiCustomerPrototypePutDeviceFreezeResidentialGwManagedOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **deleteAllExceptSuspend** | **optional.Bool**|  | 
 **schedules** | **optional.String**|  | 
 **enable** | **optional.Bool**|  | 
 **name** | **optional.String**|  | 

### Return type

[**XAny**](x-any.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomerPrototypePutDeviceFreezeSuspend**
> XAny CustomerPrototypePutDeviceFreezeSuspend(ctx, id, locationId, mac, optional)
Put a device suspend for a Location ID.

<div><strong>200</strong>: Success, updated.</div> <div><strong>401</strong>: Authorization required or customer id not found.</div> <div><strong>404</strong>: Location id, does not exist.</div> <div><strong>500</strong>: Internal server error</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Customer id | 
  **locationId** | **string**|  | 
  **mac** | **string**|  | 
 **optional** | ***CustomerApiCustomerPrototypePutDeviceFreezeSuspendOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CustomerApiCustomerPrototypePutDeviceFreezeSuspendOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **deleteAllExceptSuspend** | **optional.Bool**|  | 
 **enable** | **optional.Bool**|  | 

### Return type

[**XAny**](x-any.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomerPrototypePutDeviceNickname**
> Customer CustomerPrototypePutDeviceNickname(ctx, id, mac, optional)
Nickname a Customer's device for all locations.

<div><strong>200</strong>: Success, device name has been updated<br/>but not validated as a device that <br/>has ever connected.</div> <div><strong>400</strong>: nickname value must be defined.</div> <div><strong>404</strong>: customer id and/or mac does not exist.</div> <div><strong>422</strong>: nickname value must be less than 33 characters.</div> <div><strong>500</strong>: internal server error</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Customer id | 
  **mac** | **string**|  | 
 **optional** | ***CustomerApiCustomerPrototypePutDeviceNicknameOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CustomerApiCustomerPrototypePutDeviceNicknameOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **nickname** | **optional.String**|  | 

### Return type

[**Customer**](Customer.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomerPrototypePutDevicesFreezeCommand**
> []interface{} CustomerPrototypePutDevicesFreezeCommand(ctx, id, locationId, devices)
Sets autoExpire or freezeTemplateId suspend for specified mac addresses

<div><strong>200</strong>: Succcess, updated</div> <div><strong>401</strong>: Unathorized</div> <div><strong>404</strong>: Mac not found</div> <div><strong>422</strong>: Invalid freezeTemplateID</div> <div><strong>500</strong>: Internal Server Error</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Customer id | 
  **locationId** | **string**|  | 
  **devices** | [**[]XAny**](x-any.md)|  | 

### Return type

**[]interface{}**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomerPrototypePutDevicesVisibleToGuests**
> []string CustomerPrototypePutDevicesVisibleToGuests(ctx, id, locationId, devicesVisibleToGuests)
DEPRECATED: Update home devices visible to guests.

<div><strong>200</strong>: Success, devicesVisibleToGuests returned.</div> <div><strong>400</strong>: Required fields missing.</div> <div><strong>401</strong>: Authorization required or customer id not found.</div> <div><strong>404</strong>: Customer id, location id, or WifiNetwork does not exist.</div> <div><strong>422</strong>: Device mac validation failed.</div> <div><strong>500</strong>: Internal server error</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Customer id | 
  **locationId** | **string**|  | 
  **devicesVisibleToGuests** | **string**| array of macs[] | 

### Return type

**[]string**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomerPrototypePutDhcp**
> interface{} CustomerPrototypePutDhcp(ctx, id, locationId, subnet, optional)
Record or update a new DHCP subnet/subnetMask for a Location ID.

<div><strong>200</strong>: Success, DHCP are returned.</div> <div><strong>401</strong>: Authorization required or customer id not found</div> <div><strong>422</strong>: subnet value is empty, or invalid.</div> <div><strong>500</strong>: internal server error</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Customer id | 
  **locationId** | **string**|  | 
  **subnet** | **string**|  | 
 **optional** | ***CustomerApiCustomerPrototypePutDhcpOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CustomerApiCustomerPrototypePutDhcpOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **subnetMask** | **optional.String**|  | 
 **startIp** | **optional.String**|  | 
 **endIp** | **optional.String**|  | 

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomerPrototypePutDhcpReservation**
> DhcpReservation CustomerPrototypePutDhcpReservation(ctx, id, locationId, mac, ip, optional)
Record or update a new DHCP IP Reservation for a particular MAC address at a Location ID.

<div><strong>200</strong>: Success, all DHCP Reservations are returned.</div> <div><strong>401</strong>: Authorization required or customer id not found</div> <div><strong>412</strong>: Subnet prefix is unknown.</div> <div><strong>422</strong>: IP/mac value is empty, or invalid.</div> <div><strong>500</strong>: internal server error</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Customer id | 
  **locationId** | **string**|  | 
  **mac** | **string**|  | 
  **ip** | **string**|  | 
 **optional** | ***CustomerApiCustomerPrototypePutDhcpReservationOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CustomerApiCustomerPrototypePutDhcpReservationOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **hostName** | **optional.String**|  | 

### Return type

[**DhcpReservation**](DhcpReservation.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomerPrototypePutDnsServers**
> DnsServers CustomerPrototypePutDnsServers(ctx, id, locationId, optional)
Update the DNS IPv4 server addresses for a Location ID.

<div><strong>200</strong>: Success, new DNS Servers saved.</div> <div><strong>401</strong>: Authorization required or customer id not found</div> <div><strong>422</strong>: primaryDns or secondaryDns DNS Servers value is empty.</div> <div><strong>500</strong>: internal server error</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Customer id | 
  **locationId** | **string**|  | 
 **optional** | ***CustomerApiCustomerPrototypePutDnsServersOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CustomerApiCustomerPrototypePutDnsServersOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **primaryDns** | **optional.String**|  | 
 **secondaryDns** | **optional.String**|  | 

### Return type

[**DnsServers**](DnsServers.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomerPrototypePutDppEnrollments**
> interface{} CustomerPrototypePutDppEnrollments(ctx, id, locationId, enrollments)
Create and persist a list of DPP enrollments

<div><strong>202</strong>: Success.</div> <div><strong>400</strong>: Required fields missing or field type is incorrect.</div> <div><strong>401</strong>: Authorization required or customer id not found</div> <div><strong>404</strong>: Location id or wifi network does not exist.</div> <div><strong>500</strong>: internal server error</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Customer id | 
  **locationId** | **string**|  | 
  **enrollments** | [**[]XAny**](x-any.md)|  | 

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomerPrototypePutEthernetLan**
> CustomerPrototypePutEthernetLan(ctx, id, locationId, nodeId, nodeEthernetLan)
Updates location nodes with ethernetLan modes

<div><strong>202</strong>: Success.</div> <div><strong>404</strong>: Location does not exist.</div> <div><strong>404</strong>: Node does not exist.</div> <div><strong>422</strong>: nodeEthernetLan does not exist.</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Customer id | 
  **locationId** | **string**|  | 
  **nodeId** | **string**|  | 
  **nodeEthernetLan** | **interface{}**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomerPrototypePutFirmwareUpgradeRequest**
> []interface{} CustomerPrototypePutFirmwareUpgradeRequest(ctx, id, locationId)
Request Firmware Upgrade for a Location ID

<div><strong>200</strong>: Success.</div> <div><strong>401</strong>: Authorization required or customer id not found</div> <div><strong>404</strong>: customer id or location id does not exist and is not known to Plume</div> <div><strong>500</strong>: internal server error</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Customer id | 
  **locationId** | **string**|  | 

### Return type

**[]interface{}**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomerPrototypePutFrontlineStorage**
> CustomerPrototypePutFrontlineStorage(ctx, id, locationId, data)
Create or Update the frontline storage for a Location ID

<div><strong>204</strong>: Success.</div> <div><strong>401</strong>: Authorization required or customer id not found</div> <div><strong>404</strong>: customer id or location id does not exist and is not known to Plume</div> <div><strong>500</strong>: internal server error</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Customer id | 
  **locationId** | **string**|  | 
  **data** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomerPrototypePutGroupOfUnassignedDevicesFreezeAutoExpire**
> XAny CustomerPrototypePutGroupOfUnassignedDevicesFreezeAutoExpire(ctx, id, locationId, expiresAt)
Put GroupOfUnassignedDevices autoExpire freeze for a Location ID.

<div><strong>200</strong>: Success, updated.</div> <div><strong>401</strong>: Authorization required or customer id not found.</div> <div><strong>404</strong>: Location id, does not exist.</div> <div><strong>500</strong>: Internal server error</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Customer id | 
  **locationId** | **string**|  | 
  **expiresAt** | **string**|  | 

### Return type

[**XAny**](x-any.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomerPrototypePutGroupOfUnassignedDevicesFreezeForever**
> XAny CustomerPrototypePutGroupOfUnassignedDevicesFreezeForever(ctx, id, locationId)
PUT GroupOfUnassignedDevices forever for a Location ID.

<div><strong>200</strong>: Success, updated.</div> <div><strong>401</strong>: Authorization required or customer id not found.</div> <div><strong>404</strong>: Location id, does not exist.</div> <div><strong>500</strong>: Internal server error</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Customer id | 
  **locationId** | **string**|  | 

### Return type

[**XAny**](x-any.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomerPrototypePutGroupOfUnassignedDevicesFreezeSuspend**
> XAny CustomerPrototypePutGroupOfUnassignedDevicesFreezeSuspend(ctx, id, locationId)
PUT GroupOfUnassignedDevices suspend for a Location ID.

<div><strong>200</strong>: Success, updated.</div> <div><strong>401</strong>: Authorization required or customer id not found.</div> <div><strong>404</strong>: Location id, does not exist.</div> <div><strong>500</strong>: Internal server error</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Customer id | 
  **locationId** | **string**|  | 

### Return type

[**XAny**](x-any.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomerPrototypePutIosDeviceToken**
> CustomerPrototypePutIosDeviceToken(ctx, id, deviceToken)
Inserts the iOS device token for the Customer ID, which may be used for notification services.

<div><strong>204</strong>: Success, most recent IOS device Token saved.</div> <div><strong>401</strong>: Authorization required or customer id not found</div> <div><strong>422</strong>: deviceToken value must be defined.</div> <div><strong>500</strong>: internal server error</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Customer id | 
  **deviceToken** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomerPrototypePutIspSpeedTestConfiguration**
> InlineResponse2004 CustomerPrototypePutIspSpeedTestConfiguration(ctx, id, locationId, enable, optional)
Enable|Disable ispSpeedTestConfiguration to schedule speed tests.

<div><strong>200</strong>: Success, run.</div> <div><strong>401</strong>: Authorization required.</div> <div><strong>404</strong>: Customer or location does not exists.</div> <div><strong>500</strong>: Internal server error</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Customer id | 
  **locationId** | **string**|  | 
  **enable** | **string**| boolean but marked as &#39;any&#39; because our mobile app platforms mixed string and boolean primitive | 
 **optional** | ***CustomerApiCustomerPrototypePutIspSpeedTestConfigurationOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CustomerApiCustomerPrototypePutIspSpeedTestConfigurationOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **enableAllNodes** | **optional.String**| boolean but treated as a string since it is optional | 

### Return type

[**InlineResponse2004**](inline_response_200_4.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomerPrototypePutLedMode**
> InlineResponse2003 CustomerPrototypePutLedMode(ctx, id, locationId, nodeId, optional)
Update the LED mode on a particular Node for a Location ID.

When the mode is set to \"locate\", the Node with that ID at this locationId, will have its LED blinked for locating purposes. The mode is set to \"normal\" to return the LED to its normal mode of operation. <div><strong>200</strong>: Success, your new info looks good.</div> <div><strong>401</strong>: Authorization required or customer id not found</div> <div><strong>404</strong>: location id does not exist.</div> <div><strong>422</strong>: ledMode must be defined.</div> <div><strong>422</strong>: ledMode must be \"locate\" or \"normal\".</div> <div><strong>422</strong>: nodeId must be defined.</div> <div><strong>425</strong>: nodeId must belong to the location.</div> <div><strong>500</strong>: internal server error</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Customer id | 
  **locationId** | **string**|  | 
  **nodeId** | **string**|  | 
 **optional** | ***CustomerApiCustomerPrototypePutLedModeOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CustomerApiCustomerPrototypePutLedModeOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **mode** | **optional.String**| locate/normal | 

### Return type

[**InlineResponse2003**](inline_response_200_3.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomerPrototypePutLocale**
> Locale CustomerPrototypePutLocale(ctx, id, locationId, region)
Configure locale values for a Location ID.

<div><strong>200</strong>: Success, updated.</div> <div><strong>401</strong>: Authorization required or customer id not found.</div> <div><strong>404</strong>: Location id, does not exist.</div> <div><strong>422</strong>: Region value is not valid.</div> <div><strong>500</strong>: Internal server error</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Customer id | 
  **locationId** | **string**|  | 
  **region** | **string**| during optimizations, used to determine allowed WiFi channels. Possible values: US, SINGAPORE, UK, EU, CANADA, JP. | 

### Return type

[**Locale**](Locale.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomerPrototypePutLocationSecurityPolicyOHPDeviceUuidMapping**
> interface{} CustomerPrototypePutLocationSecurityPolicyOHPDeviceUuidMapping(ctx, id, locationId, uuid, optional)
Update the Device UUID Mapping for Out of Home Protection.

<div><strong>200</strong>: Success.</div> <div><strong>400</strong>: Required fields missing or field type is incorrect.</div> <div><strong>401</strong>: Authorization required or customer id not found</div> <div><strong>404</strong>: Location id or WifiNetwork does not exist and is not known to Plume</div> <div><strong>500</strong>: Internal server error</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Customer id | 
  **locationId** | **string**|  | 
  **uuid** | **string**|  | 
 **optional** | ***CustomerApiCustomerPrototypePutLocationSecurityPolicyOHPDeviceUuidMappingOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CustomerApiCustomerPrototypePutLocationSecurityPolicyOHPDeviceUuidMappingOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **lanIpv4** | **optional.String**| Mobile device lanIpv4 address, if any | 
 **lanIpv6** | **optional.String**| Mobile device lanIpv6 address, if any | 

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomerPrototypePutLocationSecurityPolicyOHPProtectionState**
> interface{} CustomerPrototypePutLocationSecurityPolicyOHPProtectionState(ctx, id, locationId, uuid, optional)
Update the Device Protection State for Out of Home Protection.

<div><strong>200</strong>: Success.</div> <div><strong>400</strong>: Required fields missing or field type is incorrect.</div> <div><strong>401</strong>: Authorization required or customer id not found</div> <div><strong>404</strong>: Location id or WifiNetwork does not exist and is not known to Plume</div> <div><strong>500</strong>: Internal server error</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Customer id | 
  **locationId** | **string**|  | 
  **uuid** | **string**| Mobile device uuid (as was assigned by Mobolize) | 
 **optional** | ***CustomerApiCustomerPrototypePutLocationSecurityPolicyOHPProtectionStateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CustomerApiCustomerPrototypePutLocationSecurityPolicyOHPProtectionStateOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **protectionState** | **optional.String**| ProtectionState info as obtained directly from the Mobolize SDK, null if deleting ProtectionState | 

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomerPrototypePutLocationWanConfiguration**
> interface{} CustomerPrototypePutLocationWanConfiguration(ctx, id, locationId, optional)
Persist WAN Configuration for a Location ID.

<div><strong>200</strong>: Success, updated.</div> <div><strong>400</strong>: Required fields missing.</div> <div><strong>401</strong>: Authorization required or customer id not found.</div> <div><strong>404</strong>: Location id, does not exist.</div> <div><strong>422</strong>: Required fields are not valid.</div> <div><strong>500</strong>: Internal server error</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Customer id | 
  **locationId** | **string**|  | 
 **optional** | ***CustomerApiCustomerPrototypePutLocationWanConfigurationOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CustomerApiCustomerPrototypePutLocationWanConfigurationOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **pppoe** | **optional.String**|  | 
 **uplink** | **optional.String**|  | 
 **staticIPv4** | **optional.String**|  | 
 **publishedWithBLE** | **optional.Bool**|  | 

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomerPrototypePutLocationWanSettings**
> LocationWanSettings CustomerPrototypePutLocationWanSettings(ctx, id, locationId, wanSettings)
DEPRECATED: Persist WAN Settings for a Location ID.

<div><strong>200</strong>: Success, updated.</div> <div><strong>400</strong>: Required fields missing.</div> <div><strong>401</strong>: Authorization required or customer id not found.</div> <div><strong>404</strong>: Location id, does not exist.</div> <div><strong>422</strong>: Required fields are not valid.</div> <div><strong>500</strong>: Internal server error</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Customer id | 
  **locationId** | **string**|  | 
  **wanSettings** | [**LocationWanSettings**](LocationWanSettings.md)|  | 

### Return type

[**LocationWanSettings**](LocationWanSettings.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomerPrototypePutMonitorMode**
> LocationMonitorMode CustomerPrototypePutMonitorMode(ctx, id, locationId, enable)
Enable/disable monitor mode for a Location ID.

<div><strong>200</strong>: Success, updated.</div> <div><strong>400</strong>: Required fields missing.</div> <div><strong>401</strong>: Authorization required or customer id not found.</div> <div><strong>404</strong>: Location id, does not exist.</div> <div><strong>500</strong>: Internal server error</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Customer id | 
  **locationId** | **string**|  | 
  **enable** | **string**|  | 

### Return type

[**LocationMonitorMode**](LocationMonitorMode.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomerPrototypePutNetworkMode**
> InlineResponse2006 CustomerPrototypePutNetworkMode(ctx, id, locationId, optional)
Update the Network Mode for a Location ID.

<div><strong>200</strong>: Success, your new info looks good.</div> <div><strong>401</strong>: Authorization required or customer id not found</div> <div><strong>404</strong>: location id, does not exist.</div> <div><strong>500</strong>: internal server error</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Customer id | 
  **locationId** | **string**|  | 
 **optional** | ***CustomerApiCustomerPrototypePutNetworkModeOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CustomerApiCustomerPrototypePutNetworkModeOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **networkMode** | **optional.String**|  | 

### Return type

[**InlineResponse2006**](inline_response_200_6.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomerPrototypePutOptimizations**
> Optimizations CustomerPrototypePutOptimizations(ctx, id, locationId, auto, optional)
Enable/disable optimizations for a Location ID.

<div><strong>200</strong>: Success, updated.</div> <div><strong>401</strong>: Authorization required or customer id not found.</div> <div><strong>404</strong>: Location id, does not exist.</div> <div><strong>422</strong>: Invalid dfsMode, prefer160MhzMode, hopPenalty or preCACScheduler provided.</div> <div><strong>500</strong>: Internal server error</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Customer id | 
  **locationId** | **string**|  | 
  **auto** | **string**| defaults to true | 
 **optional** | ***CustomerApiCustomerPrototypePutOptimizationsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CustomerApiCustomerPrototypePutOptimizationsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **dfsMode** | **optional.String**| enum of values include: auto, enable, disable, demo, HomeNonDFSChannels, usDfs, deviceAware | [default to auto]
 **prefer160MhzMode** | **optional.String**| enum of values include: auto, enable, disable | [default to auto]
 **hopPenalty** | **optional.String**| enum of values include: auto, low, medium, high | [default to auto]
 **preCACScheduler** | **optional.String**| enum of values include: auto, enable, disable | [default to auto]

### Return type

[**Optimizations**](Optimizations.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomerPrototypePutOverlordResync**
> CustomerPrototypePutOverlordResync(ctx, id, locationId)
Push Secondary Network Configurations to Overlord.

<div><strong>204</strong>: Success.</div> <div><strong>401</strong>: Authorization required or customer id not found.</div> <div><strong>404</strong>: Location id, does not exist.</div> <div><strong>500</strong>: Internal server error</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Customer id | 
  **locationId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomerPrototypePutPersistConfigurationOnGateway**
> EthernetLan CustomerPrototypePutPersistConfigurationOnGateway(ctx, id, locationId, persistConfigurationOnGateway)
Update settings for persistConfigurationOnGateway.

Supported modes are: * enable/disable/auto  <div><strong>200</strong>: Success, new ethernetLan settings saved.</div> <div><strong>401</strong>: Authorization required or customer id not found</div> <div><strong>422</strong>: Input validation error, see output for details.</div> <div><strong>500</strong>: Internal server error</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Customer id | 
  **locationId** | **string**|  | 
  **persistConfigurationOnGateway** | [**EthernetLan**](EthernetLan.md)| ethernetLan object | 

### Return type

[**EthernetLan**](EthernetLan.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomerPrototypePutPersonFreeze**
> XAny CustomerPrototypePutPersonFreeze(ctx, id, locationId, personId, freezeTemplateId, optional)
Put a person to be frozen for a Location ID.

<div><strong>200</strong>: Success, updated.</div> <div><strong>401</strong>: Authorization required or customer id not found.</div> <div><strong>404</strong>: Location id, does not exist.</div> <div><strong>500</strong>: Internal server error</div> <div><strong>501</strong>: Not Implemented, if location is utilizing shared location freeze schedules</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Customer id | 
  **locationId** | **string**|  | 
  **personId** | **string**|  | 
  **freezeTemplateId** | **string**| Valid templates are &#39;untilMidinight&#39;, &#39;schoolNights&#39;, etc. | 
 **optional** | ***CustomerApiCustomerPrototypePutPersonFreezeOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CustomerApiCustomerPrototypePutPersonFreezeOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **deleteAllExceptSuspend** | **optional.Bool**|  | 
 **schedules** | **optional.String**|  | 
 **enable** | **optional.Bool**|  | 
 **name** | **optional.String**|  | 

### Return type

[**XAny**](x-any.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomerPrototypePutPersonFreezeAutoExpire**
> XAny CustomerPrototypePutPersonFreezeAutoExpire(ctx, id, locationId, personId, expiresAt)
Put all devices from a person to be frozen for a Location ID.

<div><strong>200</strong>: Success, updated.</div> <div><strong>401</strong>: Authorization required or customer id not found.</div> <div><strong>404</strong>: Location id, does not exist.</div> <div><strong>500</strong>: Internal server error</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Customer id | 
  **locationId** | **string**|  | 
  **personId** | **string**|  | 
  **expiresAt** | **string**|  | 

### Return type

[**XAny**](x-any.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomerPrototypePutPersonFreezeForever**
> XAny CustomerPrototypePutPersonFreezeForever(ctx, id, locationId, personId, optional)
Put a person forever freeze for a Location ID.

<div><strong>200</strong>: Success, updated.</div> <div><strong>401</strong>: Authorization required or customer id not found.</div> <div><strong>404</strong>: Location id, does not exist.</div> <div><strong>500</strong>: Internal server error</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Customer id | 
  **locationId** | **string**|  | 
  **personId** | **string**|  | 
 **optional** | ***CustomerApiCustomerPrototypePutPersonFreezeForeverOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CustomerApiCustomerPrototypePutPersonFreezeForeverOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **deleteAllExceptSuspend** | **optional.Bool**|  | 
 **enable** | **optional.Bool**|  | 

### Return type

[**XAny**](x-any.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomerPrototypePutPersonFreezeSuspend**
> XAny CustomerPrototypePutPersonFreezeSuspend(ctx, id, locationId, personId, optional)
Put a person suspend for a Location ID.

<div><strong>200</strong>: Success, updated.</div> <div><strong>401</strong>: Authorization required or customer id not found.</div> <div><strong>404</strong>: Location id, does not exist.</div> <div><strong>500</strong>: Internal server error</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Customer id | 
  **locationId** | **string**|  | 
  **personId** | **string**|  | 
 **optional** | ***CustomerApiCustomerPrototypePutPersonFreezeSuspendOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CustomerApiCustomerPrototypePutPersonFreezeSuspendOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **deleteAllExceptSuspend** | **optional.Bool**|  | 
 **enable** | **optional.Bool**|  | 

### Return type

[**XAny**](x-any.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomerPrototypePutPortForward**
> PortForward CustomerPrototypePutPortForward(ctx, id, locationId, mac, externalPort, optional)
Update an existing Port Forwarding entry for an existing DHCP IP reservation tied to a MAC address at a Location ID.

<div><strong>200</strong>: Success, all PortForwards are returned.</div> <div><strong>401</strong>: Authorization required or customer id not found</div> <div><strong>422</strong>: networkConfiguration, dhcpReservation, PortForward is empty.</div> <div><strong>422</strong>: mac is empty, or invalid, externalPort/internalPort is out of range, or protocol is invalid.</div> <div><strong>500</strong>: internal server error</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Customer id | 
  **locationId** | **string**|  | 
  **mac** | **string**|  | 
  **externalPort** | **string**|  | 
 **optional** | ***CustomerApiCustomerPrototypePutPortForwardOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CustomerApiCustomerPrototypePutPortForwardOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **internalPort** | **optional.String**|  | 
 **protocol** | **optional.String**|  | 
 **name** | **optional.String**|  | 

### Return type

[**PortForward**](PortForward.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomerPrototypePutSniffing**
> CustomerPrototypePutSniffing(ctx, id, locationId, dns, http, upnp, mdns)
Updates location sniffing toggle modes

<div><strong>202</strong>: Success.</div> <div><strong>400</strong>: Required fields missing or field type is incorrect.</div> <div><strong>404</strong>: Location does not exist.</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Customer id | 
  **locationId** | **string**|  | 
  **dns** | **string**| object with property \&quot;mode\&quot;: an enum of values which include: auto, enable, disable | 
  **http** | **string**| object with property \&quot;mode\&quot;: an enum of values which include: auto, enable, disable | 
  **upnp** | **string**| object with property \&quot;mode\&quot;: an enum of values which include: auto, enable, disable | 
  **mdns** | **string**| object with property \&quot;mode\&quot;: an enum of values which include: auto, enable, disable | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomerPrototypePutSnoozeOnDeviceAlert**
> Device CustomerPrototypePutSnoozeOnDeviceAlert(ctx, id, locationId, mac, type_, optional)
Snooze an alert on a device.

<div><strong>200</strong>: Success, updated.</div> <div><strong>401</strong>: Authorization required or customer id not found.</div> <div><strong>404</strong>: Location id does not exist or device mac not in this account's recent history.</div> <div><strong>422</strong>: Invalid alert type and/or state.</div> <div><strong>500</strong>: Internal server error</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Customer id | 
  **locationId** | **string**|  | 
  **mac** | **string**| mac of device | 
  **type_** | **string**| enum of values include: poorHealth | 
 **optional** | ***CustomerApiCustomerPrototypePutSnoozeOnDeviceAlertOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CustomerApiCustomerPrototypePutSnoozeOnDeviceAlertOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **state** | **optional.String**| enum of values include: snooze, ignore, performanceAcceptable | 

### Return type

[**Device**](Device.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomerPrototypePutSnoozeOnNodeAlert**
> Node CustomerPrototypePutSnoozeOnNodeAlert(ctx, id, locationId, nodeId, type_, optional)
Snooze an alert on a node.

<div><strong>200</strong>: Success, updated.</div> <div><strong>401</strong>: Authorization required or customer id not found.</div> <div><strong>404</strong>: Location id does not exist or nodeId not claimed to this account.</div> <div><strong>422</strong>: Invalid alert type and/or state.</div> <div><strong>500</strong>: Internal server error</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Customer id | 
  **locationId** | **string**|  | 
  **nodeId** | **string**| id of node | 
  **type_** | **string**| enum of values include: poorHealth | 
 **optional** | ***CustomerApiCustomerPrototypePutSnoozeOnNodeAlertOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CustomerApiCustomerPrototypePutSnoozeOnNodeAlertOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **state** | **optional.String**| enum of values include: snooze, ignore, performanceAcceptable, reset | 

### Return type

[**Node**](Node.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomerPrototypePutSubscription**
> interface{} CustomerPrototypePutSubscription(ctx, id, locationId, ratePlanId)
Put Subscription details for this location

<div><strong>200</strong>: Success, service level returned</div> <div><strong>401</strong>: Authorization required or customer id not found</div> <div><strong>404</strong>: customer id or location id does not exist and is not known to Plume</div> <div><strong>500</strong>: internal server error</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Customer id | 
  **locationId** | **string**|  | 
  **ratePlanId** | **string**|  | 

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomerPrototypePutUpnp**
> Upnp CustomerPrototypePutUpnp(ctx, id, locationId, optional)
Update the UPnP setting for a Location ID.

<div><strong>200</strong>: Success, new Upnp saved.</div> <div><strong>401</strong>: Authorization required or customer id not found</div> <div><strong>422</strong>: Upnp value is empty.</div> <div><strong>500</strong>: internal server error</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Customer id | 
  **locationId** | **string**|  | 
 **optional** | ***CustomerApiCustomerPrototypePutUpnpOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CustomerApiCustomerPrototypePutUpnpOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **enabled** | **optional.String**| DEPRECATED: boolean but marked as &#39;any&#39; because our mobile app platforms mixed string and boolean primitive | 
 **mode** | **optional.String**| Possible values enable/disable/auto | 

### Return type

[**Upnp**](Upnp.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomerPrototypePutWifiKey**
> []WifiNetworkKey CustomerPrototypePutWifiKey(ctx, id, locationId, accessZone, keyId, encryptionKey, enable, format, optional)
Update a WiFi Password

<div><strong>200</strong>: Success, all passwords returned</div> <div><strong>400</strong>: Required fields missing</div> <div><strong>401</strong>: Authorization required or customer id not found</div> <div><strong>404</strong>: Customer id, location id, or WifiNetwork does not exist</div> <div><strong>405</strong>: Cannot disable a read-only key</div> <div><strong>422</strong>: Password validation failed</div> <div><strong>500</strong>: Internal server error</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Customer id | 
  **locationId** | **string**|  | 
  **accessZone** | **string**| home | guests | internetAccessOnly | 
  **keyId** | **float64**| Unique password id: 0-9 | 
  **encryptionKey** | **string**|  | 
  **enable** | **bool**| devices can connect using this encryptionKey | 
  **format** | **string**| encryptionKey | phoneNumber | 
 **optional** | ***CustomerApiCustomerPrototypePutWifiKeyOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CustomerApiCustomerPrototypePutWifiKeyOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------







 **expiresAt** | **optional.Time**| UTC in ISO 8601 String format | 
 **content** | **optional.String**| Valid values: &#39;adultAndSensitive&#39; | 

### Return type

[**[]WifiNetworkKey**](WifiNetworkKey.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomerPrototypePutWifiNetwork**
> InlineResponse2002 CustomerPrototypePutWifiNetwork(ctx, id, locationId, optional)
Update a WiFi SSID and/or PSK for a Location ID.

<div><strong>200</strong>: Success, your new info looks good.</div> <div><strong>401</strong>: Authorization required or customer id not found</div> <div><strong>404</strong>: location id, or wifi network does not exist.</div> <div><strong>422</strong>: encryptionKey or ssid must be defined.</div> <div><strong>500</strong>: internal server error</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Customer id | 
  **locationId** | **string**|  | 
 **optional** | ***CustomerApiCustomerPrototypePutWifiNetworkOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CustomerApiCustomerPrototypePutWifiNetworkOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **encryptionKey** | **optional.String**|  | 
 **ssid** | **optional.String**|  | 
 **wpaMode** | **optional.String**|  | 

### Return type

[**InlineResponse2002**](inline_response_200_2.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomerPrototypeRebootLocation**
> CustomerPrototypeRebootLocation(ctx, id, locationId, optional)
Reboots a particular on-line Node for a Location ID.

<div><strong>204</strong>: Success, your new info looks good.</div> <div><strong>401</strong>: Authorization required or customer id not found.</div> <div><strong>404</strong>: location id, does not exist.</div> <div><strong>422</strong>: Delay, is not between 0 and 100000.</div> <div><strong>500</strong>: internal server error</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Customer id | 
  **locationId** | **string**|  | 
 **optional** | ***CustomerApiCustomerPrototypeRebootLocationOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CustomerApiCustomerPrototypeRebootLocationOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **delay** | **optional.String**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomerPrototypeRebootNode**
> CustomerPrototypeRebootNode(ctx, id, locationId, nodeId, optional)
Reboots a single on-line Node for a Location ID.

<div><strong>204</strong>: Success.</div> <div><strong>401</strong>: Authorization required or customer id not found.</div> <div><strong>404</strong>: location id or nodeId, does not exist.</div> <div><strong>422</strong>: Delay, is not between 0 and 100000.</div> <div><strong>500</strong>: internal server error</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Customer id | 
  **locationId** | **string**|  | 
  **nodeId** | **string**|  | 
 **optional** | ***CustomerApiCustomerPrototypeRebootNodeOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CustomerApiCustomerPrototypeRebootNodeOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **delay** | **optional.String**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomerPrototypeRejectWhitelistRequest**
> interface{} CustomerPrototypeRejectWhitelistRequest(ctx, id, locationId, requestId)
Reject an approval request for a website whitelist

<div><strong>204</strong>: Success.</div> <div><strong>400</strong>: Required fields missing or field type is incorrect.</div> <div><strong>401</strong>: Authorization required or customer id not found</div> <div><strong>404</strong>: Location id, CustomerId or requst id does not exist and is not known to Plume</div> <div><strong>500</strong>: Internal server error</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Customer id | 
  **locationId** | **string**|  | 
  **requestId** | **string**|  | 

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomerPrototypeRemoveDeviceByMac**
> CustomerPrototypeRemoveDeviceByMac(ctx, id, locationId, mac, optional)
Removes a device for a customer's location id, wiping config and setting a hidden flag.

<div><strong>204</strong>: Success, device removed from location. </div> <div><strong>404</strong>: location id or  device not found. </div> <div><strong>500</strong>: internal server error </div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Customer id | 
  **locationId** | **string**|  | 
  **mac** | **string**| mac of device | 
 **optional** | ***CustomerApiCustomerPrototypeRemoveDeviceByMacOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CustomerApiCustomerPrototypeRemoveDeviceByMacOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **daysOffline** | **optional.Float64**| exclude devices disconnected longer than daysOffline. if not set, it will be 31. for older devices, it will return 404, \&quot;not found\&quot; | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomerPrototypeRemoveDeviceVisibleToGuests**
> []string CustomerPrototypeRemoveDeviceVisibleToGuests(ctx, id, locationId, mac)
DEPRECATED: Update home devices visible to guests.

<div><strong>200</strong>: Success, devicesVisibleToGuests returned.</div> <div><strong>400</strong>: Required fields missing.</div> <div><strong>401</strong>: Authorization required or customer id not found.</div> <div><strong>404</strong>: Customer id, location id, or WifiNetwork does not exist.</div> <div><strong>422</strong>: Device mac validation failed.</div> <div><strong>500</strong>: Internal server error</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Customer id | 
  **locationId** | **string**|  | 
  **mac** | **string**| mac to be removed | 

### Return type

**[]string**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomerPrototypeRenameNode**
> Node CustomerPrototypeRenameNode(ctx, id, locationId, nodeId, nickname, optional)
Rename a particular Node for a Location ID with the option to disable the blinking LED.

Rename a particular Node for a Location ID with the option to disable the blinking LED with the option \"emitMessage\":\"on\" or \"off\". <div><strong>200</strong>: Success, a job well done.</div> <div><strong>400</strong>: Bad request, nickname is undefined or empty string.</div> <div><strong>401</strong>: Authorization required, customer id not found, <br/> or id not owned by requestor.</div> <div><strong>404</strong>: location ID or node ID not found.</div> <div><strong>500</strong>: internal server error</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Customer id | 
  **locationId** | **string**|  | 
  **nodeId** | **string**|  | 
  **nickname** | **string**|  | 
 **optional** | ***CustomerApiCustomerPrototypeRenameNodeOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CustomerApiCustomerPrototypeRenameNodeOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **emitMessage** | **optional.String**|  | 

### Return type

[**Node**](Node.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomerPrototypeResendManagerInvite**
> interface{} CustomerPrototypeResendManagerInvite(ctx, id, locationId, managerId, optional)
Resend invite to a manager that has status \"pending\".

<div><strong>204</strong>: Success.</div> <div><strong>400</strong>: Required fields missing or field type is incorrect.</div> <div><strong>404</strong>: Location or Manager does not exist.</div> <div><strong>422</strong>: Manager already accepted the invite to manage the location </div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Customer id | 
  **locationId** | **string**|  | 
  **managerId** | **string**|  | 
 **optional** | ***CustomerApiCustomerPrototypeResendManagerInviteOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CustomerApiCustomerPrototypeResendManagerInviteOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **notificationOptions** | **optional.String**|  | 

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomerPrototypeResetTos**
> XAny CustomerPrototypeResetTos(ctx, id, locationId, mac)
Resets the back-off and thresholds for the given client.

<div><strong>200</strong>: Ok.</div> <div><strong>401</strong>: Authorization required.</div> <div><strong>404</strong>: Location id, does not exist.</div> <div><strong>404</strong>: No device found with provided mac address</div> <div><strong>422</strong>: Invalid MAC.</div> <div><strong>500</strong>: Internal server error.</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Customer id | 
  **locationId** | **string**|  | 
  **mac** | **string**|  | 

### Return type

[**XAny**](x-any.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomerPrototypeSetForcedSteer**
> interface{} CustomerPrototypeSetForcedSteer(ctx, id, locationId, mac, optional)
Force a device to use the 2.4Ghz band with auto expire.

<div><strong>204</strong>: Success, forced steer enabled.</div> <div><strong>404</strong>: Location ID or Device mac not found or the device has not been online in the last 31 days</div> <div><strong>422</strong>: expiresAt is outside of the expected range 5 to 60 minutes in the future</div> <div><strong>422</strong>: expiresAt is an invalid UTC date</div> <div><strong>422</strong>: expiresAt cannot be in the past</div> <div><strong>500</strong>: Internal server error</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Customer id | 
  **locationId** | **string**| locationId | 
  **mac** | **string**| MAC address of the target device. Must have been online in the last 31 days. | 
 **optional** | ***CustomerApiCustomerPrototypeSetForcedSteerOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CustomerApiCustomerPrototypeSetForcedSteerOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **expiresAt** | **optional.String**| time of expiration in RFC 3339 format (e.g. 2021-11-24T09:13:33+00:00), must be between 5 and 60 minutes in the future. | 

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomerPrototypeSetPrimarySecondaryNetworks**
> CustomerPrototypeSetPrimarySecondaryNetworks(ctx, id, locationId, optional)
Set networks at wpa3 transition flow

<div><strong>202</strong>: Success, accepted the data</div> <div><strong>401</strong>: Authorization required or customer id not found</div> <div><strong>404</strong>: location id does not exist</div> <div><strong>422</strong>: Input validation failed.</div> <div><strong>500</strong>: Internal server error</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Customer id | 
  **locationId** | **string**|  | 
 **optional** | ***CustomerApiCustomerPrototypeSetPrimarySecondaryNetworksOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CustomerApiCustomerPrototypeSetPrimarySecondaryNetworksOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **wpa3ssid** | **optional.String**|  | 
 **wpa3encryptionKey** | **optional.String**|  | 
 **wpa3enabled** | **optional.Bool**|  | 
 **wpa2ssid** | **optional.String**|  | 
 **wpa2enabled** | **optional.Bool**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomerPrototypeShareDevice**
> interface{} CustomerPrototypeShareDevice(ctx, id, locationId, networkId, mac, groups, devices)
Share access to individual device. 

<p>This endpoint allows for a device in the first network to have access to all of the devices in the other group in the second network and/or to individual devices in the second network. In other words, by sharing access, we're allowing a single device to communicate with other devices across networks, by specifying other groups and/or individual devices.</p> <div><strong>200</strong>: Success.</div> <div><strong>422</strong>: Schema validation failed.</div> <div><strong>422</strong>: Illegal share.</div> <div><strong>404</strong>: Location does not exist.</div> <div><strong>404</strong>: Network does not exist.</div> <div><strong>404</strong>: Group does not exist.</div> <div><strong>401</strong>: Unauthorized.</div> <div><strong>400</strong>: Invalid JSON or missing arguments.</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Customer id | 
  **locationId** | **string**|  | 
  **networkId** | **string**|  | 
  **mac** | **string**|  | 
  **groups** | **string**|  | 
  **devices** | **string**|  | 

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomerPrototypeShareDeviceGroup**
> interface{} CustomerPrototypeShareDeviceGroup(ctx, id, locationId, networkId, groupId, groups, devices)
Share access for a group or employee.

<p>This endpoint allows for a device in the first network to have access to all of the devices in the other group in the second network and/or to individual devices in the second network. In other words, by sharing access, we're allowing a single device to communicate with other devices across networks, by specifying other groups and/or individual devices.</p> <div><strong>200</strong>: Success.</div> <div><strong>422</strong>: Schema validation failed.</div> <div><strong>422</strong>: Illegal share.</div> <div><strong>404</strong>: Location does not exist.</div> <div><strong>404</strong>: Network does not exist.</div> <div><strong>404</strong>: Group does not exist.</div> <div><strong>401</strong>: Unauthorized.</div> <div><strong>400</strong>: Invalid JSON or missing arguments.</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Customer id | 
  **locationId** | **string**|  | 
  **networkId** | **string**|  | 
  **groupId** | **string**|  | 
  **groups** | **string**|  | 
  **devices** | **string**|  | 

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomerPrototypeStartWps**
> XAny CustomerPrototypeStartWps(ctx, id, locationId, nodeId, optional)
Start a WPS session

<div><strong>201</strong>: Success, a WPS session was requested.</div> <div><strong>401</strong>: Authorization required </div> <div><strong>404</strong>: location id not found or nodeId missing from URL <div><strong>500</strong>: internal server error</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Customer id | 
  **locationId** | **string**|  | 
  **nodeId** | **string**|  | 
 **optional** | ***CustomerApiCustomerPrototypeStartWpsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CustomerApiCustomerPrototypeStartWpsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **keyId** | **optional.String**|  | 
 **networkId** | **optional.String**|  | 

### Return type

[**XAny**](x-any.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomerPrototypeStitchDevice**
> CustomerPrototypeStitchDevice(ctx, id, locationId, oldMac, newMac)
Delete prioritization of a single device

<div><strong>204</strong>: Success.</div> <div><strong>401</strong>: Unauthorized.</div> <div><strong>400</strong>: Missing oldMac or newMac field.</div> <div><strong>404</strong>: Location does not exist.</div> <div><strong>422</strong>: oldMac or newMac is not valid mac.</div> <div><strong>422</strong>: If oldMac and newMac are the same.</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Customer id | 
  **locationId** | **string**|  | 
  **oldMac** | **string**|  | 
  **newMac** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomerPrototypeUnapproveDevice**
> CustomerPrototypeUnapproveDevice(ctx, id, locationId, networkId, mac)
Unapprove approved devices in the network

<div><strong>204</strong>: Success.</div> <div><strong>404</strong>: Location does not exist.</div> <div><strong>404</strong>: Network does not exist.</div> <div><strong>404</strong>: Device is not approved.</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Customer id | 
  **locationId** | **string**|  | 
  **networkId** | **string**|  | 
  **mac** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomerPrototypeUnblockDevice**
> CustomerPrototypeUnblockDevice(ctx, id, locationId, mac)
Unblock blocked devices

<div><strong>204</strong>: Success.</div> <div><strong>404</strong>: Location does not exist.</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Customer id | 
  **locationId** | **string**|  | 
  **mac** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomerPrototypeUnclaimAllNodes**
> CustomerPrototypeUnclaimAllNodes(ctx, id, locationId, optional)
Unclaim all Nodes from a Location ID with the option of preserving the original Package ID.

<div><strong>204</strong>: Success, a job well done.</div> <div><strong>401</strong>: Authorization required, customer id not found, <br/> or id not owned by requestor.</div> <div><strong>404</strong>: location id not found in customer service or not found in inventory service.<p/>  <div><strong>500</strong>: internal server error</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Customer id | 
  **locationId** | **string**|  | 
 **optional** | ***CustomerApiCustomerPrototypeUnclaimAllNodesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CustomerApiCustomerPrototypeUnclaimAllNodesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **preservePackId** | **optional.Bool**| packId should remain the same | [default to false]
 **removeAccountId** | **optional.Bool**| delete account ids on the inventory nodes | [default to false]

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomerPrototypeUnclaimNode**
> CustomerPrototypeUnclaimNode(ctx, id, locationId, nodeId, optional)
Unclaim a particular Node from a Location ID with the option of preserving the original Package ID.

<div><strong>204</strong>: Success, a job well done.</div> <div><strong>400</strong>: Pod already unclaimed.</div> <div><strong>401</strong>: Authorization required, customer id not found, <br/> or id not owned by requestor.</div> <div><strong>403</strong>: the node is online, and can not be unclaimed.<br/>  <div><strong>404</strong>: location id not found, nodeId missing from URL,<br/> or location has zero owned pods.</div> <div><strong>500</strong>: internal server error</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Customer id | 
  **locationId** | **string**|  | 
  **nodeId** | **string**|  | 
 **optional** | ***CustomerApiCustomerPrototypeUnclaimNodeOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CustomerApiCustomerPrototypeUnclaimNodeOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **preservePackId** | **optional.Bool**| packId should remain the same | 
 **removeAccountId** | **optional.Bool**| delete account id on the inventory node | [default to false]

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomerPrototypeUpdateLocationName**
> XAny CustomerPrototypeUpdateLocationName(ctx, id, locationId, name)
Update the location name.

<div><strong>200</strong>: Success, updated.</div> <div><strong>400</strong>: Required fields missing.</div> <div><strong>401</strong>: Authorization required or customer id not found.</div> <div><strong>404</strong>: Location id, does not exist.</div> <div><strong>500</strong>: Internal server error</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Customer id | 
  **locationId** | **string**|  | 
  **name** | **string**|  | 

### Return type

[**XAny**](x-any.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomerPrototypeUpdateMigration**
> Migration CustomerPrototypeUpdateMigration(ctx, id, optional)
Update _migration of this model.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Customer id | 
 **optional** | ***CustomerApiCustomerPrototypeUpdateMigrationOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CustomerApiCustomerPrototypeUpdateMigrationOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **data** | [**optional.Interface of Migration**](Migration.md)|  | 

### Return type

[**Migration**](Migration.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomerPrototypeUpdateTermsAndPrivacy**
> TermsAndPrivacy CustomerPrototypeUpdateTermsAndPrivacy(ctx, id, optional)
Update a terms and privacy acceptance for customer.

<div><strong>200</strong>: Success, terms and privacy updated.</div> <div><strong>422</strong>: Input validation failed.</div> <div><strong>500</strong>: Internal server error.</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Customer id | 
 **optional** | ***CustomerApiCustomerPrototypeUpdateTermsAndPrivacyOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CustomerApiCustomerPrototypeUpdateTermsAndPrivacyOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **termsDocumentId** | **optional.Float64**|  | 
 **privacyDocumentId** | **optional.Float64**|  | 
 **termsAcceptedAt** | **optional.Time**|  | 
 **privacyAcceptedAt** | **optional.Time**|  | 

### Return type

[**TermsAndPrivacy**](TermsAndPrivacy.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomerPrototypeUserInfo**
> Customer CustomerPrototypeUserInfo(ctx, id)
Get customer details with userInfo access token.

<div><strong>200</strong>: Success, customer details returned.</div> <div><strong>401</strong>: Authorization Required.</div> <div><strong>500</strong>: Internal server error.</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Customer id | 

### Return type

[**Customer**](Customer.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomerPrototypeVerifyEmailPasswordlessToken**
> CustomerPrototypeVerifyEmailPasswordlessToken(ctx, id)
Verifies the email token and activates tokens related to it. Returns verified text with redirect to \"signup complete deep link\"

<div><strong>204</strong>: Success, return new appToken and send out the email with emailToken.</div> <div><strong>401</strong>: Authorization required or customer id not found</div> <div><strong>422</strong>: nodeId must be defined.</div> <div><strong>500</strong>: internal server error</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Customer id | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomerPrototypeVlanServices**
> XAny CustomerPrototypeVlanServices(ctx, id, locationId)
Returns vlanServices from Customer location state

<div><strong>200</strong>: Success.</div> <div><strong>401</strong>: Authorization required or customer id not found</div> <div><strong>404</strong>: Location id or WifiNetwork does not exist and is not known to Plume</div> <div><strong>500</strong>: Internal server error</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Customer id | 
  **locationId** | **string**|  | 

### Return type

[**XAny**](x-any.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomerPrototypeWpsState**
> XAny CustomerPrototypeWpsState(ctx, id, locationId)
Get WPS state

<div><strong>200</strong>: Success, a job well done.</div> <div><strong>401</strong>: Authorization required </div> <div><strong>404</strong>: location id not found or nodeId missing from URL <div><strong>500</strong>: internal server error</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Customer id | 
  **locationId** | **string**|  | 

### Return type

[**XAny**](x-any.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomerRefreshOauthAccessToken**
> interface{} CustomerRefreshOauthAccessToken(ctx, optional)
Refresh access and refresh tokens

<div><strong>200</strong>: Success, access and refresh tokens created and returned.</div> <div><strong>401</strong>: Authorization Required.</div> <div><strong>500</strong>: Internal server error.</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CustomerApiCustomerRefreshOauthAccessTokenOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CustomerApiCustomerRefreshOauthAccessTokenOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **refreshToken** | **optional.String**|  | 
 **clientId** | **optional.String**|  | 

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomerRegister**
> RegisterResponse CustomerRegister(ctx, accountId, optional)
Register/create an anonymous account with an accountId instead of with email/password.

<div><strong>200</strong>: Success, customer and location IDs returned.</div> <div><strong>400</strong>: Required fields are missing.</div> <div><strong>401</strong>: Authorization required.</div> <div><strong>422</strong>: Input validation failed.</div> <div><strong>422</strong>: Only integration role can set profile to property.</div> <div><strong>500</strong>: Internal server error.</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountId** | **string**| must be unique, a UUID is recommended, min length is 6 characters. | 
 **optional** | ***CustomerApiCustomerRegisterOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CustomerApiCustomerRegisterOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **name** | **optional.String**| Full name of customer, defaults to value of accountId | 
 **partnerId** | **optional.String**| PartnerId of customer for accountId | 
 **email** | **optional.String**|  | 
 **acceptLanguage** | **optional.String**| acceptable language | 
 **profile** | **optional.String**| location profile | 
 **onboardingCheckpoint** | **optional.String**| is the last passed onboarding step by the customer: &#39;PodsAdded&#39; or &#39;OnboardingComplete&#39;; | 

### Return type

[**RegisterResponse**](RegisterResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomerRegisterWithGroups**
> RegisterResponse CustomerRegisterWithGroups(ctx, email, accountId, name, optional)
Register/create an account with an accountId plus email/password/groups.

<div><strong>200</strong>: Success, customer and location IDs returned.</div> <div><strong>400</strong>: Required fields are missing.</div> <div><strong>401</strong>: Authorization required.</div> <div><strong>422</strong>: Input validation failed.</div> <div><strong>500</strong>: Internal server error.</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **email** | **string**|  | 
  **accountId** | **string**| must be unique, a UUID is recommended, min length is 6 characters. | 
  **name** | **string**| Full name of customer, defaults to value of accountId | 
 **optional** | ***CustomerApiCustomerRegisterWithGroupsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CustomerApiCustomerRegisterWithGroupsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **password** | **optional.String**|  | 
 **groupIds** | **optional.String**| at least one groupId | 
 **partnerId** | **optional.String**| PartnerId of customer for accountId | 
 **acceptLanguage** | **optional.String**| acceptable language | 
 **profile** | **optional.String**| location profile | 

### Return type

[**RegisterResponse**](RegisterResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomerResendEmailVerification**
> CustomerResendEmailVerification(ctx, optional)
Resend the verification email.

<div><strong>204</strong>: Successfully sent email verification.</div> <div><strong>400</strong>: Customer email is required (for this request).</div> <div><strong>404</strong>: Unable to find Customer by email address.</div> <div><strong>409</strong>: Customer email already verified.</div> <div><strong>500</strong>: Internal server error.</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CustomerApiCustomerResendEmailVerificationOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CustomerApiCustomerResendEmailVerificationOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **email** | **optional.String**| Email address that verification email will be sent to. | 
 **notificationOptions** | **optional.String**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomerResetPassword**
> CustomerResetPassword(ctx, options)
Reset password for a user with email.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **options** | **interface{}**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomerSearchFields**
> XAny CustomerSearchFields(ctx, keyword, field, optional)
Search the keyword on a particular field such as \"accountId\", \"name\", \"email\".

<div><strong>200</strong>: Success, return the search result.</div> <div><strong>401</strong>: Authorization required or customer id not found</div> <div><strong>422</strong>: \"illegal field\"</div> <div><strong>500</strong>: internal server error</div>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **keyword** | **string**|  | 
  **field** | **string**|  | 
 **optional** | ***CustomerApiCustomerSearchFieldsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CustomerApiCustomerSearchFieldsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **exactMatch** | **optional.Bool**|  | 
 **startsWith** | **optional.Bool**|  | 

### Return type

[**XAny**](x-any.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-www-form-urlencoded, application/xml, text/xml
 - **Accept**: application/json, application/xml, text/xml, application/javascript, text/javascript

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

