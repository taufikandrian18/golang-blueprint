package constants

const (
	StaticOTP = "123456"

	DefaultIsActiveValue = true

	DefaultBaseDecimal = 10
	DefaultBitSize     = 64

	DefaultAllowHeaderToken         = "token"
	DefaultAllowHeaderRefreshToken  = "refresh-token"
	DefaultAllowHeaderAuthorization = "Authorization"

	DefaultOrderValue = "created_at DESC"

	LanguageID = "id"
	LanguageEN = "en"

	RoleAdministator = 1

	QuestionTypeScale = "scale"
	QuestionTypeText  = "text"

	PointCredit = "credit"
	PointDebit  = "debit"

	OnSchedule  = "1"
	OffSchedule = "0"

	CreatedAppointmentSuccess = "1"
	CreatedAppointmentFailed  = "0"
	CancelAppointmentSuccess  = "1"
	CancelAppointmentFailed   = "0"

	// status on db.
	AppointmentReservation     = "1"
	AppointmentStatusScheduled = "scheduled"
	AppointmentStatusDone      = "done"
	AppointmentStatusCanceled  = "canceled"
	AppointmentStatusCheckIn   = "check in"
	// status on simrs
	// status : 0 = masih reservasi, 1 = sedang dijalankan (sudah check in), 2 = selesai, 3 = batal.
	SIMRSAppointmentStatusReserved = 0
	SIMRSAppointmentStatusCheckIn  = 1
	SIMRSAppointmentStatusDone     = 2
	SIMRSAppointmentStatusCanceled = 3

	PatientExist           = "1"
	DoctorScheduleUnActive = "0"
	DoctorScheduleActive   = "1"

	MenuStatusActive   = "active"
	MenuStatusUnactive = "unactive"

	GroupCorporate   = "headquarters"
	GroupCorporateID = "8bfd0df8-f035-4329-af6b-fde686ad1f43"
	SuperadminID     = "81322b1d-8211-429f-9e0b-8ac6e1315043"
	AdminID          = "6afd4ed5-03e5-4015-9274-6f5061a17733"
)

const (
	DeviceIOS             = "IOS"
	DeviceAndroid         = "Android"
	DeviceWeb             = "web"
	FirebaseStatusSuccess = "success"
	FirebaseStatusFailed  = "failed"

	NotificationAppointment   = "4c3ebfc1-a360-4929-87f3-4da6d817ca0e"
	NotificationBroadcast     = "81dcb6c1-94eb-4975-b208-7376b9f10d3e"
	NotificationPayment       = "f18e6ef0-c334-4ea2-baff-75299f91c4b1"
	NotificationSurvey        = "e0eab3b9-7c0a-4910-9bd9-501d73cacf2c"
	NotificationRegristration = "390279f5-3e12-4e1f-99a5-02e535dc8fab"
	SurveyCategoryAppointment = "ea7a4623-5952-4630-8664-983bad7e69dd"

	NotificationStatusCreated = "created"
	NotificationStatusSent    = "sent"

	NotificationCategoryAppointment   = "appointment"
	NotificationCategoryBroadcast     = "broadcast"
	NotificationCategoryPayment       = "payment"
	NotificationCategorySurvey        = "survey"
	NotificationCategoryRegristration = "registration"

	TrxRegisPatient      = "registration-patient"
	PaymentStatusPaid    = "paid"
	PaymentStatusUnpaid  = "unpaid"
	PaymentStatusFailed  = "failed"
	PaymentStatusExpired = "expired"

	PaymentMethodVA      = "06f243bd-3f1f-4960-9d0c-b97c3d9ff870" // Transfer Virtual
	PaymentMethodCC      = "8cf0c909-46a2-4c66-a63a-543b17b8f690" // Credit Card
	PaymentMethodEWallet = "c183d395-a1d5-4772-8f7a-7ea61a7521fa" // Electronic Wallet
	PaymentMethodDebit   = "0e08ccc2-a6a1-491f-9b0e-1a66366bc0f7" // Direct Debit
	PaymentMethodRetail  = "f2ffeada-f4fd-4ae0-b2d5-37887d1cd8f7" // Retail
)

const (
	XenditEWalletStatusSuccess = "SUCCEEDED"
	XenditEWalletStatusFailed  = "FAILED"

	XenditDebitStatusSuccess = "COMPLETED"
	XenditDebitStatusFailed  = "FAILED"
	XenditEventDebitExpire   = "payment_method.expiry.expired"
	XenditEventDebitExpiring = "payment_method.expiry.expiring"
	XenditEventDebitPayment  = "direct_debit.payment"

	XenditDebitCardEXPIRED = "EXPIRED"
)

type ChannelCodeEnum string

type FailureCodeEwallet string

const (
	// --- Status ENUM ----.
	StatusActive      = "active"
	StatusDeleted     = "deleted"
	StatusComplete    = "complete"
	StatusClosed      = "closed"
	StatusUnqualified = "Unqualified"

	StatusSuspend = "suspend"

	StatusDraft            = "draft"
	StatusScheduledPublish = "scheduled publish"
	StatusPublished        = "published"

	StatusRejected           = "reject"
	StatusWaitingForApproval = "waiting for approval"
	StatusApproved           = "approved"

	StatusTerminated = "terminated"

	// --- Created By Constants Value ---.
	CreatedByTemporaryBySystem = "temporary by system"

	// --- Delimiter ---.
	DefaultDelimiterStringValue = "|"

	// --- Boolean ---.
	DefaultExpiredStatus  = false
	DefaultDoneStatus     = false
	DefaultArchivedStatus = false
	DefaultSeenStatus     = false

	// --- Board Task Priority ---
	BoardTaskPriorityLow    = "Low"
	BoardTaskPriorityMedium = "Medium"
	BoardTaskPriorityHigh   = "High"
	BoardTaskNoPriority     = "No Priority"

	// --- Board Task Filter ---
	NoAssignee    = "No Assignee"
	NoLabel       = "No Label"
	TaskNoDueDate = "No Dates"
	TaskOverDue   = "Overdue"
	TaskNextDay   = "Next Day"
	TaskNextWeek  = "Next Week"
	TaskNextMonth = "Next Month"

	// --- Board Template ---
	BoardTemplateKanban = "811ab326-8e05-409c-af7f-469255fa72f4"
)

// Time Format.
const (
	TimeDateFormat             = "2006-01-02"
	TimeWithSecondFormat       = "2006-01-02"
	FloatFormat                = "3.14"
	DefaultTimePipedriveFormat = "2006-01-02 15:04:05"
)

// POST CMS, COMMUNITY & EVENT
const (
	PostTypeCMS       = "cms"
	PostTypeArticle   = "article"
	PostTypeCommunity = "community"
)

const (
	PrimaryCalendarID = "c_bg5qdcajp11g29ttkkno395jp4@group.calendar.google.com"
)

const (
	SheetHotelOwner = "Hotel Owner"
	SheetWhiteLabel = "White Label"
	SheetGroup      = "Group"
	SheetBrand      = "Brand"
	SheetProperty   = "Property"
	SheetDepartment = "Department"
	SheetJob        = "Job"
)

// Activity Logging
const (
	DefaultMapSizeActivityLogging = 3

	ActivityLoggingFullName     = "fullName"
	ActivityLoggingTaskTitle    = "taskTitle"
	ActivityLoggingColumnTitle  = "columnTitle"
	ActivityLoggingColumnTitle2 = "columnTitle2"

	ActivityAddBoardTask     = "add-board-task"
	ActivityArchiveBoardTask = "archive-board-task"
	ActivityMoveBoardTask    = "move-board-task"
)

// Artist Satisfying Survey
const (
	SurveyTypeMultipleAnswer = "Multiple Answer"
	SurveyTypeEssay          = "Essay"
)

// Property Level
const (
	EnumPropertyLevelCollection    = "1"
	EnumPropertyLevelEconomy       = "2"
	EnumPropertyLevelMidScale      = "3"
	EnumPropertyLevelLifestyle     = "4"
	EnumPropertyLevelUpScale       = "5"
	EnumPropertyLevelFoodBeverages = "6"

	PropertyLevelCollection    = "Collections"
	PropertyLevelEconomy       = "Economy"
	PropertyLevelMidScale      = "Midscale"
	PropertyLevelLifestyle     = "Lifestyle"
	PropertyLevelUpScale       = "Upscale"
	PropertyLevelFoodBeverages = "Food & Beverages"
)

// Invitation Queue Status
const (
	InvitationQueueStatusQueued  = "queued"
	InvitationQueueStatusInvited = "invited"
)

// SQL Constraint
const (
	// Table Name
	TableEmployee = "employee"

	// Employee Unique Constraint
	UniqueNIKConstraint         = "unique_nik"
	UniqueIDCardConstraint      = "unique_id_card"
	UniqueNPWPConstraint        = "unique_npwp"
	UniqueEmailConstraint       = "unique_email"
	UniquePhoneNumberConstraint = "unique_phone_number"
)

const (
	QAStatusNew        = "new"
	QAStatusInProgress = "in_progress"
	QAStatusFinished   = "finished"

	QASubmissionStatusAvailable = "available"
	QASubmissionStatusFilled    = "filled"

	QAChecklistTypeScore  = "score"
	QAChecklistTypeNumber = "number"
	QAChecklistTypeText   = "text"

	QAValueC  = "C"
	QAValueNC = "NC"
	QAValueNA = "N/A"
)

// DepartmentGroup
const (
	EnumDepartmentUnitProperty = 1
	EnumDepartmentHeadquarters = 0

	DepartmentUnitProperty = "Unit Property"
	DepartmentHeadquarters = "Headquarters"
)
