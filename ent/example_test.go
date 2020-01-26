// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"log"
	"time"

	"github.com/facebookincubator/ent/dialect/sql"

	"github.com/kcarretto/paragon/ent/credential"
	"github.com/kcarretto/paragon/ent/event"
)

// dsn for the database. In order to run the tests locally, run the following command:
//
//	 ENT_INTEGRATION_ENDPOINT="root:pass@tcp(localhost:3306)/test?parseTime=True" go test -v
//
var dsn string

func ExampleCredential() {
	if dsn == "" {
		return
	}
	ctx := context.Background()
	drv, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("failed creating database client: %v", err)
	}
	defer drv.Close()
	client := NewClient(Driver(drv))
	// creating vertices for the credential's edges.

	// create credential vertex with its edges.
	c := client.Credential.
		Create().
		SetPrincipal("string").
		SetSecret("string").
		SetKind(credential.KindPassword).
		SetFails(1).
		SaveX(ctx)
	log.Println("credential created:", c)

	// query edges.

	// Output:
}
func ExampleEvent() {
	if dsn == "" {
		return
	}
	ctx := context.Background()
	drv, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("failed creating database client: %v", err)
	}
	defer drv.Close()
	client := NewClient(Driver(drv))
	// creating vertices for the event's edges.
	j0 := client.Job.
		Create().
		SetName("string").
		SetCreationTime(time.Now()).
		SetContent("string").
		SaveX(ctx)
	log.Println("job created:", j0)
	f1 := client.File.
		Create().
		SetName("string").
		SetCreationTime(time.Now()).
		SetLastModifiedTime(time.Now()).
		SetSize(1).
		SetContent(nil).
		SetHash("string").
		SetContentType("string").
		SaveX(ctx)
	log.Println("file created:", f1)
	c2 := client.Credential.
		Create().
		SetPrincipal("string").
		SetSecret("string").
		SetKind(credential.KindPassword).
		SetFails(1).
		SaveX(ctx)
	log.Println("credential created:", c2)
	l3 := client.Link.
		Create().
		SetAlias("string").
		SetExpirationTime(time.Now()).
		SetClicks(1).
		SaveX(ctx)
	log.Println("link created:", l3)
	t4 := client.Tag.
		Create().
		SetName("string").
		SaveX(ctx)
	log.Println("tag created:", t4)
	t5 := client.Target.
		Create().
		SetName("string").
		SetPrimaryIP("string").
		SetMachineUUID("string").
		SetPublicIP("string").
		SetPrimaryMAC("string").
		SetHostname("string").
		SetLastSeen(time.Now()).
		SaveX(ctx)
	log.Println("target created:", t5)
	t6 := client.Task.
		Create().
		SetQueueTime(time.Now()).
		SetLastChangedTime(time.Now()).
		SetClaimTime(time.Now()).
		SetExecStartTime(time.Now()).
		SetExecStopTime(time.Now()).
		SetContent("string").
		SetOutput("string").
		SetError("string").
		SetSessionID("string").
		SaveX(ctx)
	log.Println("task created:", t6)
	u7 := client.User.
		Create().
		SetName("string").
		SetOAuthID("string").
		SetPhotoURL("string").
		SetSessionToken("string").
		SetActivated(true).
		SetIsAdmin(true).
		SaveX(ctx)
	log.Println("user created:", u7)
	e8 := client.Event.
		Create().
		SetCreationTime(time.Now()).
		SetKind(event.KindCREATEJOB).
		SaveX(ctx)
	log.Println("event created:", e8)
	u9 := client.User.
		Create().
		SetName("string").
		SetOAuthID("string").
		SetPhotoURL("string").
		SetSessionToken("string").
		SetActivated(true).
		SetIsAdmin(true).
		SaveX(ctx)
	log.Println("user created:", u9)

	// create event vertex with its edges.
	e := client.Event.
		Create().
		SetCreationTime(time.Now()).
		SetKind(event.KindCREATEJOB).
		SetJob(j0).
		SetFile(f1).
		SetCredential(c2).
		SetLink(l3).
		SetTag(t4).
		SetTarget(t5).
		SetTask(t6).
		SetUser(u7).
		SetEvent(e8).
		AddLikers(u9).
		SaveX(ctx)
	log.Println("event created:", e)

	// query edges.
	j0, err = e.QueryJob().First(ctx)
	if err != nil {
		log.Fatalf("failed querying job: %v", err)
	}
	log.Println("job found:", j0)

	f1, err = e.QueryFile().First(ctx)
	if err != nil {
		log.Fatalf("failed querying file: %v", err)
	}
	log.Println("file found:", f1)

	c2, err = e.QueryCredential().First(ctx)
	if err != nil {
		log.Fatalf("failed querying credential: %v", err)
	}
	log.Println("credential found:", c2)

	l3, err = e.QueryLink().First(ctx)
	if err != nil {
		log.Fatalf("failed querying link: %v", err)
	}
	log.Println("link found:", l3)

	t4, err = e.QueryTag().First(ctx)
	if err != nil {
		log.Fatalf("failed querying tag: %v", err)
	}
	log.Println("tag found:", t4)

	t5, err = e.QueryTarget().First(ctx)
	if err != nil {
		log.Fatalf("failed querying target: %v", err)
	}
	log.Println("target found:", t5)

	t6, err = e.QueryTask().First(ctx)
	if err != nil {
		log.Fatalf("failed querying task: %v", err)
	}
	log.Println("task found:", t6)

	u7, err = e.QueryUser().First(ctx)
	if err != nil {
		log.Fatalf("failed querying user: %v", err)
	}
	log.Println("user found:", u7)

	e8, err = e.QueryEvent().First(ctx)
	if err != nil {
		log.Fatalf("failed querying event: %v", err)
	}
	log.Println("event found:", e8)

	u9, err = e.QueryLikers().First(ctx)
	if err != nil {
		log.Fatalf("failed querying likers: %v", err)
	}
	log.Println("likers found:", u9)

	// Output:
}
func ExampleFile() {
	if dsn == "" {
		return
	}
	ctx := context.Background()
	drv, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("failed creating database client: %v", err)
	}
	defer drv.Close()
	client := NewClient(Driver(drv))
	// creating vertices for the file's edges.
	l0 := client.Link.
		Create().
		SetAlias("string").
		SetExpirationTime(time.Now()).
		SetClicks(1).
		SaveX(ctx)
	log.Println("link created:", l0)

	// create file vertex with its edges.
	f := client.File.
		Create().
		SetName("string").
		SetCreationTime(time.Now()).
		SetLastModifiedTime(time.Now()).
		SetSize(1).
		SetContent(nil).
		SetHash("string").
		SetContentType("string").
		AddLinks(l0).
		SaveX(ctx)
	log.Println("file created:", f)

	// query edges.
	l0, err = f.QueryLinks().First(ctx)
	if err != nil {
		log.Fatalf("failed querying links: %v", err)
	}
	log.Println("links found:", l0)

	// Output:
}
func ExampleJob() {
	if dsn == "" {
		return
	}
	ctx := context.Background()
	drv, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("failed creating database client: %v", err)
	}
	defer drv.Close()
	client := NewClient(Driver(drv))
	// creating vertices for the job's edges.
	t0 := client.Task.
		Create().
		SetQueueTime(time.Now()).
		SetLastChangedTime(time.Now()).
		SetClaimTime(time.Now()).
		SetExecStartTime(time.Now()).
		SetExecStopTime(time.Now()).
		SetContent("string").
		SetOutput("string").
		SetError("string").
		SetSessionID("string").
		SaveX(ctx)
	log.Println("task created:", t0)
	t1 := client.Tag.
		Create().
		SetName("string").
		SaveX(ctx)
	log.Println("tag created:", t1)
	j3 := client.Job.
		Create().
		SetName("string").
		SetCreationTime(time.Now()).
		SetContent("string").
		SaveX(ctx)
	log.Println("job created:", j3)

	// create job vertex with its edges.
	j := client.Job.
		Create().
		SetName("string").
		SetCreationTime(time.Now()).
		SetContent("string").
		AddTasks(t0).
		AddTags(t1).
		SetNext(j3).
		SaveX(ctx)
	log.Println("job created:", j)

	// query edges.
	t0, err = j.QueryTasks().First(ctx)
	if err != nil {
		log.Fatalf("failed querying tasks: %v", err)
	}
	log.Println("tasks found:", t0)

	t1, err = j.QueryTags().First(ctx)
	if err != nil {
		log.Fatalf("failed querying tags: %v", err)
	}
	log.Println("tags found:", t1)

	j3, err = j.QueryNext().First(ctx)
	if err != nil {
		log.Fatalf("failed querying next: %v", err)
	}
	log.Println("next found:", j3)

	// Output:
}
func ExampleLink() {
	if dsn == "" {
		return
	}
	ctx := context.Background()
	drv, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("failed creating database client: %v", err)
	}
	defer drv.Close()
	client := NewClient(Driver(drv))
	// creating vertices for the link's edges.

	// create link vertex with its edges.
	l := client.Link.
		Create().
		SetAlias("string").
		SetExpirationTime(time.Now()).
		SetClicks(1).
		SaveX(ctx)
	log.Println("link created:", l)

	// query edges.

	// Output:
}
func ExampleTag() {
	if dsn == "" {
		return
	}
	ctx := context.Background()
	drv, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("failed creating database client: %v", err)
	}
	defer drv.Close()
	client := NewClient(Driver(drv))
	// creating vertices for the tag's edges.

	// create tag vertex with its edges.
	t := client.Tag.
		Create().
		SetName("string").
		SaveX(ctx)
	log.Println("tag created:", t)

	// query edges.

	// Output:
}
func ExampleTarget() {
	if dsn == "" {
		return
	}
	ctx := context.Background()
	drv, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("failed creating database client: %v", err)
	}
	defer drv.Close()
	client := NewClient(Driver(drv))
	// creating vertices for the target's edges.
	t0 := client.Task.
		Create().
		SetQueueTime(time.Now()).
		SetLastChangedTime(time.Now()).
		SetClaimTime(time.Now()).
		SetExecStartTime(time.Now()).
		SetExecStopTime(time.Now()).
		SetContent("string").
		SetOutput("string").
		SetError("string").
		SetSessionID("string").
		SaveX(ctx)
	log.Println("task created:", t0)
	t1 := client.Tag.
		Create().
		SetName("string").
		SaveX(ctx)
	log.Println("tag created:", t1)
	c2 := client.Credential.
		Create().
		SetPrincipal("string").
		SetSecret("string").
		SetKind(credential.KindPassword).
		SetFails(1).
		SaveX(ctx)
	log.Println("credential created:", c2)

	// create target vertex with its edges.
	t := client.Target.
		Create().
		SetName("string").
		SetPrimaryIP("string").
		SetMachineUUID("string").
		SetPublicIP("string").
		SetPrimaryMAC("string").
		SetHostname("string").
		SetLastSeen(time.Now()).
		AddTasks(t0).
		AddTags(t1).
		AddCredentials(c2).
		SaveX(ctx)
	log.Println("target created:", t)

	// query edges.
	t0, err = t.QueryTasks().First(ctx)
	if err != nil {
		log.Fatalf("failed querying tasks: %v", err)
	}
	log.Println("tasks found:", t0)

	t1, err = t.QueryTags().First(ctx)
	if err != nil {
		log.Fatalf("failed querying tags: %v", err)
	}
	log.Println("tags found:", t1)

	c2, err = t.QueryCredentials().First(ctx)
	if err != nil {
		log.Fatalf("failed querying credentials: %v", err)
	}
	log.Println("credentials found:", c2)

	// Output:
}
func ExampleTask() {
	if dsn == "" {
		return
	}
	ctx := context.Background()
	drv, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("failed creating database client: %v", err)
	}
	defer drv.Close()
	client := NewClient(Driver(drv))
	// creating vertices for the task's edges.
	t0 := client.Tag.
		Create().
		SetName("string").
		SaveX(ctx)
	log.Println("tag created:", t0)

	// create task vertex with its edges.
	t := client.Task.
		Create().
		SetQueueTime(time.Now()).
		SetLastChangedTime(time.Now()).
		SetClaimTime(time.Now()).
		SetExecStartTime(time.Now()).
		SetExecStopTime(time.Now()).
		SetContent("string").
		SetOutput("string").
		SetError("string").
		SetSessionID("string").
		AddTags(t0).
		SaveX(ctx)
	log.Println("task created:", t)

	// query edges.
	t0, err = t.QueryTags().First(ctx)
	if err != nil {
		log.Fatalf("failed querying tags: %v", err)
	}
	log.Println("tags found:", t0)

	// Output:
}
func ExampleUser() {
	if dsn == "" {
		return
	}
	ctx := context.Background()
	drv, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("failed creating database client: %v", err)
	}
	defer drv.Close()
	client := NewClient(Driver(drv))
	// creating vertices for the user's edges.
	j0 := client.Job.
		Create().
		SetName("string").
		SetCreationTime(time.Now()).
		SetContent("string").
		SaveX(ctx)
	log.Println("job created:", j0)
	e1 := client.Event.
		Create().
		SetCreationTime(time.Now()).
		SetKind(event.KindCREATEJOB).
		SaveX(ctx)
	log.Println("event created:", e1)

	// create user vertex with its edges.
	u := client.User.
		Create().
		SetName("string").
		SetOAuthID("string").
		SetPhotoURL("string").
		SetSessionToken("string").
		SetActivated(true).
		SetIsAdmin(true).
		AddJobs(j0).
		AddEvents(e1).
		SaveX(ctx)
	log.Println("user created:", u)

	// query edges.
	j0, err = u.QueryJobs().First(ctx)
	if err != nil {
		log.Fatalf("failed querying jobs: %v", err)
	}
	log.Println("jobs found:", j0)

	e1, err = u.QueryEvents().First(ctx)
	if err != nil {
		log.Fatalf("failed querying events: %v", err)
	}
	log.Println("events found:", e1)

	// Output:
}
