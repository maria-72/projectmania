import { Box, CircularProgress } from '@mui/material'
import React from 'react'
import { Project } from '../../../models/project'
import { ProjectList } from './list'

interface ContentProps {
    isLoading: boolean
    projectList?: Project[]
}

export const Content: React.FC<ContentProps> = ({ isLoading, projectList }) => {
    if (isLoading) {
        return (
            <Box display="flex" justifyContent="center">
                <CircularProgress color="primary" />
            </Box>
        )
    }

    if (!projectList) {
        return null
    }

    return <ProjectList projectList={projectList} />
}
