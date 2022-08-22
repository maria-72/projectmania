import React from 'react'
import { Header } from './components/header'
import { ProjectList } from './components/list'
import { Project } from './models'

const projectList: Project[] = [
    {
        name: 'Projectmania',
    },
]

export const ProjectsPage: React.FC = () => {
    return (
        <>
            <Header />
            <ProjectList projectList={projectList} />
        </>
    )
}
