package admin

import (
	"geoserver/internal/db/models"
	"html/template"
	"net/http"
	"strconv"

	"gorm.io/gorm"
)

type AdminPanel struct {
	db *gorm.DB
}

func Init(db *gorm.DB) *http.ServeMux {
	a := AdminPanel{db: db}
	mux := http.NewServeMux()
	a.RegisterHandlers(mux)
	return mux
}

func (a *AdminPanel) RegisterHandlers(mux *http.ServeMux) {
	mux.HandleFunc("/admin/layers", a.handleLayers)
	mux.HandleFunc("/admin/layers/edit/", a.handleEditLayer)
	//mux.HandleFunc("/admin/layers/delete/", a.handleDeleteLayer)
	mux.HandleFunc("/admin/styles", a.handleStyles)
	mux.HandleFunc("/admin/users", a.handleUsers)
}

// Обработчики для Layer
func (a *AdminPanel) handleLayers(w http.ResponseWriter, r *http.Request) {
	var layers []models.Layer
	a.db.Find(&layers)

	tmpl := template.Must(template.New("layers").Parse(`
	<!DOCTYPE html>
	<html>
	<head>
		<title>Layers Admin</title>
		<link href="https://cdn.jsdelivr.net/npm/bootstrap@5.1.3/dist/css/bootstrap.min.css" rel="stylesheet">
	</head>
	<body>
		<div class="container mt-4">
			<h1>Layers</h1>
			<table class="table">
				<thead>
					<tr>
						<th>ID</th>
						<th>Name</th>
						<th>Title</th>
						<th>Source</th>
						<th>Actions</th>
					</tr>
				</thead>
				<tbody>
					{{range .}}
					<tr>
						<td>{{.ID}}</td>
						<td>{{.Name}}</td>
						<td>{{.Title}}</td>
						<td>{{.SourceType}}: {{.SourcePath}}</td>
						<td>
							<a href="/admin/layers/edit/{{.ID}}" class="btn btn-sm btn-primary">Edit</a>
							<a href="/admin/layers/delete/{{.ID}}" class="btn btn-sm btn-danger">Delete</a>
						</td>
					</tr>
					{{end}}
				</tbody>
			</table>
		</div>
	</body>
	</html>
	`))
	tmpl.Execute(w, layers)
}

func (a *AdminPanel) handleEditLayer(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Path[len("/admin/layers/edit/"):]
	id, _ := strconv.Atoi(idStr)

	var layer models.Layer
	if r.Method == http.MethodPost {
		a.db.First(&layer, id)
		// Обновляем поля из формы
		layer.Name = r.FormValue("name")
		layer.Title = r.FormValue("title")
		a.db.Save(&layer)
		http.Redirect(w, r, "/admin/layers", http.StatusSeeOther)
		return
	}

	a.db.First(&layer, id)
	tmpl := template.Must(template.New("editLayer").Parse(`
	<!DOCTYPE html>
	<html>
	<head>
		<title>Edit Layer</title>
		<link href="https://cdn.jsdelivr.net/npm/bootstrap@5.1.3/dist/css/bootstrap.min.css" rel="stylesheet">
	</head>
	<body>
		<div class="container mt-4">
			<h1>Edit Layer</h1>
			<form method="POST">
				<div class="mb-3">
					<label class="form-label">Name</label>
					<input type="text" class="form-control" name="name" value="{{.Name}}">
				</div>
				<div class="mb-3">
					<label class="form-label">Title</label>
					<input type="text" class="form-control" name="title" value="{{.Title}}">
				</div>
				<button type="submit" class="btn btn-primary">Save</button>
			</form>
		</div>
	</body>
	</html>
	`))
	tmpl.Execute(w, layer)
}

// Аналогичные обработчики для Style и User...

func (a *AdminPanel) handleStyles(w http.ResponseWriter, r *http.Request) {
	var styles []models.Style
	a.db.Preload("Layer").Find(&styles)

	tmpl := template.Must(template.New("styles").Parse(`
	<!DOCTYPE html>
	<html>
	<head>
		<title>Styles Admin</title>
		<link href="https://cdn.jsdelivr.net/npm/bootstrap@5.1.3/dist/css/bootstrap.min.css" rel="stylesheet">
	</head>
	<body>
		<div class="container mt-4">
			<h1>Styles</h1>
			<table class="table">
				<thead>
					<tr>
						<th>ID</th>
						<th>Name</th>
						<th>Layer</th>
						<th>Actions</th>
					</tr>
				</thead>
				<tbody>
					{{range .}}
					<tr>
						<td>{{.ID}}</td>
						<td>{{.Name}}</td>
						<td>{{.Layer.Title}}</td>
						<td>
							<a href="/admin/styles/edit/{{.ID}}" class="btn btn-sm btn-primary">Edit</a>
							<a href="/admin/styles/delete/{{.ID}}" class="btn btn-sm btn-danger">Delete</a>
						</td>
					</tr>
					{{end}}
				</tbody>
			</table>
		</div>
	</body>
	</html>
	`))
	tmpl.Execute(w, styles)
}

func (a *AdminPanel) handleUsers(w http.ResponseWriter, r *http.Request) {
	var users []models.User
	a.db.Find(&users)

	tmpl := template.Must(template.New("users").Parse(`
	<!DOCTYPE html>
	<html>
	<head>
		<title>Users Admin</title>
		<link href="https://cdn.jsdelivr.net/npm/bootstrap@5.1.3/dist/css/bootstrap.min.css" rel="stylesheet">
	</head>
	<body>
		<div class="container mt-4">
			<h1>Users</h1>
			<table class="table">
				<thead>
					<tr>
						<th>ID</th>
						<th>Login</th>
						<th>Actions</th>
					</tr>
				</thead>
				<tbody>
					{{range .}}
					<tr>
						<td>{{.ID}}</td>
						<td>{{.Login}}</td>
						<td>
							<a href="/admin/users/edit/{{.ID}}" class="btn btn-sm btn-primary">Edit</a>
							<a href="/admin/users/delete/{{.ID}}" class="btn btn-sm btn-danger">Delete</a>
						</td>
					</tr>
					{{end}}
				</tbody>
			</table>
		</div>
	</body>
	</html>
	`))
	tmpl.Execute(w, users)
}
