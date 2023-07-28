import React from 'react'
import { createRoot } from 'react-dom/client'

function VoteIndex() {
    return <div>Vote</div>
}

const root = createRoot(document.getElementById('app')!)
root.render(<VoteIndex />)
