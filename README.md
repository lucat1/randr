<p align="center">
  <b>Randr</b>
</p>

Randr is a `golang` library to render HTML templates for `server-side rendering`, highly inspired by [`react`](https://github.com/facebook/react), [`preact/htm`](https://github.com/developit/htm) and [`lit-element`](https://github.com/polymer/lit-html). It is still a proof-of-concept, but should perform way better compared to `html/template` because it compiles expressions to static code in golang, which means it also has *0 runtime overhead*(similarly to [`quicktemplate`](https://github.com/valyala/quicktemplate)).