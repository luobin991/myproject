{{define "navbar"}}
<!-- 
<div class="collapse navbar-collapse container">
	<a href="/" class="navbar-brand">我的博客</a>
	<div class="">
		<ul class="nav navbar-nav">
			<li {{if .IsHone}}class="active"{{end}}><a href="/">首页</a></li>
			<li {{if .IsCategory}}class="active"{{end}}><a href="/actegory">分类</a></li>
			<li {{if .IsTopic}}class="active"{{end}}><a href="/topic">文章</a></li>
		</ul>
	</div>

	<div class="pull-right">
		<ul class="nav navbar-nav">
			{{ if .IsLogin}}
			<li><a href="/login?exit=true"></a></li>
			{{else}}
			<li><a href="/login">管理员登录</a></li>
			{{end}}
		</ul>
	</div>
</div>
 -->

<nav class="navbar navbar-default navbar-fixed-top" role="navigation">
  <div class="container">
    <div class="navbar-header">
      <button type="button" class="navbar-toggle" data-toggle="collapse" data-target="#navbar-menu">
        <!-- <span class="sr-only">切换导航</span> -->
        <span class="icon-bar"></span>
        <span class="icon-bar"></span>
      </button>
      <a href="#" class="navbar-brand nav-title">Beego</a>
    </div>
    <div class="collapse navbar-collapse" id="navbar-menu">
      <ul class="nav navbar-nav navbar-left">
        <li class="cative"><a href="#">Index</a></li>
        <li class="cative"><a href="#">About</a></li>
        <li class="cative"><a href="#">Portfolio</a></li>
        <li class="cative"><a href="#">Contact</a></li>
      </ul>
      <ul class="nav navbar-nav navbar-right">
          {{if .IsLogin}}
          <li> <a href="/login?exit=0">admin | exit</a></li>
          {{else}}
          <li><a href="/login">Login</a></li>
          {{end}}

          {{range .IsLogins}}{{.}}{{end}}
      </ul>
    </div>
  </div>
</nav> 
{{end}}

