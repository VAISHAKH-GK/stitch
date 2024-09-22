// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.747
package user

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

import (
	"github.com/FulgurCode/stitch/view/layout"
)

func Item() templ.Component {
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
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<link rel=\"stylesheet\" href=\"/static/styles/user/item.css\">")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var3 = []any{Container()}
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
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `view/user/Item.templ`, Line: 1, Col: 0}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var4))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\">")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			templ_7745c5c3_Err = ImageCarousel().Render(ctx, templ_7745c5c3_Buffer)
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div class=\"description-container\"><span style=\"gap: 0\"><p class=\"no-line-height\">Shop name</p><h1 style=\"font-weight: 600;\" class=\"no-line-height\">Really big Clothe name</h1></span> <span style=\"gap: 2px;\"><p style=\"font-weight: 600;\" class=\"no-line-height\">&#8360;. 5000.00</p><p>Lorem ipsum dolor sit amet, consectetur adipiscing elit. Curabitur vitae  libero rhoncus tellus venenatis aliquam quis nec dolor. Nulla aliquet  tempor nisl</p>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			templ_7745c5c3_Err = templ.RenderScriptItems(ctx, templ_7745c5c3_Buffer, clothingSizeDialog())
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<p class=\"no-line-height\" onclick=\"")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var5 templ.ComponentScript = clothingSizeDialog()
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ_7745c5c3_Var5.Call)
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\">Size: </p></span><div class=\"size-group\"><div class=\"radio-button\"><input type=\"radio\" id=\"small-size\" name=\"size\" value=\"S\"> <label for=\"small-size\">S</label></div><div class=\"radio-button\"><input type=\"radio\" id=\"medium-size\" name=\"size\" value=\"M\"> <label for=\"medium-size\">M</label></div><div class=\"radio-button\"><input type=\"radio\" id=\"large-size\" name=\"size\" value=\"L\"> <label for=\"large-size\">L</label></div><div class=\"radio-button\"><input type=\"radio\" id=\"extra-large-size\" name=\"size\" value=\"XL\"> <label for=\"extra-large-size\">XL</label></div></div><span><button class=\"button-primary\">BUY</button> <button class=\"button-secondary\">Add to Cart</button></span><br></div></div><dialog id=\"clothingSizeDialog\"><table><tr><th>AREA</th><th>XS</th><th>S</th><th>M</th><th>L</th><th>XL</th></tr><tr><td>A - Chest</td><td>45.0</td><td>48.0</td><td>51.0</td><td>54.0</td><td>57.0</td></tr><tr><td>A - Front length</td><td>63.0</td><td>64.5</td><td>66.0</td><td>67.5</td><td>69.0</td></tr><tr><td>A - Sleeve length</td><td>21.0</td><td>22.0</td><td>23.0</td><td>24.0</td><td>25.0</td></tr><tr><td>A - Back width</td><td>37.5</td><td>40.0</td><td>42.5</td><td>45.0</td><td>47.5</td></tr><tr><td>A - Arm width</td><td>45.0</td><td>48.0</td><td>51.0</td><td>54.0</td><td>57.0</td></tr></table></dialog>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			return templ_7745c5c3_Err
		})
		templ_7745c5c3_Err = layout.User().Render(templ.WithChildren(ctx, templ_7745c5c3_Var2), templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}

func clothingSizeDialog() templ.ComponentScript {
	return templ.ComponentScript{
		Name: `__templ_clothingSizeDialog_7f75`,
		Function: `function __templ_clothingSizeDialog_7f75(){let dialog = document.getElementById("clothingSizeDialog");
    dialog.showModal();
}`,
		Call:       templ.SafeScript(`__templ_clothingSizeDialog_7f75`),
		CallInline: templ.SafeScriptInline(`__templ_clothingSizeDialog_7f75`),
	}
}

func ImageCarousel() templ.Component {
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
		templ_7745c5c3_Var6 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var6 == nil {
			templ_7745c5c3_Var6 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<style>\n        .carousel-container {\n            width: 100%;\n            max-width: 550px;\n            // margin: 0 auto;\n            position: relative;\n            display: flex;\n            flex-direction: column;\n            align-items: center;\n\n            margin-bottom: 4rem;\n        }\n        .main-image-container {\n            width: 100%;\n            padding-top: 133.33%; /* 3:4 aspect ratio */\n            position: relative;\n            overflow: hidden;\n        }\n        .main-image {\n            position: absolute;\n            top: 0;\n            left: 0;\n            width: 100%;\n            height: 100%;\n            object-fit: cover;\n        }\n        .preview-panel {\n            display: flex;\n            align-items: center;\n            width: 100%;\n            margin-top: 10px;\n        }\n        .preview-container {\n            display: flex;\n            justify-content: center;\n            overflow-x: auto;\n            flex-grow: 1;\n            scrollbar-width: none; /* Firefox */\n            -ms-overflow-style: none; /* IE and Edge */\n        }\n        .preview-container::-webkit-scrollbar {\n            display: none; /* Chrome, Safari, and Opera */\n        }\n        .preview-image {\n            width: 60px;\n            height: 80px;\n            object-fit: cover;\n            margin: 0 5px;\n            cursor: pointer;\n            opacity: 0.6;\n            transition: opacity 0.3s;\n            flex-shrink: 0;\n        }\n        .preview-image:hover, .preview-image.active {\n            opacity: 1;\n        }\n        .arrows {\n            background: none;\n            border: none;\n            font-size: 24px;\n            cursor: pointer;\n            padding: 0 10px;\n            color: var(--primary);\n        }\n        @media (max-width: 480px) {\n            .preview-image {\n                width: 45px;\n                height: 60px;\n            }\n        }\n    </style><div class=\"carousel-container\"><div class=\"main-image-container\"><img src=\"https://picsum.photos/300/400\" alt=\"Main Image\" class=\"main-image\" id=\"mainImage\"></div><div class=\"preview-panel\"><button class=\"arrows prev\" onclick=\"changeSlide(-1)\">&#10094;</button><div class=\"preview-container\"><img src=\"https://picsum.photos/300/400\" alt=\"Preview 1\" class=\"preview-image active\" onclick=\"setMainImage(0)\"> <img src=\"https://picsum.photos/301/400\" alt=\"Preview 2\" class=\"preview-image\" onclick=\"setMainImage(1)\"> <img src=\"https://picsum.photos/302/400\" alt=\"Preview 3\" class=\"preview-image\" onclick=\"setMainImage(2)\"> <img src=\"https://picsum.photos/303/400\" alt=\"Preview 4\" class=\"preview-image\" onclick=\"setMainImage(3)\"> <img src=\"https://picsum.photos/304/400\" alt=\"Preview 5\" class=\"preview-image\" onclick=\"setMainImage(4)\"></div><button class=\"arrows next\" onclick=\"changeSlide(1)\">&#10095;</button></div></div><script>\n        const images = [\n            \"https://picsum.photos/300/400\",\n            \"https://picsum.photos/301/400\",\n            \"https://picsum.photos/302/400\",\n            \"https://picsum.photos/303/400\",\n            \"https://picsum.photos/304/400\"\n        ];\n        let currentIndex = 0;\n\n        function setMainImage(index) {\n            currentIndex = index;\n            updateCarousel();\n        }\n\n        function changeSlide(direction) {\n            currentIndex = (currentIndex + direction + images.length) % images.length;\n            updateCarousel();\n        }\n\n        function updateCarousel() {\n            const mainImage = document.getElementById('mainImage');\n            mainImage.src = images[currentIndex];\n            const previews = document.querySelectorAll('.preview-image');\n            previews.forEach((preview, index) => {\n                preview.classList.toggle('active', index === currentIndex);\n            });\n        }\n    </script>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}
