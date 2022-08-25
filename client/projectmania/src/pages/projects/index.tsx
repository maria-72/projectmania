import React from 'react'
import { Header } from './components/header'
import { projectsService } from '../../services/projects.service'
import { useFetch } from '../../hooks/use-fetch'
import { Content } from './components/content'

export const ProjectsPage: React.FC = () => {
    const { isLoading, response: projectList } = useFetch(
        projectsService.getProjects
    )

    return (
        <>
            <Header />
            <Content isLoading={isLoading} projectList={projectList} />
        </>
    )
}
