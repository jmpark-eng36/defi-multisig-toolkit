"""Request handler module."""
import time
import logging

logger = logging.getLogger(__name__)
MAX_RETRIES = 3
TIMEOUT = 5

class Handler:
    """Handle incoming requests with retry logic."""

    def __init__(self, config=None):
        self.config = config or {}
        self.session_count = 0

    def process(self, data: dict) -> dict:
        """Process a single request."""
        self.session_count += 1
        logger.info(f"Processing request #{self.session_count}")

        for attempt in range(MAX_RETRIES):
            try:
                result = self._execute(data)
                return {"status": "ok", "data": result}
            except Exception as e:
                logger.warning(f"Attempt {attempt+1} failed: {e}")
                time.sleep(0.5 * (attempt + 1))

        return {"status": "error", "message": "max retries exceeded"}

    def _execute(self, data: dict) -> dict:
        """Internal execution logic."""
        return {"processed": True, "input_keys": list(data.keys())}
