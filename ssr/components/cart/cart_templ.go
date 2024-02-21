// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.543
package cart

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

import "api/models"
import "fmt"

func CartComponent(products *[]models.Product) templ.Component {
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
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div class=\"cart\"><button class=\"cart-button\" id=\"cart-button\"><img src=\"/public/icons/shopping-bag.svg\" alt=\"Cart\"> ")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if products != nil {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<span class=\"cart-count\">")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var2 string
			templ_7745c5c3_Var2, templ_7745c5c3_Err = templ.JoinStringErrs(fmt.Sprintf("%d Item(s) in Cart", len(*products)))
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `components/cart/cart.templ`, Line: 10, Col: 80}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var2))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</span>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		} else {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<span class=\"cart-count\">0 Items in Cart</span>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</button> ")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if products != nil && len(*products) > 0 {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div class=\"cart-container hidden\" id=\"cart-container\"><ul>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			for _, product := range *products {
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<li><div class=\"cart-item\"><span>")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				var templ_7745c5c3_Var3 string
				templ_7745c5c3_Var3, templ_7745c5c3_Err = templ.JoinStringErrs(product.Name)
				if templ_7745c5c3_Err != nil {
					return templ.Error{Err: templ_7745c5c3_Err, FileName: `components/cart/cart.templ`, Line: 21, Col: 28}
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var3))
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</span> <button class=\"icon-button\" type=\"button\"><img src=\"/public/icons/trash.svg\" alt=\"Remove\"></button></div></li>")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<button class=\"icon-button\" style=\"gap: 1rem; padding: 1rem; background: #3880ff; color: #fff\"><img src=\"/public/icons/shopping-cart.svg\" alt=\"\" style=\"filter: invert(100%)\"> <span>Checkout</span></button></ul></div>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</div><script>\n\t\tconst cartButton = document.getElementById('cart-button');\n\n\n\t\tcartButton.addEventListener('click', () => {\n\t\t\tconst cartContainer = document.getElementById('cart-container');\n\t\t\tif (!cartContainer) {\n\t\t\t\treturn;\n\t\t\t}\n\n\t\t\tconst cartOpen = cartContainer.classList.contains('hidden');\n\t\t\tif (cartOpen) {\n\t\t\t\tcartContainer.classList.remove('hidden');\n\t\t\t} else {\n\t\t\t\tcartContainer.classList.add('leave')\n\t\t\t\tsetTimeout(() => {\n\t\t\t\t\tcartContainer.classList.add('hidden');\n\t\t\t\t\tcartContainer.classList.remove('leave');\n\t\t\t\t}, 500);\n\t\t\t}\n\t\t});\n\t</script>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}