// Code generated by adapter-generator. DO NOT EDIT.

package adapter

import (
	"github.com/golang/protobuf/proto"
	. "go.ligato.io/vpp-agent/v2/plugins/kvscheduler/api"
	"go.ligato.io/vpp-agent/v2/proto/ligato/vpp/nat"
)

////////// type-safe key-value pair with metadata //////////

type NAT44GlobalAddressKVWithMetadata struct {
	Key      string
	Value    *vpp_nat.Nat44Global_Address
	Metadata interface{}
	Origin   ValueOrigin
}

////////// type-safe Descriptor structure //////////

type NAT44GlobalAddressDescriptor struct {
	Name                 string
	KeySelector          KeySelector
	ValueTypeName        string
	KeyLabel             func(key string) string
	ValueComparator      func(key string, oldValue, newValue *vpp_nat.Nat44Global_Address) bool
	NBKeyPrefix          string
	WithMetadata         bool
	MetadataMapFactory   MetadataMapFactory
	Validate             func(key string, value *vpp_nat.Nat44Global_Address) error
	Create               func(key string, value *vpp_nat.Nat44Global_Address) (metadata interface{}, err error)
	Delete               func(key string, value *vpp_nat.Nat44Global_Address, metadata interface{}) error
	Update               func(key string, oldValue, newValue *vpp_nat.Nat44Global_Address, oldMetadata interface{}) (newMetadata interface{}, err error)
	UpdateWithRecreate   func(key string, oldValue, newValue *vpp_nat.Nat44Global_Address, metadata interface{}) bool
	Retrieve             func(correlate []NAT44GlobalAddressKVWithMetadata) ([]NAT44GlobalAddressKVWithMetadata, error)
	IsRetriableFailure   func(err error) bool
	DerivedValues        func(key string, value *vpp_nat.Nat44Global_Address) []KeyValuePair
	Dependencies         func(key string, value *vpp_nat.Nat44Global_Address) []Dependency
	RetrieveDependencies []string /* descriptor name */
}

////////// Descriptor adapter //////////

type NAT44GlobalAddressDescriptorAdapter struct {
	descriptor *NAT44GlobalAddressDescriptor
}

func NewNAT44GlobalAddressDescriptor(typedDescriptor *NAT44GlobalAddressDescriptor) *KVDescriptor {
	adapter := &NAT44GlobalAddressDescriptorAdapter{descriptor: typedDescriptor}
	descriptor := &KVDescriptor{
		Name:                 typedDescriptor.Name,
		KeySelector:          typedDescriptor.KeySelector,
		ValueTypeName:        typedDescriptor.ValueTypeName,
		KeyLabel:             typedDescriptor.KeyLabel,
		NBKeyPrefix:          typedDescriptor.NBKeyPrefix,
		WithMetadata:         typedDescriptor.WithMetadata,
		MetadataMapFactory:   typedDescriptor.MetadataMapFactory,
		IsRetriableFailure:   typedDescriptor.IsRetriableFailure,
		RetrieveDependencies: typedDescriptor.RetrieveDependencies,
	}
	if typedDescriptor.ValueComparator != nil {
		descriptor.ValueComparator = adapter.ValueComparator
	}
	if typedDescriptor.Validate != nil {
		descriptor.Validate = adapter.Validate
	}
	if typedDescriptor.Create != nil {
		descriptor.Create = adapter.Create
	}
	if typedDescriptor.Delete != nil {
		descriptor.Delete = adapter.Delete
	}
	if typedDescriptor.Update != nil {
		descriptor.Update = adapter.Update
	}
	if typedDescriptor.UpdateWithRecreate != nil {
		descriptor.UpdateWithRecreate = adapter.UpdateWithRecreate
	}
	if typedDescriptor.Retrieve != nil {
		descriptor.Retrieve = adapter.Retrieve
	}
	if typedDescriptor.Dependencies != nil {
		descriptor.Dependencies = adapter.Dependencies
	}
	if typedDescriptor.DerivedValues != nil {
		descriptor.DerivedValues = adapter.DerivedValues
	}
	return descriptor
}

func (da *NAT44GlobalAddressDescriptorAdapter) ValueComparator(key string, oldValue, newValue proto.Message) bool {
	typedOldValue, err1 := castNAT44GlobalAddressValue(key, oldValue)
	typedNewValue, err2 := castNAT44GlobalAddressValue(key, newValue)
	if err1 != nil || err2 != nil {
		return false
	}
	return da.descriptor.ValueComparator(key, typedOldValue, typedNewValue)
}

func (da *NAT44GlobalAddressDescriptorAdapter) Validate(key string, value proto.Message) (err error) {
	typedValue, err := castNAT44GlobalAddressValue(key, value)
	if err != nil {
		return err
	}
	return da.descriptor.Validate(key, typedValue)
}

func (da *NAT44GlobalAddressDescriptorAdapter) Create(key string, value proto.Message) (metadata Metadata, err error) {
	typedValue, err := castNAT44GlobalAddressValue(key, value)
	if err != nil {
		return nil, err
	}
	return da.descriptor.Create(key, typedValue)
}

func (da *NAT44GlobalAddressDescriptorAdapter) Update(key string, oldValue, newValue proto.Message, oldMetadata Metadata) (newMetadata Metadata, err error) {
	oldTypedValue, err := castNAT44GlobalAddressValue(key, oldValue)
	if err != nil {
		return nil, err
	}
	newTypedValue, err := castNAT44GlobalAddressValue(key, newValue)
	if err != nil {
		return nil, err
	}
	typedOldMetadata, err := castNAT44GlobalAddressMetadata(key, oldMetadata)
	if err != nil {
		return nil, err
	}
	return da.descriptor.Update(key, oldTypedValue, newTypedValue, typedOldMetadata)
}

func (da *NAT44GlobalAddressDescriptorAdapter) Delete(key string, value proto.Message, metadata Metadata) error {
	typedValue, err := castNAT44GlobalAddressValue(key, value)
	if err != nil {
		return err
	}
	typedMetadata, err := castNAT44GlobalAddressMetadata(key, metadata)
	if err != nil {
		return err
	}
	return da.descriptor.Delete(key, typedValue, typedMetadata)
}

func (da *NAT44GlobalAddressDescriptorAdapter) UpdateWithRecreate(key string, oldValue, newValue proto.Message, metadata Metadata) bool {
	oldTypedValue, err := castNAT44GlobalAddressValue(key, oldValue)
	if err != nil {
		return true
	}
	newTypedValue, err := castNAT44GlobalAddressValue(key, newValue)
	if err != nil {
		return true
	}
	typedMetadata, err := castNAT44GlobalAddressMetadata(key, metadata)
	if err != nil {
		return true
	}
	return da.descriptor.UpdateWithRecreate(key, oldTypedValue, newTypedValue, typedMetadata)
}

func (da *NAT44GlobalAddressDescriptorAdapter) Retrieve(correlate []KVWithMetadata) ([]KVWithMetadata, error) {
	var correlateWithType []NAT44GlobalAddressKVWithMetadata
	for _, kvpair := range correlate {
		typedValue, err := castNAT44GlobalAddressValue(kvpair.Key, kvpair.Value)
		if err != nil {
			continue
		}
		typedMetadata, err := castNAT44GlobalAddressMetadata(kvpair.Key, kvpair.Metadata)
		if err != nil {
			continue
		}
		correlateWithType = append(correlateWithType,
			NAT44GlobalAddressKVWithMetadata{
				Key:      kvpair.Key,
				Value:    typedValue,
				Metadata: typedMetadata,
				Origin:   kvpair.Origin,
			})
	}

	typedValues, err := da.descriptor.Retrieve(correlateWithType)
	if err != nil {
		return nil, err
	}
	var values []KVWithMetadata
	for _, typedKVWithMetadata := range typedValues {
		kvWithMetadata := KVWithMetadata{
			Key:      typedKVWithMetadata.Key,
			Metadata: typedKVWithMetadata.Metadata,
			Origin:   typedKVWithMetadata.Origin,
		}
		kvWithMetadata.Value = typedKVWithMetadata.Value
		values = append(values, kvWithMetadata)
	}
	return values, err
}

func (da *NAT44GlobalAddressDescriptorAdapter) DerivedValues(key string, value proto.Message) []KeyValuePair {
	typedValue, err := castNAT44GlobalAddressValue(key, value)
	if err != nil {
		return nil
	}
	return da.descriptor.DerivedValues(key, typedValue)
}

func (da *NAT44GlobalAddressDescriptorAdapter) Dependencies(key string, value proto.Message) []Dependency {
	typedValue, err := castNAT44GlobalAddressValue(key, value)
	if err != nil {
		return nil
	}
	return da.descriptor.Dependencies(key, typedValue)
}

////////// Helper methods //////////

func castNAT44GlobalAddressValue(key string, value proto.Message) (*vpp_nat.Nat44Global_Address, error) {
	typedValue, ok := value.(*vpp_nat.Nat44Global_Address)
	if !ok {
		return nil, ErrInvalidValueType(key, value)
	}
	return typedValue, nil
}

func castNAT44GlobalAddressMetadata(key string, metadata Metadata) (interface{}, error) {
	if metadata == nil {
		return nil, nil
	}
	typedMetadata, ok := metadata.(interface{})
	if !ok {
		return nil, ErrInvalidMetadataType(key)
	}
	return typedMetadata, nil
}