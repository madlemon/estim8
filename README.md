# estim8

estim8 is a small web app made in Go that helps teams estimate tasks/projects together.
It lets you create rooms where team members can join and give their estimates. You can decide between estimating in story points or real time.
It's handy for teams working on projects and trying to figure out how long tasks or projects might take.

## Run locally
To run the application, you need to have [Air](https://github.com/cosmtrek/air) installed. 
It will reload the app when relevant code changes.

Run the application using Air:
```shell
air
```

By default, the application runs on port 1234. You can access it via http://localhost:1234.

## Main dependencies
- [Echo](https://github.com/labstack/echo): Web framework for Go.
- [htmx](https://htmx.org/): JavaScript library for creating dynamic HTML content.
- [templ](https://github.com/a-h/templ): Custom HTML templates for UI rendering.
- [Tailwind CSS](https://tailwindcss.com/): Utility-first CSS framework for styling.
- [Air](https://github.com/cosmtrek/air): Live reload tool for Go applications.

## License

This project is licensed under the GLWTS License - see the LICENSE file for details.
