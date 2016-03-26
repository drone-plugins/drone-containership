package main

type Params struct {
    ApiKey       string `json:"api_key"`
    Application  string `json:"application"`
    ClusterId    string `json:"cluster_id"`
    Image        string `json:"docker_image"`
    Organization string `json:"organization"`
}
