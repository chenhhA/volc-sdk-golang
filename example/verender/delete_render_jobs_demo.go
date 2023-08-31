package main

import (
    "fmt"

    "github.com/volcengine/volc-sdk-golang/service/verender"
)

func DeleteRenderJobsDemo() {
    v := getVerenderInstance()
    workspaceId := int64(1935)
    w, err := v.GetWorkspace(workspaceId)
    if err != nil {
        fmt.Println(err.Error())
        return
    }

    r := verender.BatchJobIDList{
        JobIDList: []string{"rb3bf60e847", "r24d6fbecae"},
    }

    if err := w.DeleteRenderJobs(&r); err != nil {
        fmt.Println(err.Error())
        return
    }

    fmt.Println("DeleteRenderJobs done")
}
