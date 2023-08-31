package services

type ConfessionsStoreStub struct {
	saveCalled       bool
	randomCalled     bool
	deleteByIDCalled bool
	forcedError      error
}

type PostedConfessionsStoreStub struct {
	MarkConfessionAsPublishedCalled bool
	IsConfessionPublishedCalled     bool
}
