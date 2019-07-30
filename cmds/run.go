package cmds

import (
	"net/http"

	"github.com/spf13/cobra"
	"gopkg.in/macaron.v1"
)

func NewCmdRun() *cobra.Command {
	cmd := &cobra.Command{
		Use:               "run",
		Short:             "Launch pos",
		DisableAutoGenTag: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			m := macaron.Classic()

			m.Use(macaron.Renderer())

			m.Get("/", func(ctx *macaron.Context) {
				ctx.HTML(http.StatusOK, "home")
			})

			//m.Group("/books", func() {
			//	m.Combo("").Get(routers.GetAllBooks). // get all books
			//						Post(routers.CreateBook) // add new book
			//
			//	m.Combo("/:id").
			//		Get(routers.GetBook).      // get the specific book
			//		Put(routers.EditBook).     // update book info
			//		Delete(routers.DeleteBook) // delete book
			//})

			if err := http.ListenAndServe("0.0.0.0:8443", m); err != nil {
				panic(err)
			}
			return nil
		},
	}

	return cmd
}
