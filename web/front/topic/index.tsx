import React from 'react'
import { createRoot } from 'react-dom/client'

function TopicIndex() {
    return <div>Topic</div>
}

const root = createRoot(document.getElementById('app')!)
root.render(<TopicIndex />)
