package dashboardtypes

// DashboardParent is an interface implemented by all dashboard run nodes which have children
type DashboardParent interface {
	GetName() string
	ChildCompleteChan() chan DashboardTreeRun
	GetChildren() []DashboardTreeRun
	ChildrenComplete() bool
}
