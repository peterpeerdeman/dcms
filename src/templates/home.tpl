{{template "header.tpl" .}}

<div class="container">

    <ul class="header top menu">
        <a href="#/">Dashboard</a>
        <a href="#/document/overview">Documents</a>
        <a href="#/template/overview">Templates</a>
        <a href="#/pages/overview">Pages</a>
        <a href="#/sitemap/overview">Sitemap</a>
    </ul>

    <div ng-view></div>

</div>

{{template "footer.tpl" .}}