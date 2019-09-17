package structs

type (
    DelugePayload struct {
        Id string `json:"id" validate:"required,strint"`
        Dir string `json:"dir"  validate:"required,linuxpath"`
        Name string `json:"name" validate:"required"`
    }

    NzbGetPayload struct {
        Dir string `json:"dir" validate:"required,linuxpath"`
        Status string `json:"status" validate:"required,nzbgetstatus"`
        Category string `json:"category" validate:"required,nzbgetcategory"`
    }
)
