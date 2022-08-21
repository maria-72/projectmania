import { Accordion, AccordionSummary, Typography } from '@mui/material'
import ExpandMoreIcon from '@mui/icons-material/ExpandMore'
import React from 'react'
import { Project } from '../models'

interface ListProps {
    projectList: Project[]
}

export const ProjectList: React.FC<ListProps> = ({ projectList }) => {
    return (
        <div>
            {projectList.map(({ name }) => (
                <Accordion>
                    <AccordionSummary
                        expandIcon={<ExpandMoreIcon />}
                        aria-controls="panel1a-content"
                        id="panel1a-header"
                    >
                        <Typography>{name}</Typography>
                    </AccordionSummary>
                </Accordion>
            ))}
        </div>
    )
}
