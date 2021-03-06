package main

import (
	"fmt"
	"github.com/HewlettPackard/oneview-golang/ov"
	"os"
)

func main() {
	var (
		clientOV           *ov.OVClient
		enc_name           = "0000A66101"
		new_enclosure_name = "RenamedEnclosure"
		path               = "/name"
		op                 = "replace"
	)
	ovc := clientOV.NewOVClient(
		os.Getenv("ONEVIEW_OV_USER"),
		os.Getenv("ONEVIEW_OV_PASSWORD"),
		os.Getenv("ONEVIEW_OV_DOMAIN"),
		os.Getenv("ONEVIEW_OV_ENDPOINT"),
		false,
		1800,
		"*")

	enclosure_create_map := ov.EnclosureCreateMap{
		EnclosureGroupUri: "/rest/enclosure_groups/05100faa-c26b-4a16-8055-911568418190",
		Hostname:          os.Getenv("ENCLOSURE_HOSTNAME"),
		Username:          os.Getenv("ENCLOSURE_USERNAME"),
		Password:          os.Getenv("ENCLOSURE_PASSWORD"),
		LicensingIntent:   "OneView",
		InitialScopeUris:  make([]string, 0),
	}

	fmt.Println("#----------------Create Enclosure---------------#")

	err := ovc.CreateEnclosure(enclosure_create_map)
	if err != nil {
		fmt.Println("Enclosure Creation Failed: ", err)
	} else {
		fmt.Println("Enclosure created successfully...")
	}

	sort := ""

	enc_list, err := ovc.GetEnclosures("", "", "", sort, "")
	if err != nil {
		fmt.Println("Enclosure Retrieval Failed: ", err)
	} else {
		fmt.Println("#----------------Enclosure List---------------#")

		for i := 0; i < len(enc_list.Members); i++ {
			fmt.Println(enc_list.Members[i].Name)
		}
	}

	enclosure, err := ovc.GetEnclosureByName(enc_name)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("#----------------Enclosure by Name----------------#")
		fmt.Println(enclosure.Name)
	}

	uri := enclosure.URI
	enclosure, err = ovc.GetEnclosurebyUri(uri)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("#----------------Enclosure by URI--------------#")
		fmt.Println(enclosure.Name)
	}

	err = ovc.UpdateEnclosure(op, path, new_enclosure_name, enclosure)
	if err != nil {
		fmt.Println(err)
	}

	enc_list, err = ovc.GetEnclosures("", "", "", sort, "")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("#----------------Enclosure List after Updating---------#")
		for i := 0; i < len(enc_list.Members); i++ {
			fmt.Println(enc_list.Members[i].Name)
		}
	}

	err = ovc.DeleteEnclosure(new_enclosure_name)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Deleted Enclosure successfully...")
	}
}
