// Code generated by "callbackgen -type S3"; DO NOT EDIT.

package fmaker

import ()

func (inc *S3) OnUpdate(cb func(val float64)) {
	inc.UpdateCallbacks = append(inc.UpdateCallbacks, cb)
}

func (inc *S3) EmitUpdate(val float64) {
	for _, cb := range inc.UpdateCallbacks {
		cb(val)
	}
}