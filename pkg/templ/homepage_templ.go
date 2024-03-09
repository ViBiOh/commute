// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.598
package templ

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import (
	"bytes"
	"context"
	"io"
	"strconv"
	"time"

	"github.com/a-h/templ"
)

func Login(uri, nonce, title, loginURL string) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, templ_7745c5c3_W io.Writer) (templ_7745c5c3_Err error) {
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templ_7745c5c3_W.(*bytes.Buffer)
		if !templ_7745c5c3_IsBuffer {
			templ_7745c5c3_Buffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templ_7745c5c3_Buffer)
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var1 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var1 == nil {
			templ_7745c5c3_Var1 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		templ_7745c5c3_Var2 := templ.ComponentFunc(func(ctx context.Context, templ_7745c5c3_W io.Writer) (templ_7745c5c3_Err error) {
			templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templ_7745c5c3_W.(*bytes.Buffer)
			if !templ_7745c5c3_IsBuffer {
				templ_7745c5c3_Buffer = templ.GetBuffer()
				defer templ.ReleaseBuffer(templ_7745c5c3_Buffer)
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<style type=\"text/css\" nonce=\"")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(nonce))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\">\n      #connect-strava {\n        width: 30rem;\n      }\n    </style> <a href=\"")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var3 templ.SafeURL = templ.URL(loginURL)
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(string(templ_7745c5c3_Var3)))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\"><svg id=\"connect-strava\" xmlns=\"http://www.w3.org/2000/svg\" viewBox=\"0 0 193 48\"><g fill=\"none\" fill-rule=\"evenodd\"><rect width=\"185\" height=\"40\" x=\"4\" y=\"4\" fill=\"#FC4C02\" rx=\"2\"></rect><path fill=\"#FFF\" d=\"m27 25.164 1.736.35c-.112 1.101-.513 1.993-1.204 2.674-.69.681-1.582 1.022-2.674 1.022-1.241 0-2.256-.434-3.045-1.302-.789-.868-1.183-2.165-1.183-3.892 0-1.605.413-2.844 1.239-3.717.826-.873 1.818-1.309 2.975-1.309 1.017 0 1.876.32 2.576.959.7.64 1.11 1.482 1.232 2.527l-1.708.266c-.243-1.484-.938-2.226-2.086-2.226-.719 0-1.318.29-1.799.868-.48.579-.721 1.465-.721 2.66 0 1.223.236 2.135.707 2.737.471.602 1.076.903 1.813.903 1.223 0 1.937-.84 2.142-2.52Zm6.519 2.604c.55 0 .996-.217 1.337-.651.34-.434.51-1.043.51-1.827s-.17-1.393-.51-1.827a1.62 1.62 0 0 0-1.337-.651c-.56 0-1.01.215-1.351.644-.34.43-.511 1.04-.511 1.834 0 .803.17 1.416.51 1.841.341.425.792.637 1.352.637Zm0 1.442c-.999 0-1.823-.345-2.471-1.036-.649-.69-.973-1.652-.973-2.884 0-1.213.333-2.17 1-2.87.668-.7 1.482-1.05 2.444-1.05.961 0 1.77.35 2.429 1.05.658.7.987 1.657.987 2.87 0 1.232-.32 2.193-.96 2.884-.639.69-1.458 1.036-2.456 1.036Zm5.245-.21v-7.42h1.54v.714h.027c.206-.261.49-.48.854-.658a2.589 2.589 0 0 1 1.148-.266c.822 0 1.468.245 1.94.735.47.49.706 1.169.706 2.037V29h-1.581v-4.438c0-1.148-.481-1.722-1.443-1.722-.485 0-.872.14-1.161.42-.29.28-.434.658-.434 1.134V29h-1.596Zm8.464 0v-7.42h1.54v.714h.028c.206-.261.49-.48.854-.658a2.589 2.589 0 0 1 1.148-.266c.822 0 1.468.245 1.94.735.47.49.706 1.169.706 2.037V29h-1.582v-4.438c0-1.148-.48-1.722-1.442-1.722-.485 0-.872.14-1.162.42-.29.28-.434.658-.434 1.134V29h-1.596Zm13.393-2.464 1.148.938c-.719 1.157-1.745 1.736-3.08 1.736-1.027 0-1.864-.357-2.513-1.071-.649-.714-.973-1.664-.973-2.849 0-1.185.322-2.135.966-2.849.644-.714 1.46-1.071 2.45-1.071.99 0 1.799.355 2.429 1.064.63.71.945 1.661.945 2.856v.476h-5.18c.019.616.189 1.108.511 1.477s.768.553 1.337.553c.27 0 .513-.037.728-.112.215-.075.404-.187.567-.336.163-.15.287-.28.371-.392.084-.112.182-.252.294-.42Zm-3.794-1.974h3.612c-.028-.523-.196-.95-.504-1.281-.308-.331-.747-.497-1.316-.497-.532 0-.957.177-1.274.532-.317.355-.49.77-.518 1.246Zm11.503 1.484 1.582.336c-.15.896-.49 1.591-1.022 2.086-.532.495-1.237.742-2.114.742-1.008 0-1.834-.345-2.478-1.036-.644-.69-.966-1.652-.966-2.884 0-1.185.324-2.135.973-2.849.648-.714 1.467-1.071 2.457-1.071.85 0 1.55.254 2.1.763s.872 1.169.966 1.981l-1.498.252c-.159-1.036-.677-1.554-1.554-1.554-.57 0-1.022.217-1.358.651-.336.434-.504 1.043-.504 1.827s.163 1.393.49 1.827c.326.434.784.651 1.372.651.868 0 1.386-.574 1.554-1.722Zm3.69.476v-3.57H70.9V21.58h1.162v-1.82h1.513v1.82h1.861v1.372h-1.848v3.402c0 .41.063.698.19.861.126.163.384.245.776.245h.588V29h-.713c-.897 0-1.522-.198-1.877-.595-.354-.397-.532-1.024-.532-1.883ZM81.992 29l-1.638-7.42h1.568l1.05 5.166H83l1.764-5.166h1.442l1.652 5.152h.028l1.19-5.152h1.54L88.838 29h-1.54l-1.806-5.572h-.028L83.518 29h-1.526Zm10.41 0v-7.42H94V29h-1.596Zm-.027-8.638V18.78h1.652v1.582h-1.652Zm4.32 6.16v-3.57h-1.12V21.58h1.162v-1.82h1.513v1.82h1.862v1.372h-1.849v3.402c0 .41.063.698.19.861.126.163.384.245.776.245h.588V29h-.713c-.897 0-1.522-.198-1.876-.595-.355-.397-.532-1.024-.532-1.883ZM101.87 29V18.78h1.596v3.528h.028c.168-.252.444-.471.826-.658.383-.187.77-.28 1.162-.28.794 0 1.438.243 1.932.728.495.485.742 1.148.742 1.988V29h-1.582v-4.536c0-.485-.13-.877-.392-1.176-.261-.299-.64-.448-1.134-.448-.476 0-.858.14-1.148.42-.29.28-.434.64-.434 1.078V29h-1.596ZM160.016 18.724l-2.442 4.97-2.444-4.97h-3.591L157.574 31l6.03-12.276h-3.588Zm-19.849 4.34c0-.374-.129-.653-.385-.833-.255-.181-.603-.272-1.04-.272h-1.634v2.261h1.618c.449 0 .801-.1 1.056-.297.256-.199.385-.475.385-.826v-.034ZM149.175 18l6.034 12.276h-3.592l-2.442-4.97-2.44 4.97h-6.712l-2.115-3.3h-.8v3.3h-3.747V18.724h5.477c1.004 0 1.828.119 2.474.355.647.237 1.166.56 1.561.966.342.351.598.748.77 1.187.17.44.256.959.256 1.55v.035c0 .847-.198 1.562-.594 2.145-.394.583-.933 1.046-1.617 1.386l1.947 2.93L149.175 18Zm16.792 0-6.032 12.276h3.591l2.44-4.97 2.444 4.97H172L165.967 18Zm-43.48 3.99h3.3v8.286h3.747V21.99h3.3v-3.266h-10.346v3.266Zm-.134 3.07c.229.408.344.904.344 1.486v.033c0 .605-.117 1.15-.353 1.634a3.472 3.472 0 0 1-.993 1.23c-.427.335-.946.593-1.554.775-.607.183-1.292.274-2.049.274-1.141 0-2.207-.164-3.195-.487-.987-.326-1.838-.813-2.553-1.46l2.001-2.458a5.9 5.9 0 0 0 1.921 1.038c.673.21 1.34.314 2.002.314.343 0 .587-.044.737-.131.15-.089.224-.21.224-.363v-.033c0-.167-.11-.306-.328-.414-.218-.11-.628-.225-1.226-.345a20.64 20.64 0 0 1-1.8-.464 6.1 6.1 0 0 1-1.505-.676 3.215 3.215 0 0 1-1.034-1.04c-.256-.419-.384-.93-.384-1.535v-.033c0-.551.104-1.063.312-1.536.207-.473.512-.886.912-1.237.4-.352.898-.627 1.491-.826.59-.198 1.272-.297 2.041-.297 1.088 0 2.04.132 2.858.397.816.262 1.55.659 2.202 1.187l-1.825 2.608a5.55 5.55 0 0 0-1.69-.868 5.717 5.717 0 0 0-1.673-.273c-.277 0-.483.044-.616.133a.395.395 0 0 0-.2.346v.033c0 .155.1.287.304.398.203.11.598.225 1.184.346.716.131 1.366.291 1.955.477.586.19 1.092.427 1.512.72.422.291.749.64.978 1.048Z\"></path></g></svg></a>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			if !templ_7745c5c3_IsBuffer {
				_, templ_7745c5c3_Err = io.Copy(templ_7745c5c3_W, templ_7745c5c3_Buffer)
			}
			return templ_7745c5c3_Err
		})
		templ_7745c5c3_Err = Root(uri, nonce, title).Render(templ.WithChildren(ctx, templ_7745c5c3_Var2), templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}

func Form(uri, nonce, title, token, staticMap string, places []Place) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, templ_7745c5c3_W io.Writer) (templ_7745c5c3_Err error) {
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templ_7745c5c3_W.(*bytes.Buffer)
		if !templ_7745c5c3_IsBuffer {
			templ_7745c5c3_Buffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templ_7745c5c3_Buffer)
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var4 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var4 == nil {
			templ_7745c5c3_Var4 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		templ_7745c5c3_Var5 := templ.ComponentFunc(func(ctx context.Context, templ_7745c5c3_W io.Writer) (templ_7745c5c3_Err error) {
			templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templ_7745c5c3_W.(*bytes.Buffer)
			if !templ_7745c5c3_IsBuffer {
				templ_7745c5c3_Buffer = templ.GetBuffer()
				defer templ.ReleaseBuffer(templ_7745c5c3_Buffer)
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<form method=\"POST\" action=\"/compute\"><input type=\"hidden\" name=\"token\" value=\"")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(token))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\"> <img src=\"")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(staticMap))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" class=\"max-full\" alt=\"Static image of found cluters\"><p><label for=\"home\" class=\"padding\">Home</label> <select id=\"home\" class=\"max-full\" name=\"home\">")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			for index, place := range places {
				if index == 0 {
					_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<option selected value=\"")
					if templ_7745c5c3_Err != nil {
						return templ_7745c5c3_Err
					}
					_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(place.Coordinates.String()))
					if templ_7745c5c3_Err != nil {
						return templ_7745c5c3_Err
					}
					_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\">")
					if templ_7745c5c3_Err != nil {
						return templ_7745c5c3_Err
					}
					var templ_7745c5c3_Var6 string
					templ_7745c5c3_Var6, templ_7745c5c3_Err = templ.JoinStringErrs(place.Name)
					if templ_7745c5c3_Err != nil {
						return templ.Error{Err: templ_7745c5c3_Err, FileName: `pkg/templ/homepage.templ`, Line: 32, Col: 80}
					}
					_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var6))
					if templ_7745c5c3_Err != nil {
						return templ_7745c5c3_Err
					}
					_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</option>")
					if templ_7745c5c3_Err != nil {
						return templ_7745c5c3_Err
					}
				} else {
					_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<option value=\"")
					if templ_7745c5c3_Err != nil {
						return templ_7745c5c3_Err
					}
					_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(place.Coordinates.String()))
					if templ_7745c5c3_Err != nil {
						return templ_7745c5c3_Err
					}
					_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\">")
					if templ_7745c5c3_Err != nil {
						return templ_7745c5c3_Err
					}
					var templ_7745c5c3_Var7 string
					templ_7745c5c3_Var7, templ_7745c5c3_Err = templ.JoinStringErrs(place.Name)
					if templ_7745c5c3_Err != nil {
						return templ.Error{Err: templ_7745c5c3_Err, FileName: `pkg/templ/homepage.templ`, Line: 34, Col: 71}
					}
					_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var7))
					if templ_7745c5c3_Err != nil {
						return templ_7745c5c3_Err
					}
					_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</option>")
					if templ_7745c5c3_Err != nil {
						return templ_7745c5c3_Err
					}
				}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</select></p><p><label for=\"work\" class=\"padding\">Work</label> <select id=\"work\" class=\"max-full\" name=\"work\">")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			for index, place := range places {
				if index == 1 {
					_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<option selected value=\"")
					if templ_7745c5c3_Err != nil {
						return templ_7745c5c3_Err
					}
					_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(place.Coordinates.String()))
					if templ_7745c5c3_Err != nil {
						return templ_7745c5c3_Err
					}
					_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\">")
					if templ_7745c5c3_Err != nil {
						return templ_7745c5c3_Err
					}
					var templ_7745c5c3_Var8 string
					templ_7745c5c3_Var8, templ_7745c5c3_Err = templ.JoinStringErrs(place.Name)
					if templ_7745c5c3_Err != nil {
						return templ.Error{Err: templ_7745c5c3_Err, FileName: `pkg/templ/homepage.templ`, Line: 44, Col: 80}
					}
					_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var8))
					if templ_7745c5c3_Err != nil {
						return templ_7745c5c3_Err
					}
					_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</option>")
					if templ_7745c5c3_Err != nil {
						return templ_7745c5c3_Err
					}
				} else {
					_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<option value=\"")
					if templ_7745c5c3_Err != nil {
						return templ_7745c5c3_Err
					}
					_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(place.Coordinates.String()))
					if templ_7745c5c3_Err != nil {
						return templ_7745c5c3_Err
					}
					_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\">")
					if templ_7745c5c3_Err != nil {
						return templ_7745c5c3_Err
					}
					var templ_7745c5c3_Var9 string
					templ_7745c5c3_Var9, templ_7745c5c3_Err = templ.JoinStringErrs(place.Name)
					if templ_7745c5c3_Err != nil {
						return templ.Error{Err: templ_7745c5c3_Err, FileName: `pkg/templ/homepage.templ`, Line: 46, Col: 71}
					}
					_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var9))
					if templ_7745c5c3_Err != nil {
						return templ_7745c5c3_Err
					}
					_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</option>")
					if templ_7745c5c3_Err != nil {
						return templ_7745c5c3_Err
					}
				}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</select></p><p><label for=\"month\" class=\"padding\">Month</label> <select id=\"month\" name=\"month\">")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			for i := 1; i <= 12; i++ {
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<option value=\"")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(strconv.Itoa(i)))
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\"")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				if int(time.Now().Month()) == i {
					_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(" selected")
					if templ_7745c5c3_Err != nil {
						return templ_7745c5c3_Err
					}
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(">")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				var templ_7745c5c3_Var10 string
				templ_7745c5c3_Var10, templ_7745c5c3_Err = templ.JoinStringErrs(time.Month(i).String())
				if templ_7745c5c3_Err != nil {
					return templ.Error{Err: templ_7745c5c3_Err, FileName: `pkg/templ/homepage.templ`, Line: 60, Col: 37}
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var10))
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</option>")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</select></p><p><input type=\"submit\" value=\"Compute\"></p></form>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			if !templ_7745c5c3_IsBuffer {
				_, templ_7745c5c3_Err = io.Copy(templ_7745c5c3_W, templ_7745c5c3_Buffer)
			}
			return templ_7745c5c3_Err
		})
		templ_7745c5c3_Err = Root(uri, nonce, title).Render(templ.WithChildren(ctx, templ_7745c5c3_Var5), templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}
