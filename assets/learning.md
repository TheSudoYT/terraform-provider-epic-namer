# Docs and Stuff I Learned

## Autogeneration of Docs

Go here https://github.com/hashicorp/terraform-plugin-docs and follow the README.md

`go generate ./...`

## Terraform.rc

```hcl
plugin_cache_dir = "C:\\Users\\Joshua\\AppData\\Roaming\\terraform.d\\plugins\\cachce"
disable_checkpoint = true

provider_installation {
  filesystem_mirror {
    path    = "C:\\Users\\Joshua\\AppData\\Roaming\\terraform.d\\plugins"
    include = ["localhost/providers/epic"]
  }
  direct {
    exclude = ["localhost/providers/epic"]
  }
}
```

## Local Testing

`go build -o terraform-provider-epic_v0.1.0`

`mv .\terraform-provider-epic_v0.1.0 C:\Users\Joshua\AppData\Roaming\terraform.d\plugins\localhost\providers\epic\0.1.0\windows_amd64\`