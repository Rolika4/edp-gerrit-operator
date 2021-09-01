package gerritproject

import (
	"strings"
	"testing"
	"time"

	"github.com/epam/edp-gerrit-operator/v2/pkg/apis/v2/v1alpha1"
	gerritClient "github.com/epam/edp-gerrit-operator/v2/pkg/client/gerrit"
	"github.com/epam/edp-gerrit-operator/v2/pkg/controller/helper"
	gerritService "github.com/epam/edp-gerrit-operator/v2/pkg/service/gerrit"
	"github.com/pkg/errors"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	"sigs.k8s.io/controller-runtime/pkg/client/fake"
)

func TestSyncBackendProjectsTick(t *testing.T) {
	scheme := runtime.NewScheme()
	utilruntime.Must(v1alpha1.AddToScheme(scheme))
	utilruntime.Must(corev1.AddToScheme(scheme))

	g := v1alpha1.Gerrit{
		ObjectMeta: metav1.ObjectMeta{
			Namespace: "ns", Name: "ger1"},
		TypeMeta: metav1.TypeMeta{
			APIVersion: "v2.edp.epam.com/v1alpha1",
			Kind:       "Gerrit",
		}}

	prj := v1alpha1.GerritProject{
		ObjectMeta: metav1.ObjectMeta{Namespace: "ns", Name: "prj1",
			OwnerReferences: []metav1.OwnerReference{
				{
					Kind: g.Kind,
					UID:  g.UID,
				},
			}},
		Spec: v1alpha1.GerritProjectSpec{Name: "sprj1"},
	}

	cl := fake.NewClientBuilder().WithScheme(scheme).WithRuntimeObjects(&g, &prj).Build()
	serviceMock := gerritService.Mock{}
	clientMock := gerritClient.Mock{}
	serviceMock.On("GetRestClient", &g).Return(&clientMock, nil)

	logger := helper.Logger{}

	rcn := Reconcile{
		client:  cl,
		log:     &logger,
		service: &serviceMock,
	}

	clientMock.On("ListProjects", "CODE").Return([]gerritClient.Project{
		{
			Name: "alphabet",
		},
	}, nil)
	clientMock.On("ListProjectBranches", "alphabet").Return([]gerritClient.Branch{
		{
			Ref: "test",
		},
	}, nil)

	if err := rcn.syncBackendProjectsTick(); err != nil {
		t.Fatal(err)
	}
}

func TestSyncBackendProjectsTick_BranchesFailure(t *testing.T) {
	scheme := runtime.NewScheme()
	utilruntime.Must(v1alpha1.AddToScheme(scheme))
	utilruntime.Must(corev1.AddToScheme(scheme))

	g := v1alpha1.Gerrit{
		ObjectMeta: metav1.ObjectMeta{
			Namespace: "ns", Name: "ger1"},
		TypeMeta: metav1.TypeMeta{
			APIVersion: "v2.edp.epam.com/v1alpha1",
			Kind:       "Gerrit",
		}}

	prj := v1alpha1.GerritProject{
		ObjectMeta: metav1.ObjectMeta{Namespace: "ns", Name: "prj1",
			OwnerReferences: []metav1.OwnerReference{
				{
					Kind: g.Kind,
					UID:  g.UID,
				},
			}},
		Spec: v1alpha1.GerritProjectSpec{Name: "sprj1"},
	}

	cl := fake.NewClientBuilder().WithScheme(scheme).WithRuntimeObjects(&g, &prj).Build()
	serviceMock := gerritService.Mock{}
	clientMock := gerritClient.Mock{}
	serviceMock.On("GetRestClient", &g).Return(&clientMock, nil)

	logger := helper.Logger{}

	rcn := Reconcile{
		client:  cl,
		log:     &logger,
		service: &serviceMock,
	}

	clientMock.On("ListProjects", "CODE").Return([]gerritClient.Project{
		{
			Name: "alphabet",
		},
	}, nil)
	clientMock.On("ListProjectBranches", "alphabet").Return(nil, errors.New("list branches fatal"))

	err := rcn.syncBackendProjectsTick()
	if err == nil {
		t.Fatal("no error returned")
	}

	if !strings.Contains(err.Error(), "list branches fatal") {
		t.Fatalf("wrong error returned: %s", err.Error())
	}
}

func TestSyncBackendProjectsTick_Failure(t *testing.T) {
	scheme := runtime.NewScheme()
	utilruntime.Must(v1alpha1.AddToScheme(scheme))
	utilruntime.Must(corev1.AddToScheme(scheme))

	g := v1alpha1.Gerrit{
		ObjectMeta: metav1.ObjectMeta{
			Namespace: "ns", Name: "ger1"},
		TypeMeta: metav1.TypeMeta{
			APIVersion: "v2.edp.epam.com/v1alpha1",
			Kind:       "Gerrit",
		}}

	cl := fake.NewClientBuilder().WithScheme(scheme).WithRuntimeObjects(&g).Build()
	serviceMock := gerritService.Mock{}
	serviceMock.On("GetRestClient", &g).
		Return(nil, errors.New("gerrit client fatal")).Once()

	rcn := Reconcile{
		client:  cl,
		service: &serviceMock,
	}

	err := rcn.syncBackendProjectsTick()
	if err == nil {
		t.Fatal("no error returned")
	}
	if !strings.Contains(err.Error(), "gerrit client fatal") {
		t.Fatalf("wrong error returned: %s", err.Error())
	}

	clientMock := gerritClient.Mock{}
	serviceMock.On("GetRestClient", &g).Return(&clientMock, nil)

	clientMock.On("ListProjects", "CODE").
		Return(nil, errors.New("list projects fatal"))

	err = rcn.syncBackendProjectsTick()
	if err == nil {
		t.Fatal("no error returned")
	}
	if !strings.Contains(err.Error(), "list projects fatal") {
		t.Fatalf("wrong error returned: %s", err.Error())
	}
}

func TestSyncBackendProjects(t *testing.T) {
	scheme := runtime.NewScheme()
	utilruntime.Must(v1alpha1.AddToScheme(scheme))
	utilruntime.Must(corev1.AddToScheme(scheme))

	g := v1alpha1.Gerrit{
		ObjectMeta: metav1.ObjectMeta{
			Namespace: "ns", Name: "ger1"},
		TypeMeta: metav1.TypeMeta{
			APIVersion: "v2.edp.epam.com/v1alpha1",
			Kind:       "Gerrit",
		}}

	cl := fake.NewClientBuilder().WithScheme(scheme).WithRuntimeObjects(&g).Build()
	serviceMock := gerritService.Mock{}
	serviceMock.On("GetRestClient", &g).
		Return(nil, errors.New("gerrit client fatal"))

	logger := helper.Logger{}

	rcn := Reconcile{
		client:  cl,
		service: &serviceMock,
		log:     &logger,
	}

	go rcn.syncBackendProjects(time.Millisecond)
	time.Sleep(time.Millisecond * 10)

	err := logger.LastError()
	if err == nil {
		t.Fatal("no error returned")
	}

	if !strings.Contains(err.Error(), "gerrit client fatal") {
		t.Fatalf("wrong error returned: %s", err.Error())
	}
}