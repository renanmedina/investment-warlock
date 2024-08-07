package announcements

import (
	"github.com/renanmedina/i-invest/internal/event_store"
)

func configureEventHandlers() map[string][]event_store.EventHandler {
	return map[string][]event_store.EventHandler{
		COMPANY_ANNOUNCEMENT_CREATED_EVENT_NAME: {
			event_store.NewSaveEventToStoreHandler(),
			NewDownloadAndSaveAnnouncementFileHandler(),
		},
		COMPANY_ANNOUNCEMENT_FILE_DOWNLOADED_AND_SAVED_EVENT_NAME: {
			event_store.NewSaveEventToStoreHandler(),
			NewNotifyCompanyNewAnnouncementToUserHandler(),
		},
		COMPANY_ANNOUNCEMENT_TRANSCRIBED_SUCCESSFULLY_EVENT_NAME: {
			event_store.NewSaveEventToStoreHandler(),
		},
	}
}
