import { Typography, Box } from '@mui/material'
import { Construction } from '@mui/icons-material'

interface PlaceholderPageProps {
  title: string
}

export function PlaceholderPage({ title }: PlaceholderPageProps) {
  return (
    <Box
      sx={{
        display: 'flex',
        flexDirection: 'column',
        alignItems: 'center',
        justifyContent: 'center',
        minHeight: '50vh',
        color: 'text.secondary',
      }}
    >
      <Construction sx={{ fontSize: 64, mb: 2 }} />
      <Typography variant="h5">{title}</Typography>
      <Typography variant="body1" sx={{ mt: 1 }}>
        功能开发中...
      </Typography>
    </Box>
  )
}
