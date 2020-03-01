{{define "navbar"}}
        <a class="navbar-brand" href="/">PanBlog</a>
        <ul class="nav nav-tabs">
            <li role="presentation" {{if .IsHome}}class="active"{{end}}><a href="/">首页</a></li>
            <li role="presentation" {{if .IsCategory}}class="active"{{end}}><a href="/category">分类</a></li>
            <li role="presentation" {{if .IsTopic}}class="active"{{end}}><a href="/topic">文章</a></li>
            <div class="pull-right">
                  <ul class="nav navbar-nav">
                      {{if.IsLogin}}
                      <li><a href="/login?exit=true">退出</a> </li>
                      {{else}}
                      <li><a href="/login">管理员登录</a></li>
                      {{end}}

                  </ul>
            </div>
        </ul>
{{end}}