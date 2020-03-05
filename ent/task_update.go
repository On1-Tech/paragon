// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/kcarretto/paragon/ent/job"
	"github.com/kcarretto/paragon/ent/predicate"
	"github.com/kcarretto/paragon/ent/tag"
	"github.com/kcarretto/paragon/ent/target"
	"github.com/kcarretto/paragon/ent/task"
)

// TaskUpdate is the builder for updating Task entities.
type TaskUpdate struct {
	config
	QueueTime          *time.Time
	LastChangedTime    *time.Time
	ClaimTime          *time.Time
	clearClaimTime     bool
	ExecStartTime      *time.Time
	clearExecStartTime bool
	ExecStopTime       *time.Time
	clearExecStopTime  bool
	Content            *string
	Output             *string
	clearOutput        bool
	Error              *string
	clearError         bool
	SessionID          *string
	clearSessionID     bool
	tags               map[int]struct{}
	job                map[int]struct{}
	target             map[int]struct{}
	removedTags        map[int]struct{}
	clearedJob         bool
	clearedTarget      bool
	predicates         []predicate.Task
}

// Where adds a new predicate for the builder.
func (tu *TaskUpdate) Where(ps ...predicate.Task) *TaskUpdate {
	tu.predicates = append(tu.predicates, ps...)
	return tu
}

// SetQueueTime sets the QueueTime field.
func (tu *TaskUpdate) SetQueueTime(t time.Time) *TaskUpdate {
	tu.QueueTime = &t
	return tu
}

// SetNillableQueueTime sets the QueueTime field if the given value is not nil.
func (tu *TaskUpdate) SetNillableQueueTime(t *time.Time) *TaskUpdate {
	if t != nil {
		tu.SetQueueTime(*t)
	}
	return tu
}

// SetLastChangedTime sets the LastChangedTime field.
func (tu *TaskUpdate) SetLastChangedTime(t time.Time) *TaskUpdate {
	tu.LastChangedTime = &t
	return tu
}

// SetClaimTime sets the ClaimTime field.
func (tu *TaskUpdate) SetClaimTime(t time.Time) *TaskUpdate {
	tu.ClaimTime = &t
	return tu
}

// SetNillableClaimTime sets the ClaimTime field if the given value is not nil.
func (tu *TaskUpdate) SetNillableClaimTime(t *time.Time) *TaskUpdate {
	if t != nil {
		tu.SetClaimTime(*t)
	}
	return tu
}

// ClearClaimTime clears the value of ClaimTime.
func (tu *TaskUpdate) ClearClaimTime() *TaskUpdate {
	tu.ClaimTime = nil
	tu.clearClaimTime = true
	return tu
}

// SetExecStartTime sets the ExecStartTime field.
func (tu *TaskUpdate) SetExecStartTime(t time.Time) *TaskUpdate {
	tu.ExecStartTime = &t
	return tu
}

// SetNillableExecStartTime sets the ExecStartTime field if the given value is not nil.
func (tu *TaskUpdate) SetNillableExecStartTime(t *time.Time) *TaskUpdate {
	if t != nil {
		tu.SetExecStartTime(*t)
	}
	return tu
}

// ClearExecStartTime clears the value of ExecStartTime.
func (tu *TaskUpdate) ClearExecStartTime() *TaskUpdate {
	tu.ExecStartTime = nil
	tu.clearExecStartTime = true
	return tu
}

// SetExecStopTime sets the ExecStopTime field.
func (tu *TaskUpdate) SetExecStopTime(t time.Time) *TaskUpdate {
	tu.ExecStopTime = &t
	return tu
}

// SetNillableExecStopTime sets the ExecStopTime field if the given value is not nil.
func (tu *TaskUpdate) SetNillableExecStopTime(t *time.Time) *TaskUpdate {
	if t != nil {
		tu.SetExecStopTime(*t)
	}
	return tu
}

// ClearExecStopTime clears the value of ExecStopTime.
func (tu *TaskUpdate) ClearExecStopTime() *TaskUpdate {
	tu.ExecStopTime = nil
	tu.clearExecStopTime = true
	return tu
}

// SetContent sets the Content field.
func (tu *TaskUpdate) SetContent(s string) *TaskUpdate {
	tu.Content = &s
	return tu
}

// SetOutput sets the Output field.
func (tu *TaskUpdate) SetOutput(s string) *TaskUpdate {
	tu.Output = &s
	return tu
}

// SetNillableOutput sets the Output field if the given value is not nil.
func (tu *TaskUpdate) SetNillableOutput(s *string) *TaskUpdate {
	if s != nil {
		tu.SetOutput(*s)
	}
	return tu
}

// ClearOutput clears the value of Output.
func (tu *TaskUpdate) ClearOutput() *TaskUpdate {
	tu.Output = nil
	tu.clearOutput = true
	return tu
}

// SetError sets the Error field.
func (tu *TaskUpdate) SetError(s string) *TaskUpdate {
	tu.Error = &s
	return tu
}

// SetNillableError sets the Error field if the given value is not nil.
func (tu *TaskUpdate) SetNillableError(s *string) *TaskUpdate {
	if s != nil {
		tu.SetError(*s)
	}
	return tu
}

// ClearError clears the value of Error.
func (tu *TaskUpdate) ClearError() *TaskUpdate {
	tu.Error = nil
	tu.clearError = true
	return tu
}

// SetSessionID sets the SessionID field.
func (tu *TaskUpdate) SetSessionID(s string) *TaskUpdate {
	tu.SessionID = &s
	return tu
}

// SetNillableSessionID sets the SessionID field if the given value is not nil.
func (tu *TaskUpdate) SetNillableSessionID(s *string) *TaskUpdate {
	if s != nil {
		tu.SetSessionID(*s)
	}
	return tu
}

// ClearSessionID clears the value of SessionID.
func (tu *TaskUpdate) ClearSessionID() *TaskUpdate {
	tu.SessionID = nil
	tu.clearSessionID = true
	return tu
}

// AddTagIDs adds the tags edge to Tag by ids.
func (tu *TaskUpdate) AddTagIDs(ids ...int) *TaskUpdate {
	if tu.tags == nil {
		tu.tags = make(map[int]struct{})
	}
	for i := range ids {
		tu.tags[ids[i]] = struct{}{}
	}
	return tu
}

// AddTags adds the tags edges to Tag.
func (tu *TaskUpdate) AddTags(t ...*Tag) *TaskUpdate {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return tu.AddTagIDs(ids...)
}

// SetJobID sets the job edge to Job by id.
func (tu *TaskUpdate) SetJobID(id int) *TaskUpdate {
	if tu.job == nil {
		tu.job = make(map[int]struct{})
	}
	tu.job[id] = struct{}{}
	return tu
}

// SetJob sets the job edge to Job.
func (tu *TaskUpdate) SetJob(j *Job) *TaskUpdate {
	return tu.SetJobID(j.ID)
}

// SetTargetID sets the target edge to Target by id.
func (tu *TaskUpdate) SetTargetID(id int) *TaskUpdate {
	if tu.target == nil {
		tu.target = make(map[int]struct{})
	}
	tu.target[id] = struct{}{}
	return tu
}

// SetNillableTargetID sets the target edge to Target by id if the given value is not nil.
func (tu *TaskUpdate) SetNillableTargetID(id *int) *TaskUpdate {
	if id != nil {
		tu = tu.SetTargetID(*id)
	}
	return tu
}

// SetTarget sets the target edge to Target.
func (tu *TaskUpdate) SetTarget(t *Target) *TaskUpdate {
	return tu.SetTargetID(t.ID)
}

// RemoveTagIDs removes the tags edge to Tag by ids.
func (tu *TaskUpdate) RemoveTagIDs(ids ...int) *TaskUpdate {
	if tu.removedTags == nil {
		tu.removedTags = make(map[int]struct{})
	}
	for i := range ids {
		tu.removedTags[ids[i]] = struct{}{}
	}
	return tu
}

// RemoveTags removes tags edges to Tag.
func (tu *TaskUpdate) RemoveTags(t ...*Tag) *TaskUpdate {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return tu.RemoveTagIDs(ids...)
}

// ClearJob clears the job edge to Job.
func (tu *TaskUpdate) ClearJob() *TaskUpdate {
	tu.clearedJob = true
	return tu
}

// ClearTarget clears the target edge to Target.
func (tu *TaskUpdate) ClearTarget() *TaskUpdate {
	tu.clearedTarget = true
	return tu
}

// Save executes the query and returns the number of rows/vertices matched by this operation.
func (tu *TaskUpdate) Save(ctx context.Context) (int, error) {
	if tu.Content != nil {
		if err := task.ContentValidator(*tu.Content); err != nil {
			return 0, fmt.Errorf("ent: validator failed for field \"Content\": %v", err)
		}
	}
	if tu.SessionID != nil {
		if err := task.SessionIDValidator(*tu.SessionID); err != nil {
			return 0, fmt.Errorf("ent: validator failed for field \"SessionID\": %v", err)
		}
	}
	if len(tu.job) > 1 {
		return 0, errors.New("ent: multiple assignments on a unique edge \"job\"")
	}
	if tu.clearedJob && tu.job == nil {
		return 0, errors.New("ent: clearing a unique edge \"job\"")
	}
	if len(tu.target) > 1 {
		return 0, errors.New("ent: multiple assignments on a unique edge \"target\"")
	}
	return tu.sqlSave(ctx)
}

// SaveX is like Save, but panics if an error occurs.
func (tu *TaskUpdate) SaveX(ctx context.Context) int {
	affected, err := tu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (tu *TaskUpdate) Exec(ctx context.Context) error {
	_, err := tu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tu *TaskUpdate) ExecX(ctx context.Context) {
	if err := tu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (tu *TaskUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   task.Table,
			Columns: task.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: task.FieldID,
			},
		},
	}
	if ps := tu.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value := tu.QueueTime; value != nil {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  *value,
			Column: task.FieldQueueTime,
		})
	}
	if value := tu.LastChangedTime; value != nil {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  *value,
			Column: task.FieldLastChangedTime,
		})
	}
	if value := tu.ClaimTime; value != nil {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  *value,
			Column: task.FieldClaimTime,
		})
	}
	if tu.clearClaimTime {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: task.FieldClaimTime,
		})
	}
	if value := tu.ExecStartTime; value != nil {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  *value,
			Column: task.FieldExecStartTime,
		})
	}
	if tu.clearExecStartTime {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: task.FieldExecStartTime,
		})
	}
	if value := tu.ExecStopTime; value != nil {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  *value,
			Column: task.FieldExecStopTime,
		})
	}
	if tu.clearExecStopTime {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: task.FieldExecStopTime,
		})
	}
	if value := tu.Content; value != nil {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  *value,
			Column: task.FieldContent,
		})
	}
	if value := tu.Output; value != nil {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  *value,
			Column: task.FieldOutput,
		})
	}
	if tu.clearOutput {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: task.FieldOutput,
		})
	}
	if value := tu.Error; value != nil {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  *value,
			Column: task.FieldError,
		})
	}
	if tu.clearError {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: task.FieldError,
		})
	}
	if value := tu.SessionID; value != nil {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  *value,
			Column: task.FieldSessionID,
		})
	}
	if tu.clearSessionID {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: task.FieldSessionID,
		})
	}
	if nodes := tu.removedTags; len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   task.TagsTable,
			Columns: task.TagsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: tag.FieldID,
				},
			},
		}
		for k, _ := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tu.tags; len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   task.TagsTable,
			Columns: task.TagsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: tag.FieldID,
				},
			},
		}
		for k, _ := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if tu.clearedJob {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   task.JobTable,
			Columns: []string{task.JobColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: job.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tu.job; len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   task.JobTable,
			Columns: []string{task.JobColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: job.FieldID,
				},
			},
		}
		for k, _ := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if tu.clearedTarget {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   task.TargetTable,
			Columns: []string{task.TargetColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: target.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tu.target; len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   task.TargetTable,
			Columns: []string{task.TargetColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: target.FieldID,
				},
			},
		}
		for k, _ := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, tu.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// TaskUpdateOne is the builder for updating a single Task entity.
type TaskUpdateOne struct {
	config
	id                 int
	QueueTime          *time.Time
	LastChangedTime    *time.Time
	ClaimTime          *time.Time
	clearClaimTime     bool
	ExecStartTime      *time.Time
	clearExecStartTime bool
	ExecStopTime       *time.Time
	clearExecStopTime  bool
	Content            *string
	Output             *string
	clearOutput        bool
	Error              *string
	clearError         bool
	SessionID          *string
	clearSessionID     bool
	tags               map[int]struct{}
	job                map[int]struct{}
	target             map[int]struct{}
	removedTags        map[int]struct{}
	clearedJob         bool
	clearedTarget      bool
}

// SetQueueTime sets the QueueTime field.
func (tuo *TaskUpdateOne) SetQueueTime(t time.Time) *TaskUpdateOne {
	tuo.QueueTime = &t
	return tuo
}

// SetNillableQueueTime sets the QueueTime field if the given value is not nil.
func (tuo *TaskUpdateOne) SetNillableQueueTime(t *time.Time) *TaskUpdateOne {
	if t != nil {
		tuo.SetQueueTime(*t)
	}
	return tuo
}

// SetLastChangedTime sets the LastChangedTime field.
func (tuo *TaskUpdateOne) SetLastChangedTime(t time.Time) *TaskUpdateOne {
	tuo.LastChangedTime = &t
	return tuo
}

// SetClaimTime sets the ClaimTime field.
func (tuo *TaskUpdateOne) SetClaimTime(t time.Time) *TaskUpdateOne {
	tuo.ClaimTime = &t
	return tuo
}

// SetNillableClaimTime sets the ClaimTime field if the given value is not nil.
func (tuo *TaskUpdateOne) SetNillableClaimTime(t *time.Time) *TaskUpdateOne {
	if t != nil {
		tuo.SetClaimTime(*t)
	}
	return tuo
}

// ClearClaimTime clears the value of ClaimTime.
func (tuo *TaskUpdateOne) ClearClaimTime() *TaskUpdateOne {
	tuo.ClaimTime = nil
	tuo.clearClaimTime = true
	return tuo
}

// SetExecStartTime sets the ExecStartTime field.
func (tuo *TaskUpdateOne) SetExecStartTime(t time.Time) *TaskUpdateOne {
	tuo.ExecStartTime = &t
	return tuo
}

// SetNillableExecStartTime sets the ExecStartTime field if the given value is not nil.
func (tuo *TaskUpdateOne) SetNillableExecStartTime(t *time.Time) *TaskUpdateOne {
	if t != nil {
		tuo.SetExecStartTime(*t)
	}
	return tuo
}

// ClearExecStartTime clears the value of ExecStartTime.
func (tuo *TaskUpdateOne) ClearExecStartTime() *TaskUpdateOne {
	tuo.ExecStartTime = nil
	tuo.clearExecStartTime = true
	return tuo
}

// SetExecStopTime sets the ExecStopTime field.
func (tuo *TaskUpdateOne) SetExecStopTime(t time.Time) *TaskUpdateOne {
	tuo.ExecStopTime = &t
	return tuo
}

// SetNillableExecStopTime sets the ExecStopTime field if the given value is not nil.
func (tuo *TaskUpdateOne) SetNillableExecStopTime(t *time.Time) *TaskUpdateOne {
	if t != nil {
		tuo.SetExecStopTime(*t)
	}
	return tuo
}

// ClearExecStopTime clears the value of ExecStopTime.
func (tuo *TaskUpdateOne) ClearExecStopTime() *TaskUpdateOne {
	tuo.ExecStopTime = nil
	tuo.clearExecStopTime = true
	return tuo
}

// SetContent sets the Content field.
func (tuo *TaskUpdateOne) SetContent(s string) *TaskUpdateOne {
	tuo.Content = &s
	return tuo
}

// SetOutput sets the Output field.
func (tuo *TaskUpdateOne) SetOutput(s string) *TaskUpdateOne {
	tuo.Output = &s
	return tuo
}

// SetNillableOutput sets the Output field if the given value is not nil.
func (tuo *TaskUpdateOne) SetNillableOutput(s *string) *TaskUpdateOne {
	if s != nil {
		tuo.SetOutput(*s)
	}
	return tuo
}

// ClearOutput clears the value of Output.
func (tuo *TaskUpdateOne) ClearOutput() *TaskUpdateOne {
	tuo.Output = nil
	tuo.clearOutput = true
	return tuo
}

// SetError sets the Error field.
func (tuo *TaskUpdateOne) SetError(s string) *TaskUpdateOne {
	tuo.Error = &s
	return tuo
}

// SetNillableError sets the Error field if the given value is not nil.
func (tuo *TaskUpdateOne) SetNillableError(s *string) *TaskUpdateOne {
	if s != nil {
		tuo.SetError(*s)
	}
	return tuo
}

// ClearError clears the value of Error.
func (tuo *TaskUpdateOne) ClearError() *TaskUpdateOne {
	tuo.Error = nil
	tuo.clearError = true
	return tuo
}

// SetSessionID sets the SessionID field.
func (tuo *TaskUpdateOne) SetSessionID(s string) *TaskUpdateOne {
	tuo.SessionID = &s
	return tuo
}

// SetNillableSessionID sets the SessionID field if the given value is not nil.
func (tuo *TaskUpdateOne) SetNillableSessionID(s *string) *TaskUpdateOne {
	if s != nil {
		tuo.SetSessionID(*s)
	}
	return tuo
}

// ClearSessionID clears the value of SessionID.
func (tuo *TaskUpdateOne) ClearSessionID() *TaskUpdateOne {
	tuo.SessionID = nil
	tuo.clearSessionID = true
	return tuo
}

// AddTagIDs adds the tags edge to Tag by ids.
func (tuo *TaskUpdateOne) AddTagIDs(ids ...int) *TaskUpdateOne {
	if tuo.tags == nil {
		tuo.tags = make(map[int]struct{})
	}
	for i := range ids {
		tuo.tags[ids[i]] = struct{}{}
	}
	return tuo
}

// AddTags adds the tags edges to Tag.
func (tuo *TaskUpdateOne) AddTags(t ...*Tag) *TaskUpdateOne {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return tuo.AddTagIDs(ids...)
}

// SetJobID sets the job edge to Job by id.
func (tuo *TaskUpdateOne) SetJobID(id int) *TaskUpdateOne {
	if tuo.job == nil {
		tuo.job = make(map[int]struct{})
	}
	tuo.job[id] = struct{}{}
	return tuo
}

// SetJob sets the job edge to Job.
func (tuo *TaskUpdateOne) SetJob(j *Job) *TaskUpdateOne {
	return tuo.SetJobID(j.ID)
}

// SetTargetID sets the target edge to Target by id.
func (tuo *TaskUpdateOne) SetTargetID(id int) *TaskUpdateOne {
	if tuo.target == nil {
		tuo.target = make(map[int]struct{})
	}
	tuo.target[id] = struct{}{}
	return tuo
}

// SetNillableTargetID sets the target edge to Target by id if the given value is not nil.
func (tuo *TaskUpdateOne) SetNillableTargetID(id *int) *TaskUpdateOne {
	if id != nil {
		tuo = tuo.SetTargetID(*id)
	}
	return tuo
}

// SetTarget sets the target edge to Target.
func (tuo *TaskUpdateOne) SetTarget(t *Target) *TaskUpdateOne {
	return tuo.SetTargetID(t.ID)
}

// RemoveTagIDs removes the tags edge to Tag by ids.
func (tuo *TaskUpdateOne) RemoveTagIDs(ids ...int) *TaskUpdateOne {
	if tuo.removedTags == nil {
		tuo.removedTags = make(map[int]struct{})
	}
	for i := range ids {
		tuo.removedTags[ids[i]] = struct{}{}
	}
	return tuo
}

// RemoveTags removes tags edges to Tag.
func (tuo *TaskUpdateOne) RemoveTags(t ...*Tag) *TaskUpdateOne {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return tuo.RemoveTagIDs(ids...)
}

// ClearJob clears the job edge to Job.
func (tuo *TaskUpdateOne) ClearJob() *TaskUpdateOne {
	tuo.clearedJob = true
	return tuo
}

// ClearTarget clears the target edge to Target.
func (tuo *TaskUpdateOne) ClearTarget() *TaskUpdateOne {
	tuo.clearedTarget = true
	return tuo
}

// Save executes the query and returns the updated entity.
func (tuo *TaskUpdateOne) Save(ctx context.Context) (*Task, error) {
	if tuo.Content != nil {
		if err := task.ContentValidator(*tuo.Content); err != nil {
			return nil, fmt.Errorf("ent: validator failed for field \"Content\": %v", err)
		}
	}
	if tuo.SessionID != nil {
		if err := task.SessionIDValidator(*tuo.SessionID); err != nil {
			return nil, fmt.Errorf("ent: validator failed for field \"SessionID\": %v", err)
		}
	}
	if len(tuo.job) > 1 {
		return nil, errors.New("ent: multiple assignments on a unique edge \"job\"")
	}
	if tuo.clearedJob && tuo.job == nil {
		return nil, errors.New("ent: clearing a unique edge \"job\"")
	}
	if len(tuo.target) > 1 {
		return nil, errors.New("ent: multiple assignments on a unique edge \"target\"")
	}
	return tuo.sqlSave(ctx)
}

// SaveX is like Save, but panics if an error occurs.
func (tuo *TaskUpdateOne) SaveX(ctx context.Context) *Task {
	t, err := tuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return t
}

// Exec executes the query on the entity.
func (tuo *TaskUpdateOne) Exec(ctx context.Context) error {
	_, err := tuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tuo *TaskUpdateOne) ExecX(ctx context.Context) {
	if err := tuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (tuo *TaskUpdateOne) sqlSave(ctx context.Context) (t *Task, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   task.Table,
			Columns: task.Columns,
			ID: &sqlgraph.FieldSpec{
				Value:  tuo.id,
				Type:   field.TypeInt,
				Column: task.FieldID,
			},
		},
	}
	if value := tuo.QueueTime; value != nil {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  *value,
			Column: task.FieldQueueTime,
		})
	}
	if value := tuo.LastChangedTime; value != nil {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  *value,
			Column: task.FieldLastChangedTime,
		})
	}
	if value := tuo.ClaimTime; value != nil {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  *value,
			Column: task.FieldClaimTime,
		})
	}
	if tuo.clearClaimTime {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: task.FieldClaimTime,
		})
	}
	if value := tuo.ExecStartTime; value != nil {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  *value,
			Column: task.FieldExecStartTime,
		})
	}
	if tuo.clearExecStartTime {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: task.FieldExecStartTime,
		})
	}
	if value := tuo.ExecStopTime; value != nil {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  *value,
			Column: task.FieldExecStopTime,
		})
	}
	if tuo.clearExecStopTime {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: task.FieldExecStopTime,
		})
	}
	if value := tuo.Content; value != nil {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  *value,
			Column: task.FieldContent,
		})
	}
	if value := tuo.Output; value != nil {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  *value,
			Column: task.FieldOutput,
		})
	}
	if tuo.clearOutput {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: task.FieldOutput,
		})
	}
	if value := tuo.Error; value != nil {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  *value,
			Column: task.FieldError,
		})
	}
	if tuo.clearError {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: task.FieldError,
		})
	}
	if value := tuo.SessionID; value != nil {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  *value,
			Column: task.FieldSessionID,
		})
	}
	if tuo.clearSessionID {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: task.FieldSessionID,
		})
	}
	if nodes := tuo.removedTags; len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   task.TagsTable,
			Columns: task.TagsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: tag.FieldID,
				},
			},
		}
		for k, _ := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tuo.tags; len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   task.TagsTable,
			Columns: task.TagsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: tag.FieldID,
				},
			},
		}
		for k, _ := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if tuo.clearedJob {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   task.JobTable,
			Columns: []string{task.JobColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: job.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tuo.job; len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   task.JobTable,
			Columns: []string{task.JobColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: job.FieldID,
				},
			},
		}
		for k, _ := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if tuo.clearedTarget {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   task.TargetTable,
			Columns: []string{task.TargetColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: target.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tuo.target; len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   task.TargetTable,
			Columns: []string{task.TargetColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: target.FieldID,
				},
			},
		}
		for k, _ := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	t = &Task{config: tuo.config}
	_spec.Assign = t.assignValues
	_spec.ScanValues = t.scanValues()
	if err = sqlgraph.UpdateNode(ctx, tuo.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return t, nil
}
