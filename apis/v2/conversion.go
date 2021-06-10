package v2

import (
	"sigs.k8s.io/controller-runtime/pkg/conversion"

	v1 "github.com/aws-controllers-k8s/ecr-controller/apis/v1"
)

// ConvertTo converts this Repository to the Hub version (v1).
func (src *Repository) ConvertTo(dstRaw conversion.Hub) {
	dst := dstRaw.(*v1.Repository)
	dst.ObjectMeta = src.ObjectMeta

	// Copy scalar types
	dst.Spec.Name = src.Spec.Name
	dst.Spec.ImageTagMutability = src.Spec.ImageTagMutability

	// Copy struct using type conversion
	dst.Spec.EncryptionConfiguration = (*v1.EncryptionConfiguration)(src.Spec.EncryptionConfiguration)

	// Copy struct the hard way
	if src.Spec.ScanConfig != nil {
		dst.Spec.ImageScanningConfiguration = (*v1.ImageScanningConfiguration)(src.Spec.ScanConfig)
	}

	// Copying arrays is messier
	tags := make([]*v1.Tag, 0, len(src.Spec.Tags))
	for _, t := range src.Spec.Tags {
		tags = append(tags, (*v1.Tag)(t))
	}
	dst.Spec.Tags = tags

	// Copy status
	dst.Status = v1.RepositoryStatus(src.Status)
}

// ConvertFrom converts from the Hub version (v1) to this version.
func (dst *Repository) ConvertFrom(srcRaw conversion.Hub) {
	src := srcRaw.(*v1.Repository)
	dst.ObjectMeta = src.ObjectMeta

	// Copy scalar types
	dst.Spec.Name = src.Spec.Name
	dst.Spec.ImageTagMutability = src.Spec.ImageTagMutability

	// Copy struct using type conversion
	dst.Spec.EncryptionConfiguration = (*EncryptionConfiguration)(src.Spec.EncryptionConfiguration)

	// Copy struct the hard way
	if src.Spec.ImageScanningConfiguration != nil {
		dst.Spec.ScanConfig = (*ImageScanningConfiguration)(src.Spec.ImageScanningConfiguration)
	}

	// Copying arrays is messier
	tags := make([]*Tag, 0, len(src.Spec.Tags))
	for _, t := range src.Spec.Tags {
		tags = append(tags, (*Tag)(t))
	}
	dst.Spec.Tags = tags

	// Copy status
	dst.Status = RepositoryStatus(src.Status)
}
