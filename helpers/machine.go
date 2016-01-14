package helpers

import (
	"encoding/json"
	"strings"
	"errors"
)

func uploadMachineSchema(drivers []string) error {
	err := uploadMachineServiceJson(drivers)
	err2 := uploadMachineProjectJson(drivers)
	err3 := uploadMachineUserJson(drivers)
	if err != nil || err2 != nil || err3 != nil {
		return errors.New("Failed to upload one of the machine jsons.")
	}
	return nil
}


func genFieldSchema(resourceFieldStruct ResourceFieldConfigs, field, fieldType, auth string) {
	resourceFieldStruct[field] = ResourceFieldConfig{
		Nullable: true,
		Type: fieldType,
		Create:strings.Contains(auth, "c"),
		Update:strings.Contains(auth, "u"),
	}
}


func uploadMachineServiceJson(drivers []string) error {
	resourceFieldStruct := make(map[string]interface{})
	resourceFieldMap := make(ResourceFieldConfigs)
	resourceFieldStruct["collectionMethods"] = []string{"GET", "POST", "PUT", "DELETE"}
	resourceFieldStruct["resourceMethods"] = []string{"GET", "PUT", "DELETE"}
	resourceFieldStruct["resourceFields"] = resourceFieldMap
	for _, driver := range drivers {
		genFieldSchema(resourceFieldMap, driver + "Config", driver + "Config", "c")
	}
	genFieldSchema(resourceFieldMap, "authCertificateAuthority", "string","c")
	genFieldSchema(resourceFieldMap, "authKey", "string","c")
	genFieldSchema(resourceFieldMap, "data", "json","cu")
	genFieldSchema(resourceFieldMap, "dockerVersion", "string","c")
	genFieldSchema(resourceFieldMap, "driver", "string","")
	genFieldSchema(resourceFieldMap, "engineEnv", "map[string]","c")
	genFieldSchema(resourceFieldMap, "engineInsecureRegistry", "array[string]","c")
	genFieldSchema(resourceFieldMap, "engineInstallUrl", "string","c")
	genFieldSchema(resourceFieldMap, "engineLabel", "map[string]","c")
	genFieldSchema(resourceFieldMap, "engineOpt", "map[string]","c")
	genFieldSchema(resourceFieldMap, "engineRegistryMirror", "array[string]","c")
	genFieldSchema(resourceFieldMap, "engineStorageDriver", "string","c")
	genFieldSchema(resourceFieldMap, "externalId", "string","")
	genFieldSchema(resourceFieldMap, "extractedConfig", "string","u")
	genFieldSchema(resourceFieldMap, "labels", "map[string]","cu")

	jsonData, err := json.MarshalIndent(resourceFieldStruct, "", "    ")
	if err != nil {
		return err
	}
	return uploadDynamicSchema("machine", string(jsonData), "physicalHost", []string{"service"})
}

func uploadMachineProjectJson(drivers []string) error {
	resourceFieldStruct := make(map[string]interface{})
	resourceFieldMap := make(ResourceFieldConfigs)
	resourceFieldStruct["collectionMethods"] = []string{"GET", "POST", "DELETE"}
	resourceFieldStruct["resourceMethods"] = []string{"GET", "DELETE"}
	resourceFieldStruct["resourceFields"] = resourceFieldMap
	for _, driver := range drivers {
		genFieldSchema(resourceFieldMap, driver + "Config", driver + "Config", "c")
	}
	genFieldSchema(resourceFieldMap, "authCertificateAuthority", "string","c")
	genFieldSchema(resourceFieldMap, "authKey", "string","c")
	genFieldSchema(resourceFieldMap, "dockerVersion", "string","c")
	genFieldSchema(resourceFieldMap, "driver", "string","")
	genFieldSchema(resourceFieldMap, "engineEnv", "map[string]","c")
	genFieldSchema(resourceFieldMap, "engineInsecureRegistry", "array[string]","c")
	genFieldSchema(resourceFieldMap, "engineInstallUrl", "string","c")
	genFieldSchema(resourceFieldMap, "engineLabel", "map[string]","c")
	genFieldSchema(resourceFieldMap, "engineOpt", "map[string]","c")
	genFieldSchema(resourceFieldMap, "engineRegistryMirror", "array[string]","c")
	genFieldSchema(resourceFieldMap, "engineStorageDriver", "string","c")
	genFieldSchema(resourceFieldMap, "externalId", "string","")
	genFieldSchema(resourceFieldMap, "labels", "map[string]","c")

	jsonData, err := json.MarshalIndent(resourceFieldStruct, "", "    ")
	if err != nil {
		return err
	}
	return uploadDynamicSchema("machine", string(jsonData), "physicalHost",
		[]string{"project", "member", "owner"})
}

func uploadMachineUserJson(drivers []string) error {
	resourceFieldStruct := make(map[string]interface{})
	resourceFieldMap := make(ResourceFieldConfigs)
	resourceFieldStruct["collectionMethods"] = []string{"GET"}
	resourceFieldStruct["resourceMethods"] = []string{"GET"}
	resourceFieldStruct["resourceFields"] = resourceFieldMap
	for _, driver := range drivers {
		genFieldSchema(resourceFieldMap, driver + "Config", driver + "Config", "")
	}
	genFieldSchema(resourceFieldMap, "authCertificateAuthority", "string","")
	genFieldSchema(resourceFieldMap, "authKey", "string","")
	genFieldSchema(resourceFieldMap, "dockerVersion", "string","")
	genFieldSchema(resourceFieldMap, "driver", "string","")
	genFieldSchema(resourceFieldMap, "engineEnv", "map[string]","")
	genFieldSchema(resourceFieldMap, "engineInsecureRegistry", "array[string]","")
	genFieldSchema(resourceFieldMap, "engineInstallUrl", "string","")
	genFieldSchema(resourceFieldMap, "engineLabel", "map[string]","")
	genFieldSchema(resourceFieldMap, "engineOpt", "map[string]","")
	genFieldSchema(resourceFieldMap, "engineRegistryMirror", "array[string]","")
	genFieldSchema(resourceFieldMap, "engineStorageDriver", "string","")
	genFieldSchema(resourceFieldMap, "externalId", "string","")
	genFieldSchema(resourceFieldMap, "labels", "map[string]","")

	jsonData, err := json.MarshalIndent(resourceFieldStruct, "", "    ")
	if err != nil {
		return err
	}
	return uploadDynamicSchema("machine", string(jsonData), "physicalHost", []string{"admin", "user", "readAdmin"})
}