import React from 'react'
import { createRoot } from 'react-dom/client'

import '../../style.css'

function TopicIndex() {
    return <div>Topic</div>
}

const root = createRoot(document.getElementById('app')!)
root.render(<TopicIndex />)
