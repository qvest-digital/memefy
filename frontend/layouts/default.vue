<template>
  <div class="wrapper">
    <header class="main-header">
      <a href="" class="logo">
        <span class="logo-mini">MEMEfy</span>
        <span class="logo-lg">MEMEfy</span>
      </a>
      <nav class="navbar navbar-static-top">
        <a href="javascript:void(0);" class="sidebar-toggle" data-toggle="push-menu" role="button">
          <span class="sr-only">Toggle navigation</span>
        </a>

        <!-- Navbar Right Menu -->
        <div class="navbar-custom-menu">
          <ul class="nav navbar-nav">


          </ul>
        </div>
      </nav>
    </header>
    <aside class="main-sidebar">
      <section class="sidebar">
        <ul class="sidebar-menu tree" data-widget="tree">
          <nuxt-link :to="{name: 'index' }" tag="li" active-class="active" exact>
            <a>
              <i class="glyphicon glyphicon-film"></i>
              <span>MEME's</span>
            </a>
          </nuxt-link>
          <nuxt-link :to="{name: 'upload'}" tag="li" active-class="active">
            <a>
              <i class="glyphicon glyphicon-open"></i>
              <span>Upload</span>
            </a>
          </nuxt-link>
        </ul>
      </section>
    </aside>
    <div class="content-wrapper">
      <section class="content-header"></section>
      <section class="content container-fluid">
        <nuxt/>
      </section>
    </div>

    <footer class="main-footer">
      <div class="row">
        <div class="col-xs-12 col-lg-6">
          <strong class="well-sm">Version</strong>
        </div>
        <div class="col-xs-12 col-lg-6 text-right">
          <strong class="well-sm">© Popel</strong>
        </div>
      </div>
    </footer>
    <!-- Add the sidebar's background. This div must be placed immediately after the control sidebar -->
    <div class="control-sidebar-bg"></div>
  </div>
</template>

<style>

</style>
