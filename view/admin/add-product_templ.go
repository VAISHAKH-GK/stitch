// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.747
package admin

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

import (
	"github.com/FulgurCode/stitch/view/layout"
)

func AdminAddProduct() templ.Component {
	return templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
		templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templruntime.GetBuffer(templ_7745c5c3_W)
		if !templ_7745c5c3_IsBuffer {
			defer func() {
				templ_7745c5c3_BufErr := templruntime.ReleaseBuffer(templ_7745c5c3_Buffer)
				if templ_7745c5c3_Err == nil {
					templ_7745c5c3_Err = templ_7745c5c3_BufErr
				}
			}()
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var1 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var1 == nil {
			templ_7745c5c3_Var1 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		templ_7745c5c3_Var2 := templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
			templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
			templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templruntime.GetBuffer(templ_7745c5c3_W)
			if !templ_7745c5c3_IsBuffer {
				defer func() {
					templ_7745c5c3_BufErr := templruntime.ReleaseBuffer(templ_7745c5c3_Buffer)
					if templ_7745c5c3_Err == nil {
						templ_7745c5c3_Err = templ_7745c5c3_BufErr
					}
				}()
			}
			ctx = templ.InitializeContext(ctx)
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<link rel=\"stylesheet\" href=\"/static/styles/admin/item.css\">")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var3 = []any{Container(), innerScroll()}
			templ_7745c5c3_Err = templ.RenderCSSItems(ctx, templ_7745c5c3_Buffer, templ_7745c5c3_Var3...)
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div class=\"")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var4 string
			templ_7745c5c3_Var4, templ_7745c5c3_Err = templ.JoinStringErrs(templ.CSSClasses(templ_7745c5c3_Var3).String())
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `view/admin/add-product.templ`, Line: 1, Col: 0}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var4))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\"><h1 style=\"align-self: start;\">Add Product</h1><div class=\"item-container\"><div class=\"image\"><img src=\"https://picsum.photos/60/100\"> <img src=\"https://picsum.photos/60/110\"> <span onclick=\"document.getElementById(&#39;fileInput&#39;).click()\"><input type=\"file\" id=\"fileInput\" accept=\"image/*\" style=\"display: none\"> <span>+</span></span></div><form class=\"description-container\" hx-post=\"/admin/add-product\" hx-target=\"body\"><p class=\"no-line-height\">Title:</p><input type=\"text\" style=\"font-weight: 600; font-size: 1.5rem;\" placeholder=\"Title\" name=\"name\" required> <span style=\"gap: 2px;\"><p class=\"no-line-height\">Price:</p><p style=\"margin: 0;\"><input type=\"number\" placeholder=\"price\" name=\"price\" required> &#8360;.</p><p class=\"no-line-height\">Description:</p><textarea name=\"description\" placeholder=\"Description\" id=\"auto-expand-textarea\" required></textarea><p class=\"no-line-height\">Include Size:</p></span><div class=\"size-group\"><div class=\"checkbox-button\"><input type=\"checkbox\" id=\"small-size\" name=\"size\" value=\"S\"> <label for=\"small-size\">S</label></div><div class=\"checkbox-button\"><input type=\"checkbox\" id=\"medium-size\" name=\"size\" value=\"M\"> <label for=\"medium-size\">M</label></div><div class=\"checkbox-button\"><input type=\"checkbox\" id=\"large-size\" name=\"size\" value=\"L\"> <label for=\"large-size\">L</label></div><div class=\"checkbox-button\"><input type=\"checkbox\" id=\"extra-large-size\" name=\"size\" value=\"XL\"> <label for=\"extra-large-size\">XL</label></div></div><p class=\"no-line-height\">Category:</p><div class=\"category-group\"><div class=\"radio-button\"><input type=\"radio\" id=\"shirt\" name=\"category\" value=\"shirt\"> <label for=\"shirt\">shirt</label></div><div class=\"radio-button\"><input type=\"radio\" id=\"t-shirt\" name=\"category\" value=\"t-shirt\"> <label for=\"t-shirt\">t-shirt</label></div><div class=\"radio-button\"><input type=\"radio\" id=\"pants\" name=\"category\" value=\"pants\"> <label for=\"pants\">pants</label></div></div><input type=\"submit\" class=\"button-primary\" style=\"margin-top: var(--padding-inner)\" value=\"Add Product\"></form></div></div><script>\n            document.addEventListener('DOMContentLoaded', () => {\n            const textarea = document.getElementById('auto-expand-textarea');\n\n            const adjustHeight = () => {\n                textarea.style.height = 'auto'; // Reset height to auto to shrink if needed\n                textarea.style.height = `${textarea.scrollHeight}px`; // Set height to the scrollHeight\n            };\n\n            textarea.addEventListener('input', adjustHeight);\n\n            // Initial adjustment to handle pre-filled content\n            adjustHeight();\n        });\n        </script>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			return templ_7745c5c3_Err
		})
		templ_7745c5c3_Err = layout.Admin().Render(templ.WithChildren(ctx, templ_7745c5c3_Var2), templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}
