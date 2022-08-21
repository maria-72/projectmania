import { Box, Typography } from '@mui/material'
import React from 'react'

export const Header: React.FC = () => {
    return (
        <Box margin="40px 0">
            <Typography
                variant="h2"
                gutterBottom
                display="flex"
                justifyContent="center"
            >
                Projects
            </Typography>
        </Box>
    )
}
