package chrome

import (
	"app/chrome/protocol"
	"encoding/json"

	log "github.com/Sirupsen/logrus"
)

/*
LayerTree - https://chromedevtools.github.io/devtools-protocol/tot/LayerTree/
EXPERIMENTAL
*/
type LayerTree struct{}

/*
CompositingReasons provides the reasons why the given layer was composited.
*/
func (LayerTree) CompositingReasons(
	socket *Socket,
	params *layer_tree.CompositingReasonsParams,
) (layer_tree.CompositingReasonsResult, error) {
	command := &protocol.Command{
		method: "LayerTree.compositingReasons",
		params: params,
	}
	socket.SendCommand(command)
	return layer_tree.CompositingReasonsResult, command.Err
}

/*
Disable disables compositing tree inspection.
*/
func (LayerTree) Disable(
	socket *Socket,
) (nil, error) {
	command := &protocol.Command{
		method: "LayerTree.disable",
		params: params,
	}
	socket.SendCommand(command)
	return nil, command.Err
}

/*
Enable enables compositing tree inspection.
*/
func (LayerTree) Enable(
	socket *Socket,
) (nil, error) {
	command := &protocol.Command{
		method: "LayerTree.enable",
		params: params,
	}
	socket.SendCommand(command)
	return nil, command.Err
}

/*
LoadSnapshot returns the snapshot identifier.
*/
func (LayerTree) LoadSnapshot(
	socket *Socket,
	params *layer_tree.LoadSnapshotParams,
) (layer_tree.LoadSnapshotResult, error) {
	command := &protocol.Command{
		method: "LayerTree.loadSnapshot",
		params: params,
	}
	socket.SendCommand(command)
	return layer_tree.LoadSnapshotResult, command.Err
}

/*
MakeSnapshot returns the layer snapshot identifier.
*/
func (LayerTree) MakeSnapshot(
	socket *Socket,
	params *layer_tree.MakeSnapshotParams,
) (layer_tree.MakeSnapshotResult, error) {
	command := &protocol.Command{
		method: "LayerTree.makeSnapshot",
		params: params,
	}
	socket.SendCommand(command)
	return layer_tree.MakeSnapshotResult, command.Err
}

/*
ProfileSnapshot profiles a snapshot.
*/
func (LayerTree) ProfileSnapshot(
	socket *Socket,
	params *layer_tree.ProfileSnapshotParams,
) (layer_tree.ProfileSnapshotResult, error) {
	command := &protocol.Command{
		method: "LayerTree.profileSnapshot",
		params: params,
	}
	socket.SendCommand(command)
	return layer_tree.ProfileSnapshotResult, command.Err
}

/*
ReleaseSnapshot releases layer snapshot captured by the back-end.
*/
func (LayerTree) ReleaseSnapshot(
	socket *Socket,
	params *layer_tree.ReleaseSnapshotParams,
) (nil, error) {
	command := &protocol.Command{
		method: "LayerTree.releaseSnapshot",
		params: params,
	}
	socket.SendCommand(command)
	return nil, command.Err
}

/*
ReplaySnapshot replays the layer snapshot and returns the resulting bitmap.
*/
func (LayerTree) ReplaySnapshot(
	socket *Socket,
	params *layer_tree.ReplaySnapshotParams,
) (layer_tree.ReplaySnapshotResult, error) {
	command := &protocol.Command{
		method: "LayerTree.replaySnapshot",
		params: params,
	}
	socket.SendCommand(command)
	return layer_tree.ReplaySnapshotResult, command.Err
}

/*
SnapshotCommandLog replays the layer snapshot and returns canvas log.
*/
func (LayerTree) SnapshotCommandLog(
	socket *Socket,
	params *layer_tree.SnapshotCommandLogParams,
) (layer_tree.SnapshotCommandLogResult, error) {
	command := &protocol.Command{
		method: "LayerTree.snapshotCommandLog",
		params: params,
	}
	socket.SendCommand(command)
	return layer_tree.SnapshotCommandLogResult, command.Err
}

/*
OnLayerPainted adds a handler to the LayerTree.layerPainted event. LayerTree.layerPainted fires when the layer
is painted.
*/
func (LayerTree) OnLayerPainted(
	socket *Socket,
	callback func(event *layer_tree.LayerPaintedEvent),
) {
	handler := protocol.NewEventHandler(
		"LayerTree.layerPainted",
		func(name string, params []byte) {
			event := &layer_tree.LayerPaintedEvent{}
			if err := json.Unmarshal(params, event); err != nil {
				log.Error(err)
			} else {
				callback(event)
			}
		},
	)
	socket.AddEventHandler(handler)
}

/*
OnLayerTreeDidChange adds a handler to the LayerTree.layerTreeDidChange event. LayerTree.layerTreeDidChange
fires when the layer tree changes.
*/
func (LayerTree) OnLayerTreeDidChange(
	socket *Socket,
	callback func(event *layer_tree.LayerTreeDidChangeEvent),
) {
	handler := protocol.NewEventHandler(
		"LayerTree.layerTreeDidChange",
		func(name string, params []byte) {
			event := &layer_tree.LayerTreeDidChangeEvent{}
			if err := json.Unmarshal(params, event); err != nil {
				log.Error(err)
			} else {
				callback(event)
			}
		},
	)
	socket.AddEventHandler(handler)
}
