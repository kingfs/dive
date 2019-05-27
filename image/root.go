package image

type AnalyzerFactory func(string) Analyzer

func GetAnalyzer(imageID string) Analyzer {
	// todo: add ability to have multiple image formats... for the meantime only use docker
	var factory AnalyzerFactory = newDockerImageAnalyzer

	return factory(imageID)
}
