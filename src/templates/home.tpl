{{template "header.tpl" .}}

<div class="container">

    <ul class="nav nav-pills">
        <li><a href="#/">Dashboard</a></li>
        <li><a href="#/document/overview">Documents</a></li>
        <li><a href="#/template/overview">Templates</a></li>
        <li><a href="#/pages/overview">Pages</a></li>
        <li><a href="#/sitemap/overview">Sitemap</a></li>
    </ul>

    <div ng-view></div>

</div>

{{template "footer.tpl" .}}