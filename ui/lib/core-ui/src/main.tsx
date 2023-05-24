import React from 'react'
import ReactDOM from 'react-dom/client'


import { PageLayout } from '.'
import {Test} from "./test.tsx";

ReactDOM.createRoot(document.getElementById('root') as HTMLElement).render(
  <React.StrictMode>
    <PageLayout>
        <Test/>
    </PageLayout>
  </React.StrictMode>,
)
