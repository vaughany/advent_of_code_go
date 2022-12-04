package inputs

import "embed"

//go:embed 2015/*.txt
//go:embed 2016/*.txt
//go:embed 2017/*.txt
//go:embed 2018/*.txt
//go:embed 2019/*.txt
//go:embed 2020/*.txt
//go:embed 2021/*.txt
//go:embed 2022/*.txt
var EmbeddedFiles embed.FS
