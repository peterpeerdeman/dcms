{{template "header.tpl" .}}

<div class="container">

    <h3 class="muted">dCMS</h3>

    <div class="navbar">
        <div class="navbar-inner">
            <ul class="nav nav-pills">
                <li><a href="#/">Dashboard</a></li>
                <li><a href="#/document/overview">Documents</a></li>
                <li><a href="#/document-type/overview">Document types</a></li>
                <li><a href="#/channel/overview">Channels</a></li>
                <li><a href="#/template/overview">Templates</a></li>
                <li><a href="#/pages/overview">Pages</a></li>
                <li><a href="#/sitemap/overview">Sitemap</a></li>
            </ul>
        </div>
    </div>

    <div class="row-fluid" ng-view></div>

</div>

{{template "footer.tpl" .}}