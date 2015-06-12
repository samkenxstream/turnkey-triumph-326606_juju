// Copyright 2015 Canonical Ltd.
// Licensed under the AGPLv3, see LICENCE file for details.

package lease

// These constants define the field names and type values used by documents in
// a lease storage collection. They *must* remain in sync with the bson
// marshallling annotations in leaseDoc and clockDoc.
const (
	// fieldType and fieldNamespace identify the Type and Namespace fields in
	// both leaseDoc and clockDoc.
	fieldType      = "type"
	fieldNamespace = "namespace"

	// typeLease and typeClock are the acceptable values for fieldType.
	typeLease = "lease"
	typeClock = "clock"

	// fieldLease* identify the fields in a leaseDoc.
	fieldLeaseLease   = "lease"
	fieldLeaseHolder  = "holder"
	fieldLeaseExpiry  = "expiry"
	fieldLeaseWriter  = "writer"
	fieldLeaseWritten = "written"

	// fieldClock* identify the fields in a clockDoc.
	fieldClockWriters = "writers"
)

// leaseDocId returns the _id for the document holding details of the supplied
// namespace and lease.
func leaseDocId(namespace, lease string) string {
	return fmt.Sprintf("%s#%s#%s#", typeLease, namespace, lease)
}

// leaseDoc is used to serialise lease entries.
type leaseDoc struct {
	// Id is always "<Type>#<Namespace>#<Lease>#", and <Type> is always "lease",
	// so that we can extract useful information from a stream of watcher events
	// without incurring extra DB hits.
	// Apart from checking validity on load, though, there's little reason
	// to use Id elsewhere; Namespace and Lease are the sources of truth.
	Id        string `bson:"_id"`
	Type      string `bson:"type"`      // TODO(fwereade) add index
	Namespace string `bson:"namespace"` // TODO(fwereade) add index
	Lease     string `bson:"lease"`     // TODO(fwereade) add index?

	// Holder, Expiry, Writer and Written map directly to entry. The time values
	// are stored as UnixNano; not so much because we *need* the precision, as
	// because it's yucky when serialisation throws precision away, and life is
	// easier when we can compare leases exactly.
	Holder  string `bson:"holder"`
	Expiry  int64  `bson:"expiry"`
	Writer  string `bson:"writer"`
	Written int64  `bson:"written"`
}

// validate returns an error if any fields are invalid or inconsistent.
func (doc leaseDoc) validate() error {
	return fmt.Errorf("validation!")
}

// entry returns the lease name and an entry corresponding to the document. If
// the document cannot be validated, it returns an error.
func (doc leaseDoc) entry() (string, entry, error) {
	if err := doc.validate(); err != nil {
		return "", entry{}, errors.Trace(err)
	}
	entry := entry{
		holder:  doc.Holder,
		expiry:  time.Unix(0, doc.Expiry),
		writer:  doc.Writer,
		written: time.Unix(0, doc.Written),
	}
	return doc.Lease, entry, nil
}

// newLeaseDoc returns a valid lease document encoding the supplied lease and
// entry in the supplied namespace, or an error.
func newLeaseDoc(namespace, lease string, entry entry) (*leaseDoc, error) {
	doc := &leaseDoc{
		Id:        s.leaseDocId(lease),
		Namespace: s.config.Namespace,
		Lease:     lease,
		Holder:    entry.holder,
		Expiry:    entry.expiry.UnixNano(),
		Writer:    entry.writer,
		Written:   entry.written.UnixNano(),
	}
	if err := doc.validate(); err != nil {
		return nil, errors.Trace(err)
	}
	return doc, nil
}

// clockDocId returns the _id for the document holding clock skew information
// for the storage instances in the supplied namespace.
func clockDocId(namespace string) string {
	return fmt.Sprintf("%s#%s#", typeClock, namespace)
}

// clockDoc is used to synchronise different storage instances.
type clockDoc struct {
	// Id is always "<Type>#<Namespace>#", and <Type> is always "clock", for
	// consistency with leaseDoc and ease of querying within the collection.
	Id        string `bson:"_id"`
	Type      string `bson:"type"`
	Namespace string `bson:"namespace"`

	// Writers holds a map of storage instances to the most recent time they
	// acknowledge has already passed, stored as UnixNano.
	Writers map[string]int64 `bson:"writers"`
}

// validate returns an error if any fields are invalid or inconsistent.
func (doc clockDoc) validate() error {
	return fmt.Errorf("validation!")
}

// skews returns clock skew information for all writers recorded in the
// document, given that the document was read between the supplied local
// times. It will return an error if the clock document is not valid, or
// if the times don't make sense.
func (doc clockDoc) skews(readAfter, readBefore time.Time) (map[string]Skew, error) {
	if err := doc.validate(); err != nil {
		return nil, errors.Trace(err)
	}
	if readBefore.Before(readAfter) {
		return nil, errors.New("end of read window preceded beginning")
	}
	skews := make(map[string]Skew)
	for writer, written := range doc.Writers {
		skews[writer] = Skew{
			LastWrite:  time.Unix(0, written),
			ReadAfter:  readAfter,
			ReadBefore: readBefore,
		}
	}
	return skews, nil
}

// newClockDoc returns an empty clockDoc for the supplied namespace.
func newClockDoc(namespace string) (clockDoc, error) {
	doc := clockDoc{
		Id:        clockDocId(namespace),
		Type:      typeClock,
		Namespace: namespace,
		Writers:   make(map[string]int64),
	}
	if err := doc.validate(); err != nil {
		return nil, errors.Trace(err)
	}
	return doc, nil
}
