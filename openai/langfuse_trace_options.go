package openai

type LangfuseTraceOption func(*LangfuseTraceOptions)

type LangfuseTraceOptions struct {
	TraceID             string
	ParentObservationID string
	SessionID           string
	UserID              string
	Metadata            map[string]any
	Tags                []string
	Release             string
}

func WithTraceID(traceID string) LangfuseTraceOption {
	return func(o *LangfuseTraceOptions) {
		o.TraceID = traceID
	}
}

func WithParentObservationID(parentObservationID string) LangfuseTraceOption {
	return func(o *LangfuseTraceOptions) {
		o.ParentObservationID = parentObservationID
	}
}

func WithSessionID(sessionID string) LangfuseTraceOption {
	return func(o *LangfuseTraceOptions) {
		o.SessionID = sessionID
	}
}

func WithUserID(userID string) LangfuseTraceOption {
	return func(o *LangfuseTraceOptions) {
		o.UserID = userID
	}
}

func WithMetadata(metadata map[string]any) LangfuseTraceOption {
	return func(o *LangfuseTraceOptions) {
		o.Metadata = metadata
	}
}

func WithTags(tags []string) LangfuseTraceOption {
	return func(o *LangfuseTraceOptions) {
		o.Tags = tags
	}
}

func WithRelease(release string) LangfuseTraceOption {
	return func(o *LangfuseTraceOptions) {
		o.Release = release
	}
}

func ApplyTraceOptions(options []LangfuseTraceOption) *LangfuseTraceOptions {
	opts := &LangfuseTraceOptions{}
	for _, opt := range options {
		opt(opts)
	}
	return opts
}
