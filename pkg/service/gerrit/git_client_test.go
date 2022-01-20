package gerrit

import (
	"context"
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/client/fake"

	mock "github.com/epam/edp-gerrit-operator/v2/mock/platform"
	"github.com/epam/edp-gerrit-operator/v2/pkg/apis/v2/v1alpha1"
)

type testChild struct{}

func (testChild) GetNamespace() string {
	return "ns"
}

func (testChild) OwnerName() string {
	return "gerrit"
}

func TestComponentService_GetGitClient_Failure(t *testing.T) {
	sch := runtime.NewScheme()
	plt := mock.PlatformService{}
	plt.AssertExpectations(t)

	err := corev1.AddToScheme(sch)
	assert.NoError(t, err)
	v1alpha1.RegisterTypes(sch)

	s := ComponentService{
		PlatformService: &plt,
		k8sScheme:       sch,
		client:          fake.NewClientBuilder().WithScheme(sch).Build(),
	}

	testCh := testChild{}

	_, err = s.GetGitClient(context.Background(), testCh, "")
	assert.Error(t, err)
	assert.EqualError(t, err,
		"unable to get parent gerrit: gerrits.v2.edp.epam.com \"gerrit\" not found")

	rootGerrit := v1alpha1.Gerrit{ObjectMeta: metav1.ObjectMeta{Name: testCh.OwnerName(),
		Namespace: testCh.GetNamespace()}}
	s.client = fake.NewClientBuilder().WithScheme(sch).WithRuntimeObjects(&rootGerrit).Build()
	plt.On("GetSecretData", testCh.GetNamespace(), fmt.Sprintf("%v-admin-password", rootGerrit.Name)).
		Return(nil, errors.New("secret fatal")).Once()

	_, err = s.GetGitClient(context.Background(), testCh, "")
	assert.Error(t, err)
	assert.EqualError(t, err,
		"Failed to get Gerrit admin password from secret for ns/gerrit: Failed to get Secret gerrit-admin-password for ns/gerrit: secret fatal")
}

func TestComponentService_GetGitClient(t *testing.T) {
	sch := runtime.NewScheme()
	plt := mock.PlatformService{}
	plt.AssertExpectations(t)

	err := corev1.AddToScheme(sch)
	assert.NoError(t, err)
	v1alpha1.RegisterTypes(sch)

	testCh := testChild{}
	rootGerrit := v1alpha1.Gerrit{ObjectMeta: metav1.ObjectMeta{Name: testCh.OwnerName(),
		Namespace: testCh.GetNamespace()}}

	s := ComponentService{
		PlatformService: &plt,
		k8sScheme:       sch,
		client:          fake.NewClientBuilder().WithScheme(sch).WithRuntimeObjects(&rootGerrit).Build(),
	}

	plt.On("GetSecretData", testCh.GetNamespace(), fmt.Sprintf("%v-admin-password", rootGerrit.Name)).
		Return(map[string][]byte{"password": []byte("secret")}, nil)
	plt.On("GetExternalEndpoint", testCh.GetNamespace(), testCh.OwnerName()).
		Return("", "", nil)

	_, err = s.GetGitClient(context.Background(), testCh, "")
	assert.NoError(t, err)
}
