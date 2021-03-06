package main

import (
	"fmt"
	"github.com/HewlettPackard/oneview-golang/ov"
	"os"
	"time"
)

func main() {

	var (
		ClientOV       *ov.OVClient
		name_to_create = "VolumeTemplateExample"
		name_to_update = "VolumeTemplateExample- updated"
	)

	ovc := ClientOV.NewOVClient(
		os.Getenv("ONEVIEW_OV_USER"),
		os.Getenv("ONEVIEW_OV_PASSWORD"),
		os.Getenv("ONEVIEW_OV_DOMAIN"),
		os.Getenv("ONEVIEW_OV_ENDPOINT"),
		false,
		1800,
		"*")

	name_properties := ov.TemplatePropertyDatatypeStructString{
		Required:    true,
		Type:        "string",
		Title:       "Volume name",
		Description: "A volume name between 1 and 100 characters",
		Maxlength:   100,
		Minlength:   1,
		Meta: &ov.Meta{
			Locked: false,
		},
	}

	storage_pool_properties := ov.TemplatePropertyDatatypeStructString{
		Required:    true,
		Type:        "string",
		Title:       "Storage Pool",
		Description: "StoragePoolURI the volume should be added to",
		Default:     "/rest/storage-pools/7D2A17CD-414A-45EF-B221-ABF1006D6F44",
		Meta: &ov.Meta{
			Locked:       false,
			CreateOnly:   true,
			SemanticType: "device-storage-pool",
		},
		Format: "x-uri-reference",
	}

	size_properties := ov.TemplatePropertyDatatypeStructInt{
		Required:    true,
		Type:        "integer",
		Title:       "Capacity",
		Default:     1073741824,
		Minimum:     4194304,
		Description: "Capacity of the volume in bytes",
		Meta: &ov.Meta{
			Locked:       false,
			SemanticType: "capacity",
		},
	}

	dataProtectionLevel_properties := ov.TemplatePropertyDatatypeStructString{
		Required: true,
		Type:     "string",
		Enum: []string{"NetworkRaid0None",
			"NetworkRaid5SingleParity",
			"NetworkRaid10Mirror2Way",
			"NetworkRaid10Mirror3Way",
			"NetworkRaid10Mirror4Way",
			"NetworkRaid6DualParity",
		},
		Title:       "Data Protection Level",
		Default:     "NetworkRaid10Mirror2Way",
		Description: "Indicates the number and configuration of data copies in the Storage Pool",
		Meta: &ov.Meta{
			Locked:       false,
			SemanticType: "device-dataProtectionLevel",
		},
	}

	template_version_properties := ov.TemplatePropertyDatatypeStructString{
		Required:    true,
		Type:        "string",
		Title:       "Template version",
		Description: "Version of the template",
		Default:     "1.1",
		Meta: &ov.Meta{
			Locked: true,
		},
	}

	description_properties := ov.TemplatePropertyDatatypeStructString{
		Required:    false,
		Type:        "string",
		Title:       "Description",
		Description: "A description for the volume",
		Default:     "A description for the volume",
		Maxlength:   2000,
		Minlength:   1,
		Meta: &ov.Meta{
			Locked: false,
		},
	}

	provisioning_type_properties := ov.TemplatePropertyDatatypeStructString{
		Required:    false,
		Title:       "Provisioning Type",
		Type:        "string",
		Description: "The provisioning type for the volume",
		Default:     "Thin",
		Enum:        []string{"Thin", "Full"},
		Meta: &ov.Meta{
			Locked:       true,
			CreateOnly:   true,
			SemanticType: "device-provisioningType",
		},
	}

	adaptive_optimization_properties := ov.TemplatePropertyDatatypeStructBool{
		Meta: &ov.Meta{
			Locked: true,
		},
		Type:        "boolean",
		Description: "",
		Default:     true,
		Required:    false,
		Title:       "Adaptive Optimization",
	}

	is_shareable_properties := ov.TemplatePropertyDatatypeStructBool{
		Meta: &ov.Meta{
			Locked: true,
		},
		Type:        "boolean",
		Description: "The shareability of the volume",
		Default:     true,
		Required:    false,
		Title:       "Is Shareable",
	}

	Properties := ov.TemplateProperties{
		Name:                          &name_properties,
		StoragePool:                   &storage_pool_properties,
		Size:                          &size_properties,
		DataProtectionLevel:           &dataProtectionLevel_properties,
		TemplateVersion:               &template_version_properties,
		Description:                   &description_properties,
		ProvisioningType:              &provisioning_type_properties,
		IsAdaptiveOptimizationEnabled: &adaptive_optimization_properties,
		IsShareable:                   &is_shareable_properties,
	}

	storageVolumeTemplate := ov.StorageVolumeTemplate{
		TemplateProperties: &Properties,
		Name:               name_to_create,
		Description:        "Volume template Example",
		RootTemplateUri:    "/rest/storage-volume-templates/11f5fb2e-068a-4fe3-a1a6-abf1006d6f65",
	}

	err := ovc.CreateStorageVolumeTemplate(storageVolumeTemplate)
	if err != nil {
		fmt.Println("Could not create the volume Template", err)
	} else {
		fmt.Println("Volume Template created successfully", storageVolumeTemplate.Name)
	}

	// Get the volume template by name
	update_vol_template, err := ovc.GetStorageVolumeTemplateByName(name_to_create)
	if err != nil {
		fmt.Println(err)
	} else {
		update_vol_template.Name = name_to_update
		update_vol_template.Description = "Updating description"

		// Update the previously created storage volume template
		err = ovc.UpdateStorageVolumeTemplate(update_vol_template)
		if err != nil {
			fmt.Println("Could not update the volume template", err)
		} else {
			fmt.Println("Volume template updated")
		}
	}
	time.Sleep(2 * time.Second)

	// Get All the volume templates present
	fmt.Println("\nGetting all the storage volume templates present in the system: \n")
	sort := "name:desc"
	vol_temp_list, err := ovc.GetStorageVolumeTemplates("", sort, "", "")
	if err != nil {
		fmt.Println("Error Getting the storage volume templates ", err)
	}
	for i := 0; i < len(vol_temp_list.Members); i++ {
		fmt.Println(vol_temp_list.Members[i].Name)
	}

	// Delete the created volume template
	fmt.Println("\nDeleting the volume with name : ", name_to_update)
	err = ovc.DeleteStorageVolumeTemplate(name_to_update)
	if err != nil {
		fmt.Println("Delete Unsuccessful", err)
	}

}
